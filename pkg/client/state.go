package client

import (
	net "github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/room"
	srv "github.com/sonr-io/core/internal/service"
	s "github.com/sonr-io/core/internal/state"
	data "github.com/sonr-io/core/pkg/data"
)

const (
	// Connection: States, Each State needs an Action
	Connecting       s.StateType = "Connecting"
	ConnectionFailed s.StateType = "ConnectionFailed"
	Connected        s.StateType = "Connected"

	// Connection: Events
	Connect           s.EventType = "Connect"
	FailConnection    s.EventType = "FailConnection"
	SucceedConnection s.EventType = "SucceedConnection"

	// Bootstrap: States, Each State needs an Action
	Bootstrapping   s.StateType = "Bootstrapping"
	BootstrapFailed s.StateType = "BootstrapFailed"
	Bootstrapped    s.StateType = "Bootstrapped"

	// Bootstrap: Events
	Bootstrap        s.EventType = "Bootstrap"
	FailBootstrap    s.EventType = "FailBootstrap"
	SucceedBootstrap s.EventType = "SucceedBootstrap"
)

type ConnectionContext struct {
	connReq *data.ConnectionRequest
	c       *client
	peer    *data.Peer
	err     *data.SonrError
}

type ConnectingAction struct{}

func (a *ConnectingAction) Execute(eventCtx s.EventContext) s.EventType {
	ctx := eventCtx.(*ConnectionContext)

	// Set Request
	ctx.c.request = ctx.connReq
	ctx.c.isLinker = ctx.connReq.GetIsLinker()

	// Set Host
	hn, err := net.NewHost(ctx.c.ctx, ctx.connReq, ctx.c.account.AccountKeys(), ctx.c.emitter)
	if err != nil {
		ctx.err = err
		return FailConnection
	}

	// Get MultiAddrs
	maddr, err := hn.MultiAddr()
	if err != nil {
		ctx.err = err
		return FailConnection
	}

	// Set Peer
	peer, _ := ctx.c.account.CurrentDevice().SetPeer(hn.ID(), maddr, ctx.connReq.GetIsLinker())
	ctx.peer = peer

	// Set Host
	ctx.c.Host = hn
	return SucceedConnection
}

type BootstrapContext struct {
	connReq *data.ConnectionRequest
	c       *client
	room    *room.RoomManager
	peer    *data.Peer
	err     *data.SonrError
}

type BootstrappingAction struct{}

func (a *BootstrappingAction) Execute(eventCtx s.EventContext) s.EventType {
	ctx := eventCtx.(*BootstrapContext)

	// Bootstrap Host
	err := ctx.c.Host.Bootstrap(ctx.c.account.CurrentDevice().GetId())
	if err != nil {
		ctx.err = err
		return FailBootstrap
	}

	// Start Services
	s, err := srv.NewService(ctx.c.ctx, ctx.c.Host, ctx.c.account.CurrentDevice(), ctx.c.request, ctx.c.emitter)
	if err != nil {
		ctx.err = err
		return FailBootstrap
	}
	ctx.c.Service = s

	// Join Local
	RoomName := ctx.c.account.CurrentDevice().NewLocalRoom(ctx.connReq.GetServiceOptions())
	if t, err := room.JoinRoom(ctx.c.ctx, ctx.c.Host, ctx.c.account, RoomName, ctx.c.emitter); err != nil {
		ctx.err = err
		return FailBootstrap
	} else {
		// Check if Auto Update Events
		if ctx.connReq.GetServiceOptions().GetAutoUpdate() {
			go ctx.c.sendPeriodicRoomEvents(t)
		}
		ctx.room = t
		return SucceedBootstrap
	}
}

func newStateMachine() *s.StateMachine {
	return &s.StateMachine{
		States: s.States{
			s.Default: s.State{
				Events: s.Events{
					Connect: Connecting,
				},
			},
			Connecting: s.State{
				Action: &ConnectingAction{},
				Events: s.Events{
					SucceedConnection: Bootstrapping,
					FailConnection:    ConnectionFailed,
				},
			},
			Bootstrapping: s.State{
				Action: &BootstrappingAction{},
				Events: s.Events{
					SucceedBootstrap: Bootstrapped,
					FailBootstrap:    BootstrapFailed,
				},
			},
		},
	}
}
