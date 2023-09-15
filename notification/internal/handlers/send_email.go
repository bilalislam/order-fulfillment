package handlers

import (
	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/models"
	log "github.com/sirupsen/logrus"
)

// SendEmail will construct and send an email based on the supplied notification information
func SendEmail(notification models.Notification) error {
	log.Info("attempting to send an email notification")

	log.WithField("subject", notification.Subject).
		WithField("body", notification.Body).
		WithField("recipient", notification.Recipient).
		WithField("from", notification.From).
		Info("would send email here")

	return nil
}
