# Implementation Notes

## Running Kafka and Setting up the Topic
1. Start Zookeeper
    ```shell
    $> $KAFKA_HOME/bin/zookeeper-server-start.sh config/zookeeper.properties
    ```

1. Start Kafka
    ```shell
    $> $KAFKA_HOME/bin/kafka-server-start.sh config/server.properties
    ```

1. Create the Topic
    ```shell
    $> $KAFKA_HOME/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic Notification
    ```

## Running the Service
1. The program is written using Go modules, so you will need to ensure modules are turned on: https://blog.golang.org/using-go-modules

1. Navigate to the directory containing the _code_ 
    ```shell
    $> cd Asynchronous-Event-Handling-Using-Microservices-and-Kafka//code/notification/
    ```

1. Start the service
    ```shell
    $> go run main.go
    ```

## Testing the Service
1. Publish an event to the *Notification* topic
    ```shell
    $> $KAFKA_HOME/bin/kafka-console-producer.sh --bootstrap-server localhost:9092 --topic Notification
    > {"EventBase":{"EventID":"d6c54946-0bdc-4d63-898f-447ccdb7ccf4","EventTimestamp":"2020-08-16T09:01:59.757509-04:00"},"EventBody":{"type":"email","recipient":"tom.hardy@email.com","from":"orders@ppe4all.com","subject":"Hi Tom, your order has been received!","body":"<p>We have received your order and it is being fulfilled!</p>"}}
    ```

1. Verify an event reached the *Notification* service. You should see a response logged from the *Notification* service

