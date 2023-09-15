package events

import (
	"time"

	"github.com/google/uuid"
)

// Error represents an error that has occurred when trying to process an event in the system and will be published to our messaging system
type Error struct {
	EventBase BaseEvent
	EventBody Event
}

// ID returns the unique identifier of the event
func (e Error) ID() uuid.UUID {
	return e.EventBase.EventID
}

// Name returns the name of the event
func (e Error) Name() string {
	return "Error"
}

// Timestamp returns the unique timestamp of the event
func (e Error) Timestamp() time.Time {
	return e.EventBase.EventTimestamp
}

// Body returns the body content of the event
func (e Error) Body() interface{} {
	return e.EventBody
}
