package events

import (
	"time"

	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/models"
	"github.com/google/uuid"
)

// OrderPickedAndPacked represents an event when customers order has been picked and packed by the warehouse
type OrderPickedAndPacked struct {
	EventBase BaseEvent
	EventBody models.Order
}

// ID returns the unique identifier of the event
func (n OrderPickedAndPacked) ID() uuid.UUID {
	return n.EventBase.EventID
}

// Name returns the name of the event
func (n OrderPickedAndPacked) Name() string {
	return "OrderPickedAndPacked"
}

// Timestamp returns the unique timestamp of the event
func (n OrderPickedAndPacked) Timestamp() time.Time {
	return n.EventBase.EventTimestamp
}

// Body returns the body content of the event
func (n OrderPickedAndPacked) Body() interface{} {
	return n.EventBody
}
