package channel

import (
	"context"
	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"google.golang.org/protobuf/proto"
)

// NewStore creates a new store
func NewStore(opts *options) *Store {
	// Create a new store
	return &Store{
		Data:     make(map[string]*StoreEntry),
		Capacity: int32(opts.capacity),
		Modified: time.Now().Unix(),
		Ttl:      opts.ttl.Milliseconds(),
	}
}

// Delete deletes an entry from the store and publishes an event
func (s *Store) Delete(key string, b *channel) error {
	// Fetch the entry
	entry := s.Data[key]
	if entry == nil {
		return ErrNotFound
	}

	// Check if the entry is owned by this node
	if entry.GetCreator() != b.n.HostID().String() {
		return ErrNotOwner
	}

	// Delete the entry
	delete(s.Data, key)
	s.Modified = time.Now().Unix()

	// Create Delete Event
	event := b.NewDeleteEvent(key)
	return event.Publish(b.ctx, b.topic)
}

// Get returns the value of the entry
func (s *Store) Get(key string) ([]byte, error) {

	entry := s.Data[key]
	if entry == nil {
		return nil, ErrNotFound
	}
	return entry.Value, nil
}

// Handle checks the event type and handles it with the store
func (s *Store) Handle(e *Event, b *channel) error {
	// Check if the event is valid
	if b.n.HostID().String() == e.GetCreator() {
		return nil
	}

	switch e.Type {
	case EventType_EVENT_TYPE_DELETE:
		delete(s.Data, e.Entry.Key)
		return nil
	case EventType_EVENT_TYPE_SYNC:
		if e.Store != nil {
			if s.Modified > e.Store.Modified && len(s.Data) < int(e.Store.Capacity) {
				s.Data = e.Store.Data
				s.Modified = e.Store.Modified
				logger.Debug("Store - Updated store to pushed earlier version")
			}
		}
		return nil
	case EventType_EVENT_TYPE_PUSH:
		if e.Entry != nil {
			s.Data[e.Entry.Key] = e.Entry
			s.Modified = time.Now().Unix()
			logger.Debug("Store - Added new Store Entry")
		}
		return nil
	case EventType_EVENT_TYPE_SET:
		if e.Entry != nil {
			s.Data[e.Entry.Key] = e.Entry
			s.Modified = time.Now().Unix()
			logger.Debug("Store - Set Updated Store Entry")
		}
		return nil
	}
	return nil
}

// Put puts an entry into the store and publishes an event
func (s *Store) Put(key string, value []byte, b *channel) error {
	// Fetch the entry
	entry := s.Data[key]
	if entry == nil {
		// Create new entry with Event
		event, entry := b.NewPushEvent(key, value)
		s.Data[key] = entry
		s.Modified = time.Now().Unix()
		return event.Publish(b.ctx, b.topic)
	}

	// Get existing entry and update it
	event, err := entry.Set(value, b.n.HostID().String())
	if err != nil {
		return err
	}
	s.Modified = time.Now().Unix()
	return event.Publish(b.ctx, b.topic)
}

// Set updates the entry in the store and publishes an event
func (se *StoreEntry) Set(value []byte, selfID string) (*Event, error) {
	if se.GetCreator() != selfID {
		return nil, ErrNotOwner
	}
	se.Value = value
	se.Modified = time.Now().Unix()
	return &Event{
		Type:    EventType_EVENT_TYPE_SET,
		Entry:   se,
		Creator: se.GetCreator(),
	}, nil
}

// NewPutEvent creates a new put event
func (b *channel) NewPushEvent(key string, value []byte) (*Event, *StoreEntry) {
	entry := &StoreEntry{
		Key:      key,
		Value:    value,
		Creator:  b.n.HostID().String(),
		Created:  time.Now().Unix(),
		Modified: time.Now().Unix(),
	}
	event := &Event{
		Type:    EventType_EVENT_TYPE_PUSH,
		Creator: b.n.HostID().String(),
		Entry:   entry,
	}
	return event, entry
}

// NewSyncEvent creates a new sync event
func (b *channel) NewSyncEvent() *Event {
	return &Event{
		Type:    EventType_EVENT_TYPE_SYNC,
		Creator: b.n.HostID().String(),
		Store:   b.store,
	}
}

// NewDeleteEvent creates a new delete event
func (b *channel) NewDeleteEvent(key string) *Event {
	entry := &StoreEntry{
		Key:      key,
		Creator:  b.n.HostID().String(),
		Modified: time.Now().Unix(),
	}
	event := &Event{
		Type:    EventType_EVENT_TYPE_DELETE,
		Creator: b.n.HostID().String(),
		Entry:   entry,
	}
	return event
}

// Marshal converts the event to a protobuf message and returns buffer
func (e *Event) Marshal() ([]byte, error) {
	return proto.Marshal(e)
}

// Publish publishes the event to the topic
func (e *Event) Publish(ctx context.Context, t *pubsub.Topic) error {
	buf, err := e.Marshal()
	if err != nil {
		return err
	}
	return t.Publish(ctx, buf)
}
