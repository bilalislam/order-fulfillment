# Overview of Order Fulfillemnt Implementation 

# What Was Created?
In this milestone:
1. two KPIs were defined and published to two new Kafka topics
    1. A [Metric](./metrics/order_count.go) that tracks each time an order is placed and how many products were ordered is defined to expose information which could be used to answer things like:
        * how the system is performing (i.e. 0 orders in an hour might be indicative of a larger problem)
        * time of day customers tend to shop
        * how many orders per (second/minute/hours etc)
        * average number of products per order
    1. A [Metric](./metrics/order_time.go) that can be used to track the total time it takes to process an order, from when it was received by the warehouse and when it was picked and packed. This can be used to answer things like:
        * average time it takes to pick and pack an order
        * what times of day are slower than others

# How to Test?
I was able to test all of the code created in this milestone on my local machine. The instructions below assume you are running on your local machine. I implemented this on a Mac, so references to the command-line will show as a UNIX shell.

1. Kafka and Zookeeper need to be running
    1. The *OrderReceived* topic should be created
    1. The *OrderPickedAndPacked* topic should be created
    1. The *Notification* topic should be created
    1. The *DeadLetterQueue* topic should be created
        1. All the topics need to be created: [click here for more information](./scripts/create_topics.sh)
        1. Start kafka: [click here for more information](./scripts/start_kafka.sh)
        1. Start zookeeper: [click here for more information](./scripts/start_zookeeper.sh)
    1. A new topic *OrderCount* is created
        ```shell
        $ $KAFKA_HOME/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic OrderCount
        ```
    1. A new topic *OrderTime* is created
        ```shell
        $ $KAFKA_HOME/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic OrderTime
        ```
1. The Postgress database needs to be running
    1. The right database, schema and table need to be created: [click here for more information](./db/db.go)
1. The *Order* service needs to be running (assumes you are in the `/code` folder)
    1. If this is the first time you are running this code, you will need to setup Go modules
        1. In your `~/.bash_profile`, make sure you have the following ENV var set: `export GO111MODULE=on` and make sure the file is sourced.
        1. Initialize Go modules
            ```shell
            $ go mod init
            ```
            ```shell
            $ go mod tidy
            ```
    1. Run the *Order* service
        ```shell
        $ go run order/main.go
        ```
1. The *Inventory* consumer needs to be running (assumes you are in the `/code` folder)
    1. If this is the first time you are running this code, you will need to setup Go modules
        1. In your `~/.bash_profile`, make sure you have the following ENV var set: `export GO111MODULE=on` and make sure the file is sourced.
        1. Initialize Go modules
            ```shell
            $ go mod init
            ```
            ```shell
            $ go mod tidy
            ```
    1. Run the *Inventory* consumer service
        ```shell
        $ go run inventory/main.go
        ```
1. The *Warehouse* consumer needs to be running (assumes you are in the `/code` folder)
    1. If this is the first time you are running this code, you will need to setup Go modules
        1. In your `~/.bash_profile`, make sure you have the following ENV var set: `export GO111MODULE=on` and make sure the file is sourced.
        1. Initialize Go modules
            ```shell
            $ go mod init
            ```
            ```shell
            $ go mod tidy
            ```
    1. Run the *Warehouse* consumer service
        ```shell
        $ go run warehouse/main.go
        ```
1. The *Notification* consumer needs to be running (assumes you are in the `/code` folder)
    1. If this is the first time you are running this code, you will need to setup Go modules
        1. In your `~/.bash_profile`, make sure you have the following ENV var set: `export GO111MODULE=on` and make sure the file is sourced.
        1. Initialize Go modules
            ```shell
            $ go mod init
            ```
            ```shell
            $ go mod tidy
            ```
    1. Run the *Notification* consumer service
        ```shell
        $ go run notification/main.go
        ```
1. The *Shipper* consumer needs to be running (assumes you are in the `/code` folder)
    1. If this is the first time you are running this code, you will need to setup Go modules
        1. In your `~/.bash_profile`, make sure you have the following ENV var set: `export GO111MODULE=on` and make sure the file is sourced.
        1. Initialize Go modules
            ```shell
            $ go mod init
            ```
            ```shell
            $ go mod tidy
            ```
    1. Run the *Shipper* consumer service
        ```shell
        $ go run shipper/main.go
        ```
1. Send a HTTP request to the order service:
    ```shell
    $ curl -v -H "Content-Type: application/json" -d '{"id":"6e042f29-350b-4d51-8849-5e36456dfa48","products":[{"productCode":"12345","quantity":2}],"customer":{"firstName":"Tom","lastName":"Hardy","emailAddress":"tom.hardy@email.com","shippingAddress":{"line1":"123 Anywhere St","city":"Anytown","state":"AL","postalCode":"12345"}}}' http://localhost:8080/orders
    ```
1. You should see output in the console of the order service, and no errors. You can also check the contents of the *OrderReceived* topic in Kafka.
    ```shell
    $ $KAFKA_HOME/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic OrderReceived --from-beginning
    ```
1. You can also check the contents of the *Notification* topic in Kafka.
    ```shell
    $ $KAFKA_HOME/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic Notification --from-beginning
    ```
1. You should see output in the console of the inventory consumer, and no errors.
1. You should see output in the console of the warehouse consumer, and no errors.
1. You should see output in the console of the notification consumer, and no errors.
1. To check the the shipping consumer, you will need to manually publish a message to the *OrderPickedAndPacked* topic in Kafka.
    ```shell
    $ $KAFKA_HOME/bin/kafka-console-producer.sh --bootstrap-server localhost:9092 --topic OrderPickedAndPacked
    ```
    Then you can paste a payload like this (though it should correlate to the same order being tested with in previous steps)
    ```json
    {"EventBase":{"EventID":"4a651ef8-a851-4d77-a58b-3d8af748a570","EventTimestamp":"2020-08-16T16:03:05.258542-04:00"},"EventBody":{"id":"c6b37316-b4da-4b25-94c8-14c08bad95e6","products":[{"productCode":"12345","quantity":2}],"customer":{"firstName":"Tom","lastName":"Hardy","emailAddress":"tom.hardy@email.com","shippingAddress":{"line1":"123 Anywhere St","city":"Anytown","state":"AL","postalCode":"12345"}}}}
    ```
1. You should see output in the console of the shipper consumer, and no errors.

# Project Conclusions

After all milestones are complete, you will have a collection of microservices and consumers that can independently scale and evolve over time, are loosely coupled, and are highly cohesive. This outcome is made possible by the introduction of a distributed messaging system like Kafka. Using Kafka adds fault tolerance and durable storage of events leveraged in the system. When used together, these components create an asynchronous, event-driven architecture that makes your system adaptable, scalable, and resilient over time. You should be able to take the lessons learned here and apply the concepts to future projects that participate within a distributed system.

## Review what you have learned and accomplished:

1. Interacting with Kafka by using provided command-line tools
1. Creating and altering Kafka topics
1. Publishing events to, and consuming events from, Kafka topics through the command line and programmatically through Go
1. Using asynchronous communication patterns across independent services and consumers
1. Handling situations when events can’t be processed
1. Evaluating performance of the system through the lens of the business by defining and publishing KPIs 