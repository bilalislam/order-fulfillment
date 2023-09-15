package consumer

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/config"
	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/db"
	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/events"
	hdlr "github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/handlers"
	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/inventory/internal/handlers"
	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/models"
	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/publisher"
	"github.com/google/uuid"
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
		"group.id":              c.Group + "-inventory",
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "earliest"})

	if err != nil {
		log.WithField("error", err).Error("Failed to create consumer")

		return err
	}

	log.WithField("consumer", kc).Info("Created Consumer")

	if err = kc.SubscribeTopics([]string{c.Topic}, nil); err != nil {
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

		var event events.OrderReceived
		if err = json.Unmarshal([]byte(string(msg.Value)), &event); err != nil {
			log.WithField("error", err).Error("an issue occurred unmarshalling event from message received")

			continue
		}

		var order models.Order
		if order, err = extractOrder(event); err != nil {
			log.WithField("error", err).Error("an issue occurred trying to extract order information from the order recieved event")

			hdlr.HandleError(event)
			continue
		}

		if err = processEvent(event, order); err != nil {
			log.WithField("error", err).Error("an issue occurred trying to process the event")

			hdlr.HandleError(event)
			continue
		}

		if err = publishOrderConfirmedEvent(order); err != nil {
			log.WithField("error", err).Error("an issue occurred trying to publish an order confirmed event")

			hdlr.HandleError(event)
			continue
		}
	}
}

func extractOrder(event events.OrderReceived) (models.Order, error) {
	log.Info("attempting to extract order from event")

	body := event.Body()
	order, ok := body.(models.Order)
	if !ok {
		return models.Order{}, errors.New("event body can't be cast as an order")
	}

	return order, nil
}

func processEvent(event events.Event, order models.Order) error {
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

	// event hasn't been processed yet, decrement the inventory
	if err = handlers.DecrementInventory(order); err != nil {
		log.WithField("error", err).Error("an issue occurred trying to decrement the inventory")

		return err
	}

	// mark the event as processed
	if err = db.InsertEvent(event, tx); err != nil {
		log.WithField("error", err).Error("an issue occurred trying to insert the event")
		return err
	}

	return nil
}

func publishOrderConfirmedEvent(o models.Order) error {
	// publish an order confirmed event
	e := translateOrderToEvent(o)

	log.WithField("event", e).Info("transformed order to event")

	var err error
	if err = publisher.PublishEvent(e, config.OrderConfirmedTopicName); err != nil {
		return err
	}

	log.WithField("event", e).Info("published event")

	return nil
}

func translateOrderToEvent(o models.Order) events.Event {
	var event = events.OrderConfirmed{
		EventBase: events.BaseEvent{
			EventID:        uuid.New(),
			EventTimestamp: time.Now(),
		},
		EventBody: o,
	}

	return event
}
