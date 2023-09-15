package events

import (
	"time"

	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/metrics"
	"github.com/google/uuid"
)

// OrderCountMetric represents an order metric that represents an order was recieved and how many products were in the order
type OrderCountMetric struct {
	EventBase BaseEvent
	EventBody metrics.OrderCount
}

// ID returns the unique identifier of the event
func (ocm OrderCountMetric) ID() uuid.UUID {
	return ocm.EventBase.EventID
}

// Name returns the name of the event
func (ocm OrderCountMetric) Name() string {
	return "OrderCountMetric"
}

// Timestamp returns the unique timestamp of the event
func (ocm OrderCountMetric) Timestamp() time.Time {
	return ocm.EventBase.EventTimestamp
}

// Body returns the body content of the event
func (ocm OrderCountMetric) Body() interface{} {
	return ocm.EventBody
}

// TranslateToOrderCountMetricEvent will translate an order count metric to to an event
func TranslateToOrderCountMetricEvent(oc metrics.OrderCount) Event {
	var event = OrderCountMetric{
		EventBase: BaseEvent{
			EventID:        uuid.New(),
			EventTimestamp: time.Now(),
		},
		EventBody: oc,
	}

	return event
}
