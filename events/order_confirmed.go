package events

import (
	"time"

	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/models"
	"github.com/google/uuid"
)

// OrderConfirmed represents an order that was received in the system and will be published to our messaging system
type OrderConfirmed struct {
	EventBase BaseEvent
	EventBody models.Order
}

// ID returns the unique identifier of the event
func (or OrderConfirmed) ID() uuid.UUID {
	return or.EventBase.EventID
}

// Name returns the name of the event
func (or OrderConfirmed) Name() string {
	return "OrderConfirmed"
}

// Timestamp returns the unique timestamp of the event
func (or OrderConfirmed) Timestamp() time.Time {
	return or.EventBase.EventTimestamp
}

// Body returns the body content of the event
func (or OrderConfirmed) Body() interface{} {
	return or.EventBody
}
