package events

import (
	"time"

	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/models"
	"github.com/google/uuid"
)

// OrderReceived represents an order that was received in the system and will be published to our messaging system
type OrderReceived struct {
	EventBase BaseEvent
	EventBody models.Order
}

// ID returns the unique identifier of the event
func (or OrderReceived) ID() uuid.UUID {
	return or.EventBase.EventID
}

// Name returns the name of the event
func (or OrderReceived) Name() string {
	return "OrderReceived"
}

// Timestamp returns the unique timestamp of the event
func (or OrderReceived) Timestamp() time.Time {
	return or.EventBase.EventTimestamp
}

// Body returns the body content of the event
func (or OrderReceived) Body() interface{} {
	return or.EventBody
}
