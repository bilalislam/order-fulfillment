package events

import (
	"time"

	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/metrics"
	"github.com/google/uuid"
)

// OrderTimeMetric represents an order metric that represents an order was recieved and how many products were in the order
type OrderTimeMetric struct {
	EventBase BaseEvent
	EventBody metrics.OrderTime
}

// ID returns the unique identifier of the event
func (otm OrderTimeMetric) ID() uuid.UUID {
	return otm.EventBase.EventID
}

// Name returns the name of the event
func (otm OrderTimeMetric) Name() string {
	return "OrderTimeMetric"
}

// Timestamp returns the unique timestamp of the event
func (otm OrderTimeMetric) Timestamp() time.Time {
	return otm.EventBase.EventTimestamp
}

// Body returns the body content of the event
func (otm OrderTimeMetric) Body() interface{} {
	return otm.EventBody
}

// TranslateToOrderTimeMetricEvent will translate an order time metric to to an event
func TranslateToOrderTimeMetricEvent(ot metrics.OrderTime) Event {
	var event = OrderTimeMetric{
		EventBase: BaseEvent{
			EventID:        uuid.New(),
			EventTimestamp: time.Now(),
		},
		EventBody: ot,
	}

	return event
}
