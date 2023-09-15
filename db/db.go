package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/config"
	"github.com/bilalislam/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/code/events"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

// DB represents an interaction with a database
type DB struct {
	Address  string
	Username string
	Password string
	Database string
}

// NewDB returns a default localhost reference to a DB
func NewDB() DB {
	return DB{
		Address:  config.DatabaseAddress(),
		Username: config.DatabaseUsername(),
		Password: config.DatabasePassword(),
		Database: config.DatabaseName(),
	}
}

// EventExists will check to see if an event has already been processed
func (db DB) EventExists(event events.Event, tx pgx.Tx) (bool, error) {
	var id string
	var err error

	if err = tx.QueryRow(context.Background(), "select id from events.processed_events where id=$1 and event_name=$2", event.ID(), event.Name()).Scan(&id); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			log.WithField("code", pgErr.Code).
				WithField("message", pgErr.Message).
				Error("encountered an issue querying for the event")
			return false, err
		}
		return false, nil
	}

	return len(id) > 0, nil
}

// InsertEvent will insert a row into the processed_events table to indicate an event was processed.
func (db DB) InsertEvent(event events.Event, tx pgx.Tx) error {
	var err error

	if _, err = tx.Exec(context.Background(), "insert into events.processed_events (id, event_name, processed_timestamp) values ($1, $2, $3)", event.ID(), event.Name(), time.Now()); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			log.WithField("code", pgErr.Code).
				WithField("message", pgErr.Message).
				Error("encountered an issue inserting the event into the DB")
		}

		return err
	}

	return nil
}

// Connect creates a new conenction to the database
func (db DB) Connect() (*pgx.Conn, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s/%s", db.Username, db.Password, db.Address, db.Database)
	log.WithField("url", url).Info("attempting to connect to DB")

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}

	log.Info("successfully connected to DB")

	return conn, nil
}
