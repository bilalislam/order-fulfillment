#!/bin/bash

# Set this to the location where Kafka has been installed
KAFKA_HOME=~/devops/kafka-3.5.1-src

# Create the OrderReceived topic
$KAFKA_HOME/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic OrderReceived

# Create the OrderConfirmed topic
$KAFKA_HOME/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic OrderConfirmed

# Create the OrderPickedAndPacked topic
$KAFKA_HOME/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic OrderPickedAndPacked

# Create the Notification topic
$KAFKA_HOME/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic Notification

# Create the DeadLetterQueue topic
$KAFKA_HOME/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic DeadLetterQueue