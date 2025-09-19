A platform for real-time data handling – primarily through a data stream-processing engine and a distributed event store


### Kafka Architecture

Kafka’s architecture consists of several key components:

a. Topics
    • Topics are categories or streams of data where messages are sent.
    • Each topic is split into partitions for parallelism.
    • Messages in a topic are ordered within a partition.

b. Producers
    • Producers publish messages to specific topics.
    • Messages can be sent to a specific partition or distributed across partitions (using keys or round-robin).

c. Consumers
    • Consumers subscribe to topics and read messages.
    • Consumer groups allow multiple consumers to share the load of a topic’s partitions.

e. Partitions
    • Each topic has multiple partitions, allowing messages to be distributed for scalability.
    • Partitions are replicated across brokers for fault tolerance.

f. Zookeeper
    • Manages cluster metadata, leader election, and configurations.

Note: Kafka is moving toward replacing Zookeeper with its own metadata quorum in newer versions.)


###  Kafka Use Cases

Kafka is used in a variety of domains. Some common use cases include:

    1. Log Aggregation: Centralize logs from multiple services or servers.

    2. Real-Time Analytics: Process data streams for insights in real-time (e.g., user behavior analysis).

    3. Data Pipelines: Ingest, process, and store data across systems (ETL pipelines).

    4. Event Sourcing: Store and replay event histories for debugging or system reconstruction.

    5. Metrics and Monitoring: Collect metrics and system telemetry in real time.

    6. Stream Processing: Use Kafka Streams API to transform and aggregate real-time data.


### Key Features of Kafka

    1. Distributed: Kafka runs as a cluster of nodes (brokers) and handles massive volumes of data.

    2. High Performance: Kafka can handle millions of messages per second with minimal latency.

    3. Durability: Messages are persisted on disk, ensuring fault tolerance.

    4. Scalable: Kafka can be scaled horizontally by adding more brokers or partitions.

    5. Pub-Sub Messaging: Supports producers (publishers) and consumers (subscribers).

    6. Replayable Logs: Consumers can replay messages by reading from a specific offset in the log.

    7. Streaming API: Provides real-time processing through the Kafka Streams API.

### Kafka Workflow Example

Here’s a simple Kafka workflow:

    1. A producer sends messages to a topic.
    2. Kafka brokers store these messages in partitions.
    3. Consumers subscribe to the topic and read messages, either in real-time or from a specific offset.
    4. Messages are processed by downstream applications.


### Implementing Kafka in Golang

Let’s create a producer and consumer in Golang using the Confluent Kafka Go library.

Step 1: Install the Kafka Library

    go get github.com/confluentinc/confluent-kafka-go/kafka


Step 2: Create a Kafka Producer

    package main
    import (
        "log"
        "os"
    "github.com/confluentinc/confluent-kafka-go/kafka"
    )
    func main() {
        // Create a Kafka producer
        producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
        if err != nil {
            log.Fatalf("Failed to create producer: %s", err)
        }
        defer producer.Close()
    // Get the topic from the command line
        topic := "my-topic"
    // Send a message
        message := "Hello Kafka!"
        err = producer.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
            Value:          []byte(message),
        }, nil)
        if err != nil {
            log.Printf("Failed to produce message: %s", err)
        } else {
            log.Printf("Produced message: %s", message)
        }
    // Wait for delivery reports
        producer.Flush(15 * 1000)
    }

Step 3: Create a Kafka Consumer

    package main
    import (
        "log"
    "github.com/confluentinc/confluent-kafka-go/kafka"
    )
    func main() {
        // Create a Kafka consumer
        consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
            "bootstrap.servers": "localhost:9092",
            "group.id":          "my-group",
            "auto.offset.reset": "earliest",
        })
        if err != nil {
            log.Fatalf("Failed to create consumer: %s", err)
        }
        defer consumer.Close()
    // Subscribe to a topic
        topic := "my-topic"
        consumer.SubscribeTopics([]string{topic}, nil)
    // Poll messages
        for {
            msg, err := consumer.ReadMessage(-1)
            if err == nil {
                log.Printf("Received message: %s\n", string(msg.Value))
            } else {
                log.Printf("Error while consuming message: %v\n", err)
            }
        }
    }


### Common Kafka Questions

    1. What are Kafka partitions? How do they improve performance?
        ○ Partitions allow parallel processing by distributing data across multiple brokers.
        
    2. How does Kafka handle fault tolerance?
        ○ By replicating partitions across brokers. The leader handles reads/writes, while replicas sync.
        
    3. What is the role of Zookeeper?
        ○ Metadata management, leader election, and configuration storage. (Soon being replaced by Kafka’s KRaft mode.)
        
    4. How does Kafka ensure message ordering?
        ○ Messages are ordered within a partition but not across partitions. Use keys to control partitioning.
        
    5. Difference between Kafka Streams and Kafka Consumer API?
        ○ Kafka Streams is for data processing (transform, aggregate).
        ○ Kafka Consumer API is for reading messages from Kafka topics.

