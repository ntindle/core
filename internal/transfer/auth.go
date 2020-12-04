package transfer

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ****************** //
// ** GRPC Service ** //
// ****************** //
// Argument is AuthMessage protobuf
type AuthArgs struct {
	Data []byte
	From string
}

// Reply is also AuthMessage protobuf
type AuthReply struct {
	Data     []byte
	Decision bool
}

// Service Struct
type AuthService struct {
	// Current Data
	currArgs   AuthArgs
	currReply  *AuthReply
	inviteCall OnProtobuf
}

// GRPC Callback
type OnGRPCall func(data []byte, from string) error

// ^ Calls Invite on Remote Peer ^ //
func (as *AuthService) InviteRequest(ctx context.Context, args AuthArgs, reply *AuthReply) error {
	log.Println("Received a Invite call: ", args.Data)
	// Set Current Data
	as.currArgs = args
	as.currReply = reply

	// Send Callback
	as.inviteCall(args.Data)
	return nil
}

// ^ Send SendInvite to a Peer ^ //
func (pc *PeerConnection) SendInvite(h host.Host, id peer.ID, info *md.Peer, sm *sf.SafeMetadata) {
	// Set SafeFile
	pc.safeFile = sm

	// Create Client
	rpcClient := gorpc.NewClient(h, protocol.ID("/sonr/rpc/auth"))

	// Create Invite Message
	reqMsg := md.AuthMessage{
		Event:    md.AuthMessage_REQUEST,
		From:     info,
		Metadata: sm.GetMetadata(),
	}

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(&reqMsg)
	if err != nil {
		onError(err, "sendInvite")
		log.Println(err)
	}

	// Set Data
	var reply AuthReply
	var args AuthArgs
	args.Data = msgBytes
	startTime := time.Now()

	// Call to Peer
	err = rpcClient.Call(id, "AuthService", "InviteRequest", args, &reply)
	if err != nil {
		// Track Execution
		endTime := time.Now()
		diff := endTime.Sub(startTime)
		log.Printf("Failed to call %s: time=%s\n", id, diff)

		// Send Error
		onError(err, "sendInvite")
		log.Panicln(err)
	}

	// End Tracking
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	log.Printf("Response from %s: time=%s\n", id, diff)

	// Send Callback and Reset
	pc.respondedCall(reply.Data)

	// Handle Response
	if reply.Decision {
		// Begin Transfer
		pc.SendFile(h)
	}
}

// ^ Send Accept Message on Stream ^ //
func (pc *PeerConnection) SendResponse(decision bool, selfInfo *md.Peer) {
	// Initialize Message
	var respMsg *md.AuthMessage

	// Check Decision
	if decision {
		if pc.currMessage != nil {
			// Initialize Transfer
			savePath := "/" + pc.currMessage.Metadata.Name + "." + pc.currMessage.Metadata.Mime.Subtype
			pc.transfer = NewTransfer(savePath, pc.currMessage.Metadata, pc.currMessage.From, pc.progressCall, pc.completedCall)

			// Create Accept Response
			respMsg = &md.AuthMessage{
				From:  selfInfo,
				Event: md.AuthMessage_ACCEPT,
			}
		} else {
			err := errors.New("AuthMessage wasnt cached")
			onError(err, "sendInvite")
			log.Panicln(err)
		}
	} else {
		// Reset Peer Info
		pc.peerID = ""
		pc.currMessage = nil

		// Create Decline Response
		respMsg = &md.AuthMessage{
			From:  selfInfo,
			Event: md.AuthMessage_DECLINE,
		}
	}

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(respMsg)
	if err != nil {
		onError(err, "sendInvite")
		log.Println(err)
	}

	pc.ascv.currReply.Data = msgBytes
	pc.ascv.currReply.Decision = decision
}
