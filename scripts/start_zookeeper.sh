#!/bin/bash

# Set this to the location where Kafka has been installed
KAFKA_HOME=~/devops/kafka-3.5.1-src

# Start Zookeeper
$KAFKA_HOME/bin/zookeeper-server-start.sh $KAFKA_HOME/config/zookeeper.properties