package transmit

import (
	"bufio"
	"errors"
	"io"
	"os"
	"time"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/device"
	"github.com/sonr-io/core/node"
	"github.com/sonr-io/core/types/go/common/v1"
	"github.com/sonr-io/core/types/go/node/motor/v1"
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

// NewInSession creates a new Session from the given payload with Incoming direction.
func NewInSession(payload *common.Payload, from *common.Peer, to *common.Peer) *Session {
	// Create Session Items
	sessionPayload := NewSessionPayload(payload)
	return &Session{
		Direction:    common.Direction_DIRECTION_INCOMING,
		Payload:      payload,
		From:         from,
		To:           to,
		LastUpdated:  int64(time.Now().Unix()),
		Items:        sessionPayload.CreateItems(common.Direction_DIRECTION_INCOMING),
		CurrentIndex: 0,
		Results:      make(map[int32]bool),
	}
}

// NewOutSession creates a new Session from the given payload with Outgoing direction.
func NewOutSession(payload *common.Payload, to *common.Peer, from *common.Peer) *Session {
	// Create Session Items
	sessionPayload := NewSessionPayload(payload)
	return &Session{
		Direction:    common.Direction_DIRECTION_OUTGOING,
		Payload:      payload,
		To:           to,
		From:         from,
		LastUpdated:  int64(time.Now().Unix()),
		Items:        sessionPayload.CreateItems(common.Direction_DIRECTION_OUTGOING),
		CurrentIndex: 0,
		Results:      make(map[int32]bool),
	}
}

// FinalIndex returns the final index of the session.
func (s *Session) FinalIndex() int {
	return len(s.Items) - 1
}

// HasRead returns true if all files have been read.
func (s *Session) HasRead() bool {
	return s.IsIn() && s.IsDone()
}

// HasWrote returns true if all files have been written.
func (s *Session) HasWrote() bool {
	return s.IsOut() && s.IsDone()
}

// IsDone returns true if all files have been read or written.
func (s *Session) IsDone() bool {
	return int(s.GetCurrentIndex()) >= s.FinalIndex()
}

// IsOut returns true if the session is outgoing.
func (s *Session) IsOut() bool {
	return s.Direction == common.Direction_DIRECTION_OUTGOING
}

// IsIn returns true if the session is incoming.
func (s *Session) IsIn() bool {
	return s.Direction == common.Direction_DIRECTION_INCOMING
}

// Event returns the complete event for the session.
func (s *Session) Event() *motor.OnTransmitCompleteResponse {
	return &motor.OnTransmitCompleteResponse{
		From:       s.GetFrom(),
		To:         s.GetTo(),
		Direction:  s.GetDirection(),
		Payload:    s.GetPayload(),
		CreatedAt:  s.GetPayload().GetCreatedAt(),
		ReceivedAt: int64(time.Now().Unix()),
		Results:    s.GetResults(),
	}
}

// RouteStream is used to route the given stream to the given peer.
func (s *Session) RouteStream(stream network.Stream, n node.CallbackImpl) (*motor.OnTransmitCompleteResponse, error) {
	// Initialize Params
	logger.Debugf("Beginning %s Transmit Stream", s.Direction.String())
	doneChan := make(chan bool)

	// Check for Incoming
	if s.IsIn() {
		// Handle incoming stream
		go func(stream network.Stream, dchan chan bool) {
			// Create reader
			rs := msgio.NewReader(stream)

			// Read all items
			for _, v := range s.GetItems() {
				// Read Stream to File
				if err := v.ReadFromStream(n, rs); err != nil {
					logger.Errorf("Error reading stream: %v", err)
					dchan <- false
				} else {
					dchan <- true
				}
			}

			// Close Stream on Done Reading
			stream.Close()
		}(stream, doneChan)
	}

	// Check for Outgoing
	if s.IsOut() {
		// Handle outgoing stream
		go func(stream network.Stream, dchan chan bool) {
			// Create writer
			wc := msgio.NewWriter(stream)

			// Write all items
			for _, v := range s.GetItems() {
				// Write File to Stream
				if err := v.WriteToStream(n, wc); err != nil {
					logger.Errorf("Error writing file: %v", err)
					dchan <- false
				} else {
					dchan <- true
				}
			}
		}(stream, doneChan)
	}

	// Wait for all files to be written
	for {
		select {
		case r := <-doneChan:
			// Set Result
			if complete := s.UpdateCurrent(r); !complete {
				continue
			} else {
				return s.Event(), nil
			}
		}
	}
}

// UpdateCurrent updates the current index of the session.
func (s *Session) UpdateCurrent(result bool) bool {
	logger.Debugf("Item (%v) transmit result: %v", s.GetCurrentIndex(), result)
	s.Results[s.GetCurrentIndex()] = result
	s.CurrentIndex = s.GetCurrentIndex() + 1
	return int(s.GetCurrentIndex()) >= s.FinalIndex()
}

// NewSessionPayload creates session payload
func NewSessionPayload(p *common.Payload) *SessionPayload {
	return &SessionPayload{
		Payload: p,
	}
}

// CreateItems creates list of sessionItems
func (sp *SessionPayload) CreateItems(dir common.Direction) []*SessionItem {
	// Initialize Properties
	count := len(sp.GetPayload().GetItems())
	items := make([]*SessionItem, 0)

	// Iterate over items
	for i, v := range sp.GetPayload().GetItems() {
		// Get default payload item properties
		fi := v.GetFile()
		path := fi.GetPath()

		// Set Path for Incoming
		if dir == common.Direction_DIRECTION_INCOMING {
			inpath, err := fi.SetPathFromFolder(device.Downloads)
			if err == nil {
				path = inpath
			} else {
				logger.Errorf("%s - Failed to generate path for file: %s", err, fi.Name)
			}
		}

		// Create Session Item
		item := &SessionItem{
			Item:      fi,
			Index:     int32(i),
			TotalSize: sp.GetPayload().GetSize(),
			Size:      fi.GetSize(),
			Count:     int32(count),
			Direction: dir,
			Written:   0,
			Path:      path,
		}
		items = append(items, item)
	}
	return items
}

// ReadFromStream reads the item from the stream
func (si *SessionItem) ReadFromStream(node node.CallbackImpl, reader msgio.ReadCloser) error {
	// Create New File
	dst, err := os.Create(si.GetPath())
	defer dst.Close()
	if err != nil {
		return err
	}

	// Route Data from Stream
	for {
		// Read Next Message
		buf, err := reader.ReadMsg()
		if buf == nil {
			logger.Debug("Completed reading from stream: " + si.GetPath())
			return nil
		}

		if err != nil {
			if err == io.EOF {
				logger.Debug("Completed reading from stream: " + si.GetPath())
				return nil
			}
			return err
		}

		// Write Chunk to File
		n, err := dst.Write(buf)
		if err != nil {
			logger.Errorf("%s - Failed to Write Buffer to File on Read Stream", err)
			return err
		}

		// Update Progress
		if done := si.Progress(n, node); done {
			return nil
		}
	}
}

// WriteToStream writes the item to the stream
func (si *SessionItem) WriteToStream(node node.CallbackImpl, writer msgio.WriteCloser) error {
	// Create New Chunker
	f, err := os.Open(si.GetPath())
	defer f.Close()
	if err != nil {
		return err
	}

	// Create New Reader
	r := bufio.NewReader(f)
	buf := make([]byte, 0, 4*1024)

	// Loop through File
	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				logger.Debug("Completed writing from stream: " + si.GetPath())
				return nil
			}
			return err
		}

		// process buf
		if err != nil && err != io.EOF {
			return err
		}

		// Write Message Bytes to Stream
		err = writer.WriteMsg(buf)
		if err != nil {
			logger.Errorf("%s - Error Writing data to msgio.Writer", err)
			return err
		}

		// Update Progress
		si.Progress(len(buf), node)
	}
}

// Progress pushes a progress event to the node. Returns true if the item is done.
func (si *SessionItem) Progress(wrt int, n node.CallbackImpl) bool {
	// Update Progress
	si.Written += int64(wrt)

	// Create Progress Event
	if (si.GetWritten() % ITEM_INTERVAL) == 0 {
		event := &motor.OnTransmitProgressResponse{
			Direction: si.GetDirection(),
			Progress:  (float64(si.GetWritten()) / float64(si.GetTotalSize())),
			Current:   int32(si.GetIndex()) + 1,
			Total:     int32(si.GetCount()),
		}

		// Push ProgressEvent to Emitter
		go n.OnProgress(event)
	}

	// Return if Done
	return si.GetWritten() >= si.GetSize()
}
