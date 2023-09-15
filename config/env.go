package config

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

const (
	// LogLevelEnvVar is the name of the environment variable that controls
	// the log level of the application logger
	LogLevelEnvVar = "LOG_LEVEL"

	// PortEnvVar is the name of the environment variable that controls the
	// value of the port the service should listen on
	PortEnvVar = "PORT"

	// BrokerAddressEnvVar is the name of the environment variable that controls the
	// value of the kafka broker address
	BrokerAddressEnvVar = "BROKER_ADDRESS"

	// ConsumerGroupEnvVar is the name of the environment variable that controls the
	// value of the kafka consumer group name
	ConsumerGroupEnvVar = "CONSUMER_GROUP"

	// DatabaseAddressEnvVar is the name of the environment variable that controls the
	// value of the database address <host>:<port>
	DatabaseAddressEnvVar = "DB_ADDRESS"

	// DatabaseUsernameEnvVar is the name of the environment variable that controls the
	// value of the database username
	DatabaseUsernameEnvVar = "DB_USERNAME"

	// DatabasePasswordEnvVar is the name of the environment variable that controls the
	// value of the database password
	DatabasePasswordEnvVar = "DB_PASSWORD"

	// DatabaseNameEnvVar is the name of the environment variable that controls the
	// value of the database name
	DatabaseNameEnvVar = "DB_NAME"

	defaultLogLevel         = logrus.DebugLevel     // used if LOG_LEVEL not set
	defaultPort             = 8080                  // used if PORT not set
	defaultBrokerAddress    = "localhost"           // used if BROKER_ADDRESS not set
	defaultConsumerGroup    = "test-consumer-group" // used if CONSUMER_GROUP not set
	defaultDatabaseAddress  = "localhost:5432"      // used if DB_ADDRESS not set
	defaultDatabaseUsername = "postgres"            // used if DB_USERNAME not set
	defaultDatabasePassword = "postgres"            // used if DB_PASSWORD not set
	defaultDatabaseName     = "liveproject"         // used if DB_NAME not set
)

// LogLevel returns the log level set in the environment, or debug if not defined
func LogLevel() logrus.Level {
	var (
		level logrus.Level
		err   error
	)

	if level, err = logrus.ParseLevel(os.Getenv(LogLevelEnvVar)); err != nil {
		return defaultLogLevel
	}

	return level
}

// Port returns the port the service should listen on, or default value if not defined or
// is not a valid port
func Port() int {
	var (
		rawPort string
		found   bool
		port    int
		err     error
	)

	if rawPort, found = os.LookupEnv(PortEnvVar); !found {
		return defaultPort
	}

	if port, err = strconv.Atoi(rawPort); err != nil {
		return defaultPort
	}

	return port
}

// BrokerAddress returns the address the kafka broker is listening on, or default value if not defined
func BrokerAddress() string {
	return value(BrokerAddressEnvVar, defaultBrokerAddress)
}

// ConsumerGroup returns the consumer group name to use for the consumer, or default value if not defined
func ConsumerGroup() string {
	return value(ConsumerGroupEnvVar, defaultConsumerGroup)
}

// DatabaseAddress returns the database address to use, or default value if not defined
func DatabaseAddress() string {
	return value(DatabaseAddressEnvVar, defaultDatabaseAddress)
}

// DatabaseUsername returns the database username to use, or default value if not defined
func DatabaseUsername() string {
	return value(DatabaseUsernameEnvVar, defaultDatabaseUsername)
}

// DatabasePassword returns the database password to use, or default value if not defined
func DatabasePassword() string {
	return value(DatabasePasswordEnvVar, defaultDatabasePassword)
}

// DatabaseName returns the database name to use, or default value if not defined
func DatabaseName() string {
	return value(DatabaseNameEnvVar, defaultDatabaseName)
}

func value(key, defaultValue string) string {
	var value string
	var found bool

	if value, found = os.LookupEnv(key); !found {
		return defaultValue
	}

	if len(value) == 0 {
		return defaultValue
	}

	return value
}
