package models

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/h2non/filetype"
)

// ^ Struct returned on GetInfo() Generate Preview/Metadata
type FileInfo struct {
	Mime    *MIME
	Payload Payload
	Name    string
	Path    string
	Size    int32
	IsImage bool
}

// ^ Struct to Pass Invite and Response Info
type AuthOpts struct {
	Decision bool
	Contact  *Contact
	Peer     *Peer
	Invite   *AuthInvite
	IsCancel bool
}

// ^ Method Returns File Info at Path ^ //
func GetFileInfo(path string) (FileInfo, error) {
	// Initialize
	var mime *MIME
	var payload Payload

	// @ 1. Get File Information
	// Open File at Path
	file, err := os.Open(path)
	if err != nil {
		return FileInfo{}, err
	}
	defer file.Close()

	// Get Info
	info, err := file.Stat()
	if err != nil {
		return FileInfo{}, err
	}

	// Read File to required bytes
	head := make([]byte, 261)
	_, err = file.Read(head)
	if err != nil {
		return FileInfo{}, err
	}

	// Get File Type
	kind, err := filetype.Match(head)
	if err != nil {
		return FileInfo{}, err
	}

	// @ 2. Create Mime Protobuf
	mime = &MIME{
		Type:    MIME_Type(MIME_Type_value[kind.MIME.Type]),
		Subtype: kind.MIME.Subtype,
		Value:   kind.MIME.Value,
	}

	// @ 3. Find Payload
	if mime.Type == MIME_image || mime.Type == MIME_video || mime.Type == MIME_audio {
		payload = Payload_MEDIA
	} else {
		// Get Extension
		ext := filepath.Ext(path)

		// Cross Check Extension
		if ext == ".pdf" {
			payload = Payload_PDF
		} else if ext == ".ppt" || ext == ".pptx" {
			payload = Payload_PRESENTATION
		} else if ext == ".xls" || ext == ".xlsm" || ext == ".xlsx" || ext == ".csv" {
			payload = Payload_SPREADSHEET
		} else if ext == ".txt" || ext == ".doc" || ext == ".docx" || ext == ".ttf" {
			payload = Payload_TEXT
		} else {
			payload = Payload_UNDEFINED
		}
	}

	// Return Object
	return FileInfo{
		Mime:    mime,
		Payload: payload,
		Name:    strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
		Path:    path,
		Size:    int32(info.Size()),
		IsImage: filetype.IsImage(head),
	}, nil
}

// ^ Method Generates new Transfer Card from Contact^ //
func NewCardFromContact(p *Peer, c *Contact, status TransferCard_Status) TransferCard {
	return TransferCard{
		// SQL Properties
		Payload:  Payload_CONTACT,
		Received: int32(time.Now().Unix()),
		Preview:  p.Profile.Picture,
		Platform: p.Platform,

		// Transfer Properties
		Status: status,

		// Owner Properties
		Username:  p.Profile.Username,
		FirstName: p.Profile.FirstName,
		LastName:  p.Profile.LastName,

		// Data Properties
		Contact: c,
	}
}

// ^ Method Generates new Transfer Card from URL ^ //
func NewCardFromUrl(p *Peer, url string, status TransferCard_Status) TransferCard {
	// Get URL Data
	urlInfo, err := GetPageInfoFromUrl(url)
	if err != nil {
		log.Println(err)

		// Return Card
		return TransferCard{
			// SQL Properties
			Payload:  Payload_URL,
			Received: int32(time.Now().Unix()),
			Platform: p.Platform,

			// Transfer Properties
			Status: status,

			// Owner Properties
			Username:  p.Profile.Username,
			FirstName: p.Profile.FirstName,
			LastName:  p.Profile.LastName,

			// Data Properties
			Url: &URLLink{
				Link: url,
			},
		}
	} else {
		// Return Card
		return TransferCard{
			// SQL Properties
			Payload:  Payload_URL,
			Received: int32(time.Now().Unix()),
			Platform: p.Platform,

			// Transfer Properties
			Status: status,

			// Owner Properties
			Username:  p.Profile.Username,
			FirstName: p.Profile.FirstName,
			LastName:  p.Profile.LastName,

			// Data Properties
			Url: urlInfo,
		}
	}
}

// ^ Method Creates AuthInvite from Request ^ //
func NewInviteFromRequest(req *InviteRequest, p *Peer) AuthInvite {
	// Initialize
	var card TransferCard
	var payload Payload

	// Determine Payload
	if req.Type == InviteRequest_Contact {
		payload = Payload_CONTACT
		card = NewCardFromContact(p, req.Contact, TransferCard_DIRECT)
	} else if req.Type == InviteRequest_URL {
		payload = Payload_URL
		card = NewCardFromUrl(p, req.Url, TransferCard_DIRECT)
	} else {
		payload = Payload_UNDEFINED
	}

	// Return Protobuf
	return AuthInvite{
		From:    p,
		Payload: payload,
		Card:    &card,
	}
}

// ^ Method Creates AuthReply from Options ^ //
func NewReplyFromDecision(opts AuthOpts) AuthReply {
	// Initialize
	offerMsg := opts.Invite
	decision := opts.Decision
	peer := opts.Peer
	contact := opts.Contact

	// @ Check for Cancel
	if opts.IsCancel {
		return AuthReply{
			From:     peer,
			Type:     AuthReply_None,
			Decision: false,
		}
	}

	// @ Pass Contact Back
	if offerMsg.Payload == Payload_CONTACT {
		// Create Accept Response
		card := NewCardFromContact(peer, contact, TransferCard_REPLY)
		return AuthReply{
			From: peer,
			Type: AuthReply_Contact,
			Card: &card,
		}
	} else {
		// @ Check Decision
		if decision {
			// Create Accept Response
			return AuthReply{
				From:     peer,
				Decision: true,
				Type:     AuthReply_Transfer,
			}
		} else {
			// Create Decline Response
			return AuthReply{
				From:     peer,
				Decision: false,
				Type:     AuthReply_Transfer,
			}
		}
	}
}