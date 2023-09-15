package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/config"
	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/events"
	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/models"
	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/publisher"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// ShipOrder will alert the customer the order is being shipped
func ShipOrder(order models.Order) error {
	log.WithField("order.id", order.ID).
		Info("attempting to alert the customer the order is being shipped")

	// notify the customer the order is being shipped
	var b strings.Builder
	for _, p := range order.Products {
		fmt.Fprintf(&b, "%d of product [%s]", p.Quantity, p.ProductCode)
	}

	address := fmt.Sprintf("<div>Shipping to Address:</div><div>%s</div><div>%s %s, %s</div>", order.Customer.ShippingAddress.Line1, order.Customer.ShippingAddress.City, order.Customer.ShippingAddress.State, order.Customer.ShippingAddress.PostalCode)
	subject := fmt.Sprintf("Hello %s, your order is being shipped!", order.Customer.FirstName)
	body := fmt.Sprintf("<div>Your order is on its way! Here is a review of the products in your order:</div><div>%s</div><div>%s</div>", b.String(), address)

	var err error
	event := events.Notification{
		EventBase: events.BaseEvent{
			EventID:        uuid.New(),
			EventTimestamp: time.Now(),
		},
		EventBody: models.Notification{
			Type:      models.Email,
			Recipient: order.Customer.EmailAddress,
			From:      "orders@ppe4all.com",
			Subject:   subject,
			Body:      body,
		},
	}

	if err = publisher.PublishEvent(event, config.NotificationTopicName); err != nil {
		log.WithField("error", err).
			WithField("topic", config.NotificationTopicName).
			Error("an issue ocurred publishing an event to Kafka")

		return err
	}

	return nil
}
