package main

import (
	"log"

	"github.com/IBM/sarama"
)

func main() {
	brokers := []string{"localhost:9092"} //Kafka's broker's address
	topic := "example-topic"

	producer, err := sarama.NewSyncProducer(brokers, nil) //create a new sync kafka producer
	if err != nil {
		log.Fatalf("Failed to start Kafka producer: %v", err)
	}
	defer func() { //ensure producer is closed when program exits
		if err := producer.Close(); err != nil {
			log.Fatalf("Failed to close producer: %v", err)
		}
	}()

	message := &sarama.ProducerMessage{ //construct a message to send to worker
		Topic: topic,
		Value: sarama.StringEncoder("Hello, Kafka!"),
	}

	partition, offset, err := producer.SendMessage(message) //send message to kafka using producer
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
}
