package models

import (
	"encoding/json"
	"errors"
)

// NotificationType the supported types of notifications
type NotificationType string

const (
	// Email represents the notification type of email
	Email = "email"
)

// UnmarshalJSON will unmarshall the notification type
func (nt *NotificationType) UnmarshalJSON(b []byte) error {
	var s string
	json.Unmarshal(b, &s)
	notificationType := NotificationType(s)
	switch notificationType {
	case Email:
		*nt = notificationType
		return nil
	}
	return errors.New("Invalid notification type")
}

// IsValid returns true if the notification type is valid
func (nt NotificationType) IsValid() error {
	switch nt {
	case Email:
		return nil
	}
	return errors.New("Invalid notification type")
}

// Notification represent a notification to be sent
type Notification struct {
	Type      NotificationType `json:"type"`
	Recipient string           `json:"recipient"`
	From      string           `json:"from"`
	Subject   string           `json:"subject"`
	Body      string           `json:"body"`
}
