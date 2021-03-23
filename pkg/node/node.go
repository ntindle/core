package node

import (
	"context"
	"fmt"
	"log"
	"time"

	sentry "github.com/getsentry/sentry-go"
	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/host"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	swarm "github.com/libp2p/go-libp2p-swarm"
	"github.com/pkg/errors"

	discLimit "github.com/libp2p/go-libp2p-core/discovery"
	disc "github.com/libp2p/go-libp2p-discovery"
	sl "github.com/sonr-io/core/internal/lobby"
	md "github.com/sonr-io/core/internal/models"
	net "github.com/sonr-io/core/internal/network"
	tf "github.com/sonr-io/core/internal/transfer"
	tr "github.com/sonr-io/core/internal/transfer"
	dq "github.com/sonr-io/core/pkg/user"
)

const discoveryInterval = time.Second * 2
const gracePeriod = time.Minute

// ^ Struct: Main Node handles Networking/Identity/Streams ^
type Node struct {
	// Properties
	ctx     context.Context
	contact *md.Contact
	device  *md.Device
	fs      *dq.SonrFS
	peer    *md.Peer
	profile *md.Profile

	// Networking Properties
	host   host.Host
	kadDHT *dht.IpfsDHT
	pubSub *pubsub.PubSub
	router *net.ProtocolRouter
	status md.Status

	// References
	call     Callback
	lobby    *sl.Lobby
	transfer *tr.TransferController
}

// ^ NewNode Initializes Node with a host and default properties ^
func NewNode(req *md.ConnectionRequest, call Callback) *Node {
	// Initialize Node Logging
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://cbf88b01a5a5468fa77101f7dfc54f20@o549479.ingest.sentry.io/5672329",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	// Create Context and Set Node Properties
	node := new(Node)
	node.ctx = context.Background()
	node.call = call

	// Create New Profile from Request
	node.profile = &md.Profile{
		Username:  req.GetUsername(),
		FirstName: req.Contact.GetFirstName(),
		LastName:  req.Contact.GetLastName(),
		Picture:   req.Contact.GetPicture(),
		Platform:  req.Device.GetPlatform(),
	}

	// Set File System
	node.fs = dq.InitFS(req, node.profile, node.FSCallback())
	node.router = net.NewProtocolRouter(req)

	// Set Default Properties
	node.contact = req.Contact
	node.device = req.Device
	node.status = md.Status_NONE
	return node
}

// ^ Start Begins Running Libp2p Host ^
func (n *Node) Start(opts *HostOptions) bool {
	// IP Address
	ip4 := net.IPv4()
	ip6 := net.IPv6()

	// Get Private Key
	privKey, err := n.fs.GetPrivateKey()
	if err != nil {
		sentry.CaptureException(err)
		n.call.OnConnected(false)
		return false
	}

	// Start Host
	h, err := libp2p.New(
		n.ctx,
		// Add listening Addresses
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0", ip4),
			fmt.Sprintf("/ip6/%s/tcp/0", ip6)),
		libp2p.Identity(privKey),
		libp2p.DefaultTransports,
		libp2p.ConnectionManager(connmgr.NewConnManager(
			10,          // Lowwater
			20,          // HighWater,
			gracePeriod, // GracePeriod
		)),
	)
	if err != nil {
		sentry.CaptureException(err)
		n.call.OnConnected(false)
		return false
	}
	n.host = h

	// Create Pub Sub
	ps, err := pubsub.NewGossipSub(n.ctx, n.host)
	if err != nil {
		sentry.CaptureException(err)
		n.call.OnConnected(false)
		return false
	}
	n.pubSub = ps

	// Set Peer Info
	n.peer = &md.Peer{
		Id:       n.fs.GetPeerID(n.device, n.profile, n.host.ID().String()),
		Profile:  n.profile,
		Platform: n.device.Platform,
		Model:    n.device.Model,
	}
	n.call.OnConnected(true)
	return true
}

// ^ Bootstrap begins bootstrap with peers ^
func (n *Node) Bootstrap(opts *HostOptions) bool {
	// Create Bootstrapper Info
	bootstrappers := opts.GetBootstrapAddrInfo()

	// Set DHT
	kadDHT, err := dht.New(
		n.ctx,
		n.host,
		dht.BootstrapPeers(bootstrappers...),
	)
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "Error while Creating routing DHT"))
		n.error(err, "Error while Creating routing DHT")
		n.call.OnReady(false)
		return false
	}
	n.kadDHT = kadDHT

	// Bootstrap DHT
	if err := kadDHT.Bootstrap(n.ctx); err != nil {
		sentry.CaptureException(errors.Wrap(err, "Error while Bootstrapping DHT"))
		n.error(err, "Error while Bootstrapping DHT")
		n.call.OnReady(false)
		return false
	}

	// Connect to bootstrap nodes, if any
	hasBootstrapped := false
	for _, pi := range bootstrappers {
		if err := n.host.Connect(n.ctx, pi); err == nil {
			hasBootstrapped = true
			break
		}
	}

	// Check if Bootstrapping Occurred
	if !hasBootstrapped {
		sentry.CaptureException(errors.New("Failed to connect to any bootstrap peers"))
		return false
	} else {
		n.call.OnReady(true)
	}

	// Set Routing Discovery, Find Peers
	routingDiscovery := disc.NewRoutingDiscovery(kadDHT)
	disc.Advertise(n.ctx, routingDiscovery, n.router.Point(), discLimit.TTL(discoveryInterval))
	go n.handlePeers(routingDiscovery)

	// Enter Lobby
	if n.lobby, err = sl.Join(n.ctx, n.LobbyCallback(), n.host, n.pubSub, n.peer, n.router); err != nil {
		sentry.CaptureException(err)
		n.error(err, "Joining Lobby")
		n.call.OnReady(false)
		return false
	}

	// Initialize Peer Connection
	if n.transfer, err = tf.Initialize(n.ctx, n.host, n.pubSub, n.fs, n.router, n.TransferCallback()); err != nil {
		sentry.CaptureException(err)
		n.error(err, "Initializing Transfer Controller")
		n.call.OnReady(false)
		return false
	}
	return true
}

// ^ Handles Peers in DHT ^
func (n *Node) handlePeers(routingDiscovery *disc.RoutingDiscovery) {
	for {
		// Find peers in DHT
		peersChan, err := routingDiscovery.FindPeers(
			n.ctx,
			n.router.Point(),
			discLimit.Limit(16),
		)
		if err != nil {
			n.error(err, "Finding DHT Peers")
			n.call.OnReady(false)
			return
		}

		// Iterate over Channel
		for pi := range peersChan {
			// Validate not Self
			if pi.ID != n.host.ID() {
				// Connect to Peer
				if err := n.host.Connect(n.ctx, pi); err != nil {
					// Capture Error
					sentry.CaptureException(errors.Wrap(err, "Failed to connect to peer in namespace"))

					// Remove Peer Reference
					n.host.Peerstore().ClearAddrs(pi.ID)
					if sw, ok := n.host.Network().(*swarm.Swarm); ok {
						sw.Backoff().Clear(pi.ID)
					}
				}
			}
		}

		// Refresh table every 4 seconds
		md.GetState().NeedsWait()
		<-time.After(discoveryInterval)
	}
}
