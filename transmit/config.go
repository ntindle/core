package transmit

import (
	"errors"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/protocol"
)

// Transfer Protocol ID's
const (
	FilePID       protocol.ID = "/transmit/file/0.0.1"
	DonePID       protocol.ID = "/transmit/done/0.0.1"
	ITEM_INTERVAL             = 25
)

// Error Definitions
var (
	logger              = golog.Default.Child("protocols/transmit")
	ErrNoSession        = errors.New("Failed to get current session, set to nil")
	ErrFailedAuth       = errors.New("Failed to Authenticate message")
	ErrInvalidDirection = errors.New("Direction was not set")
)

// Option is a function that can be applied to ExchangeProtocol config
type Option func(*options)

// options for ExchangeProtocol config
type options struct {
	interval int
}

// defaultOptions for ExchangeProtocol config
func defaultOptions() *options {
	return &options{}
}

// Apply applies the options to the ExchangeProtocol
func (o *options) Apply(p *TransmitProtocol) error {
	// Apply options
	p.mode = p.node.Role()
	return nil
}
