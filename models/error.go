package models

import "time"

// Error represents an error that has occurred in the system
type Error struct {
	Message   string
	Timestamp time.Time
}

// NewError will create a new Error object from a base error
func NewError(err error) Error {
	return Error{
		Message:   err.Error(),
		Timestamp: time.Now(),
	}
}
