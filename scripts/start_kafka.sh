#!/bin/bash

# Set this to the location where Kafka has been installed
KAFKA_HOME=~/devops/kafka-3.5.1-src

# Start Kafka
$KAFKA_HOME/bin/kafka-server-start.sh $KAFKA_HOME/config/server.properties
