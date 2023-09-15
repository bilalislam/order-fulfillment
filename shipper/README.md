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
    $> $KAFKA_HOME/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic OrderPickedAndPacked
    ```

## Running the Service
1. The program is written using Go modules, so you will need to ensure modules are turned on: https://blog.golang.org/using-go-modules

1. Navigate to the directory containing the _code_ 
    ```shell
    $> cd Asynchronous-Event-Handling-Using-Microservices-and-Kafka//code/shipper
    ```

1. Start the service
    ```shell
    $> go run main.go
    ```


## Testing the Service
1. Start a consumer for the Topic
    ```shell
    $> $KAFKA_HOME/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic OrderPickedAndPacked --from-beginning
    ```

1. Start a console producer for the Topic and post an *OrderPickedAndPacked* event
    ```shell
    $> $KAFKA_HOME/bin/kafka-console-producer.sh --bootstrap-server localhost:9092 --topic OrderPickedAndPacked
    > {"EventBase":{"EventID":"4a651ef8-a851-4d77-a58b-3d8af748a570","EventTimestamp":"2020-08-16T16:03:05.258542-04:00"},"EventBody":{"id":"c6b37316-b4da-4b25-94c8-14c08bad95e6","products":[{"productCode":"12345","quantity":2}],"customer":{"firstName":"Tom","lastName":"Hardy","emailAddress":"tom.hardy@email.com","shippingAddress":{"line1":"123 Anywhere St","city":"Anytown","state":"AL","postalCode":"12345"}}}}
    ```
    1. Verify the message was received by the topic consumer in the first testing step

    1. Verify an event reached the *Shipper* consumer. You should see a response logged from the *Shipper* consumer

