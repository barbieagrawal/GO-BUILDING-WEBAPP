package main

import (
	"log"

	"github.com/IBM/sarama"
)

func main() {
	brokers := []string{"localhost:9092"}
	topic := "example-topic"

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka consumer: %v", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalf("Failed to close consumer: %v", err)
		}
	}()

	partitions, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatalf("Failed to get partitions: %v", err)
	}

	for _, partition := range partitions {
		pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
		if err != nil {
			log.Fatalf("Failed to start consuming partition %d: %v", partition, err)
		}

		defer pc.Close()

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				log.Printf("Consumed message: key = %s, value = %s, offset = %d\n",
					string(message.Key), string(message.Value), message.Offset)
			}
		}(pc)
	}

	select {} // Block forever
}
