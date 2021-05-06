package models

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/libp2p/go-msgio"
	"google.golang.org/protobuf/proto"
)

const K_CHUNK_SIZE = 4 * 1024

// ** ─── SONRFILE MANAGEMENT ────────────────────────────────────────────────────────
// Checks if File contains single item
func (f *SonrFile) IsSingle() bool {
	return len(f.Files) == 1
}

// Checks if Single File is Media
func (f *SonrFile) IsMedia() bool {
	return f.Payload == Payload_MEDIA || (f.IsSingle() && f.Single().Mime.IsMedia())
}

// Checks if File contains multiple items
func (f *SonrFile) IsMultiple() bool {
	return len(f.Files) > 1
}

// Returns SonrFile as TransferCard given Receiver and Owner
func (f *SonrFile) CardIn(receiver *Peer, owner *Peer) *TransferCard {
	// Update Direction
	f.Direction = SonrFile_Default

	// Create Card
	return &TransferCard{
		// SQL Properties
		Payload:  f.Payload,
		Received: int32(time.Now().Unix()),

		// Owner Properties
		Owner:    owner.GetProfile(),
		Receiver: receiver.GetProfile(),

		// Data Properties
		File: f,
	}
}

// Returns SonrFile as TransferCard given Receiver and Owner
func (f *SonrFile) CardOut(receiver *Peer, owner *Peer) *TransferCard {
	// Create Card
	return &TransferCard{
		// SQL Properties
		Payload: f.Payload,

		// Owner Properties
		Receiver: receiver.GetProfile(),
		Owner:    owner.GetProfile(),
		File:     f,
	}
}

// Method Returns Metadata Item at Given Index
func (f *SonrFile) ItemAtIndex(index int) *SonrFile_Metadata {
	return f.Files[index]
}

// Method Returns Preview from Thumbnail if Single File
func (f *SonrFile) Preview() []byte {
	// Validate Single
	if f.IsSingle() {
		// Retrieve Meta
		meta := f.Files[0]
		props := meta.GetProperties()

		// Check if Thumbnail Provided
		if props.HasThumbnail {
			// Initialize
			thumbWriter := new(bytes.Buffer)
			thumbReader := bytes.NewReader(meta.Thumbnail)

			// Convert to Image Object
			img, _, err := image.Decode(thumbReader)
			if err != nil {
				log.Println(err)
				return nil
			}

			// Encode as Jpeg into buffer w/o scaling
			err = jpeg.Encode(thumbWriter, img, nil)
			if err != nil {
				log.Panicln(err)
				return nil
			}

			return thumbWriter.Bytes()
		}
	}
	return nil
}

// Method Returns Single if Applicable
func (f *SonrFile) Single() *SonrFile_Metadata {
	if f.IsSingle() {
		return f.Files[0]
	} else {
		return nil
	}
}

// ** ─── SONRFILE_Metadata MANAGEMENT ────────────────────────────────────────────────────────
// Returns Progress of File, Given the written number of bytes
func (m *SonrFile_Metadata) Progress(curr int, n int) (bool, float32) {
	// Calculate Tracking
	itemChunks := m.Size / K_CHUNK_SIZE
	interval := int(itemChunks / 100)

	// Check Interval
	if interval != 0 {
		if curr%interval == 0 {
			return true, float32(n) / float32(m.Size)
		}
	}
	return false, 0
}

// Returns/Updates Save Path for this File
func (m *SonrFile_Metadata) SetPath(d *Device) string {
	// Initialize
	var path string

	// Check for Media
	if m.Mime.IsMedia() {
		// Check for Desktop
		if d.IsDesktop() {
			path = filepath.Join(d.FileSystem.GetDownloads(), m.Name)
		} else {
			path = filepath.Join(d.FileSystem.GetTemporary(), m.Name)
		}
	} else {
		// Check for Desktop
		if d.IsDesktop() {
			path = filepath.Join(d.FileSystem.GetDownloads(), m.Name)
		} else {
			path = filepath.Join(d.FileSystem.GetDocuments(), m.Name)
		}
	}

	// Set Path and Return
	m.Path = path
	return m.Path
}

// Metadata Info returns: Total Bytes, Total Chunks, error
func (m *SonrFile_Metadata) WriteTo(writer msgio.WriteCloser, call NodeCallback) error {
	// @ Open Os File
	f, err := os.Open(m.Path)
	if err != nil {
		return errors.New(fmt.Sprintf("Error to read [file=%v]: %v", m.Name, err.Error()))
	}

	defer f.Close()

	// @ Initialize Chunk Data
	nBytes, nChunks := int32(0), int32(0)
	r := bufio.NewReader(f)
	buf := make([]byte, 0, K_CHUNK_SIZE)

	// @ Loop through File
	for {
		// Reads bytes onto chunk
		n, err := r.Read(buf[:cap(buf)])

		// Set Current chunk Value  //
		buf = buf[:n]

		// Bytes read is zero
		if n == 0 {
			// No Error
			if err == nil {
				continue
			}

			// End of File
			if err == io.EOF {
				break
			}

			// Unexpected Error
			return err
		}

		// * Process Here: Increase Chunk/TotalBytes count
		nChunks++
		nBytes += int32(len(buf))

		// Write to Stream
		writeBufToStream(buf, false, writer, call)

		// Unexpected Error
		if err != nil && err != io.EOF {
			return err
		}
	}

	// Send Completed
	writeBufToStream(nil, true, writer, call)
	return nil
}

// Writes data to provided writer until completed is called
func writeBufToStream(buf []byte, completed bool, writer msgio.WriteCloser, call NodeCallback) {
	if !completed {
		// Create Block Protobuf from Chunk
		chunk := &Chunk{
			Size:       int32(len(buf)),
			Buffer:     buf,
			IsComplete: false,
		}

		// Convert to bytes
		bytes, err := proto.Marshal(chunk)
		if err != nil {
			call.Error(NewError(err, ErrorMessage_OUTGOING))
		}

		// Write Message Bytes to Stream
		err = writer.WriteMsg(bytes)
		if err != nil {
			call.Error(NewError(err, ErrorMessage_OUTGOING))
		}
	} else {
		// Create Block Protobuf from Chunk
		chunk := &Chunk{
			IsComplete: true,
		}

		// Convert to bytes
		bytes, err := proto.Marshal(chunk)
		if err != nil {
			call.Error(NewError(err, ErrorMessage_OUTGOING))
		}

		// Write Message Bytes to Stream
		err = writer.WriteMsg(bytes)
		if err != nil {
			call.Error(NewError(err, ErrorMessage_OUTGOING))
		}
	}
}