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
    $> $KAFKA_HOME/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic OrderReceived
    ```

## Running the Service
1. The program is written using Go modules, so you will need to ensure modules are turned on: https://blog.golang.org/using-go-modules

1. Navigate to the directory containing the _code_ 
    ```shell
    $> cd Asynchronous-Event-Handling-Using-Microservices-and-Kafka//code/warehouse
    ```

1. Start the service
    ```shell
    $> go run main.go
    ```


## Testing the Service
1. Start a consumer for the Topic
    ```shell
    $> $KAFKA_HOME/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic OrderReceived --from-beginning
    ```

1. Start the *Order* Service
    1. Navigate to the directory containing the _code_ 
        ```shell
        $> cd Asynchronous-Event-Handling-Using-Microservices-and-Kafka//code/order
        ```

    1. Start the service
        ```shell
        $> go run main.go
        ```

    1. Submit a payload to the service
        ```shell
        $> curl -v -H "Content-Type: application/json" -d '{"products":[{"productCode":"12345","quantity":2}],"customer":{"firstName":"Tom","lastName":"Hardy","emailAddress":"tom.hardy@email.com","shippingAddress":{"line1":"123 Anywhere St","city":"Anytown","state":"AL","postalCode":"12345"}}}' http://localhost:8080/orders
        ```

1. Verify an event reached the *Warehouse* consumer. You should see a response logged from the *Warehouse* consumer

