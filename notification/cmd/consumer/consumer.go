package consumer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/db"
	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/events"
	hdlr "github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/handlers"
	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/models"
	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/notification/internal/handlers"
	log "github.com/sirupsen/logrus"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Consumer represents the subscription to a specified Kafka topic
type Consumer struct {
	Broker string
	Group  string
	Topic  string
}

// SubscribeAndListen will subscribe to a Kafka topic and start polling and listening for events
// Adpated from https://github.com/confluentinc/confluent-kafka-go#examples
func (c *Consumer) SubscribeAndListen() error {

	kc, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     c.Broker,
		"broker.address.family": "v4",
		"group.id":              c.Group + "-notification",
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "earliest"})

	if err != nil {
		log.WithField("error", err).Error("Failed to create consumer")

		return err
	}

	log.WithField("consumer", kc).Info("Created Consumer")

	err = kc.SubscribeTopics([]string{c.Topic}, nil)
	if err != nil {
		log.WithField("error", err).
			WithField("topic", c.Topic).
			Error("Failed to subscribe to topic")

		return err
	}

	for {
		msg, err := kc.ReadMessage(-1)
		if err != nil {
			// The client will automatically try to recover from all errors.
			log.WithField("error", err).Error(msg)

			log.Warn("Closing consumer...")
			kc.Close()

			return err
		}

		log.WithField("topic", msg.TopicPartition).Info(string(msg.Value))

		var event events.Notification
		if err = json.Unmarshal([]byte(string(msg.Value)), &event); err != nil {
			log.WithField("error", err).Error("an issue occurred unmarshalling event from message received")

			continue
		}

		notification, err := extractNotification(event)
		if err != nil {
			log.WithField("error", err).Error("an issue occurred trying to extract notification information from the notification event")

			hdlr.HandleError(event)
			continue
		}

		if err = processEvent(event, notification); err != nil {
			log.WithField("error", err).Error("an issue occurred trying to process the event")

			hdlr.HandleError(event)
			continue
		}
	}
}

func extractNotification(event events.Notification) (models.Notification, error) {
	log.Info("attempting to extract notification from event")

	body := event.Body()
	notification, ok := body.(models.Notification)
	if !ok {
		return models.Notification{}, errors.New("event body can't be cast as a Notification")
	}

	return notification, nil
}

func processEvent(event events.Event, notification models.Notification) error {
	var err error

	db := db.NewDB()
	conn, err := db.Connect()
	if err != nil {
		log.WithField("error", err).Error("an issue occurred trying to make a connection to the database")
		return err
	}

	// begin a transaction
	tx, err := conn.Begin(context.Background())
	if err != nil {
		log.WithField("error", err).Error("an issue occurred trying to start a database transaction")
		return err
	}

	defer func() {
		log.Info("committing DB transaction")
		if err = tx.Commit(context.Background()); err != nil {
			log.WithField("error", err).Error("an issue occurred trying to commit the transaction")
		}

		log.Info("closing connection to database")
		conn.Close(context.Background())
	}()

	// check to see if event has already been processed
	eventAlreadyProcessed, err := db.EventExists(event, tx)
	if err != nil {
		log.WithField("error", err).Error("an issue occurred trying to check if an event was already processed")
		return err
	}

	// if event has already been processed, nothing more to do
	if eventAlreadyProcessed {
		log.WithField("event.id", event.ID()).
			WithField("event.name", event.Name()).
			Info("event was processed previously")

		return nil
	}

	// event hasn't been processed yet, send the notification
	switch notification.Type {
	case models.Email:
		if err = handlers.SendEmail(notification); err != nil {
			log.WithField("error", err).Error("an issue occurred trying to send an email to the customer")

			return err
		}
	default:
		log.WithField("notification.type", notification.Type).Error("notification type is not supported at this time")

		return fmt.Errorf("notification type, \"%s\" is not supported", notification.Type)
	}

	// mark the event as processed
	if err = db.InsertEvent(event, tx); err != nil {
		log.WithField("error", err).Error("an issue occurred trying to insert the event")
		return err
	}

	return nil
}
