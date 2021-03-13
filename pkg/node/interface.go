package node

import (
	"context"
	"log"

	olc "github.com/google/open-location-code/go"
	"github.com/libp2p/go-libp2p-core/host"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	sh "github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/lobby"
	tr "github.com/sonr-io/core/internal/transfer"
	dq "github.com/sonr-io/core/pkg/data"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ Interface: Callback is implemented from Plugin to receive updates ^
type Callback interface {
	OnRefreshed(data []byte)   // Lobby Updates
	OnEvent(data []byte)       // Lobby Event
	OnInvited(data []byte)     // User Invited
	OnDirected(data []byte)    // User Direct-Invite from another Device
	OnResponded(data []byte)   // Peer has responded
	OnProgress(data float32)   // File Progress Updated
	OnReceived(data []byte)    // User Received File
	OnTransmitted(data []byte) // User Sent File
	OnError(data []byte)       // Internal Error
}

// ^ Struct: Main Node handles Networking/Identity/Streams ^
type Node struct {
	// Properties
	ctx         context.Context
	olc         string
	directories *md.Directories
	device      *md.Device
	peer        *md.Peer
	contact     *md.Contact
	status      md.Status

	// Networking Properties
	host   host.Host
	pubSub *pubsub.PubSub

	// References
	call     Callback
	lobby    *lobby.Lobby
	peerConn *tr.TransferController
	queue    *dq.FileQueue
}

// ^ NewNode Initializes Node with a host and default properties ^
func NewNode(reqBytes []byte, call Callback) *Node {
	// ** Create Context and Node - Begin Setup **
	node := new(Node)
	node.ctx = context.Background()
	node.call = call

	// err := sentry.Init(sentry.ClientOptions{
	// 	Dsn: "https://cbf88b01a5a5468fa77101f7dfc54f20@o549479.ingest.sentry.io/5672329",
	// })
	// if err != nil {
	// 	log.Fatalf("sentry.Init: %s", err)
	// }

	// ** Unmarshal Request **
	req := md.ConnectionRequest{}
	err := proto.Unmarshal(reqBytes, &req)
	if err != nil {
		node.error(err, "NewNode")
		return nil
	}

	// Create New Profile from Request
	profile := &md.Profile{
		Username:  req.GetUsername(),
		FirstName: req.Contact.GetFirstName(),
		LastName:  req.Contact.GetLastName(),
		Picture:   req.Contact.GetPicture(),
		Platform:  req.Device.GetPlatform(),
	}

	// @1. Set OLC, Create Host, and Start Discovery
	node.queue = dq.NewQueue(req.Directories, profile, node.queued, node.multiQueued, node.error)
	node.status = md.Status_NONE
	node.olc = olc.Encode(float64(req.Latitude), float64(req.Longitude), 8)
	node.host, err = sh.NewHost(node.ctx, req.Directories, node.olc)
	if err != nil {
		node.error(err, "NewNode")
		return nil
	}

	// @3. Set Node User Information
	if err = node.setInfo(&req, profile); err != nil {
		node.error(err, "NewNode")
		return nil
	}

	// @4. Setup Connection w/ Lobby and Set Stream Handlers
	if err = node.setConnection(node.ctx); err != nil {
		node.error(err, "NewNode")
		return nil
	}

	// ** Callback Node User Information ** //
	return node
}

// ^ queued Callback, Sends File Invite to Peer, and Notifies Client ^
func (sn *Node) queued(card *md.TransferCard, req *md.InviteRequest) {
	// Get PeerID
	id, _, err := sn.lobby.Find(req.To.Id.Peer)

	// Check error
	if err != nil {
		sn.error(err, "Queued")
	}

	// Retreive Current File
	currFile := sn.queue.CurrentFile()
	card.Status = md.TransferCard_INVITE
	sn.peerConn.NewOutgoing(currFile)

	// Create Invite Message
	invMsg := md.AuthInvite{
		From:    sn.peer,
		Payload: card.Payload,
		Card:    card,
	}

	// Check if ID in PeerStore
	go func(inv *md.AuthInvite) {
		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(inv)
		if err != nil {
			sn.error(err, "Marshal")
		}

		sn.peerConn.Request(sn.host, id, msgBytes)
	}(&invMsg)
}

// ^ multiQueued Callback, Sends File Invite to Peer, and Notifies Client ^
func (sn *Node) multiQueued(card *md.TransferCard, req *md.InviteRequest, count int) {
	// Get PeerID
	id, _, err := sn.lobby.Find(req.To.Id.Peer)

	// Check error
	if err != nil {
		sn.error(err, "Queued")
	}

	// Retreive Current File
	currFile := sn.queue.CurrentFile()
	card.Status = md.TransferCard_INVITE
	sn.peerConn.NewOutgoing(currFile)

	// Create Invite Message
	invMsg := md.AuthInvite{
		From:    sn.peer,
		Payload: card.Payload,
		Card:    card,
	}

	// Check if ID in PeerStore
	go func(inv *md.AuthInvite) {
		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(inv)
		if err != nil {
			sn.error(err, "Marshal")
		}

		sn.peerConn.Request(sn.host, id, msgBytes)
	}(&invMsg)
}

// ^ invite Callback with data for Lifecycle ^ //
func (sn *Node) invited(data []byte) {
	// Update Status
	sn.status = md.Status_INVITED
	// Callback with Data
	sn.call.OnInvited(data)
}

// ^ transmitted Callback middleware post transfer ^ //
func (sn *Node) transmitted(peer *md.Peer) {
	// Update Status
	sn.status = md.Status_AVAILABLE

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(peer)
	if err != nil {
		log.Println(err)
	}

	// Callback with Data
	sn.call.OnTransmitted(msgBytes)
}

// ^ received Callback middleware post transfer ^ //
func (sn *Node) received(card *md.TransferCard) {
	// Update Status
	sn.status = md.Status_AVAILABLE

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(card)
	if err != nil {
		log.Println(err)
	}

	// Callback with Data
	sn.call.OnReceived(msgBytes)
}

// ^ error Callback with error instance, and method ^
func (sn *Node) error(err error, method string) {
	// Create Error ProtoBuf
	errorMsg := md.ErrorMessage{
		Message: err.Error(),
		Method:  method,
	}

	// Convert Message to bytes
	bytes, err := proto.Marshal(&errorMsg)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}
	// Send Callback
	sn.call.OnError(bytes)

	// Log In Core
	log.Fatalf("[Error] At Method %s : %s", err.Error(), method)
}
