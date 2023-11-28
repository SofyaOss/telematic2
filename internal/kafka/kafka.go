package kafka

import (
	"encoding/json"
	"log"

	"practice/storage"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Producer contains information about kafka producer
type Producer struct {
	prod  *kafka.Producer
	topic string
}

// Consumer contains information about kafka consumer
type Consumer struct {
	cons *kafka.Consumer
}

// NewProducer creates new kafka producer
func NewProducer(addr string) (*Producer, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": addr, // "kafka:9092"
	}
	topic := "telematicTopic"
	producer, err := kafka.NewProducer(config)
	if err != nil {
		log.Printf("could not connect to kafka: %s", err)
		return nil, err
	}
	p := Producer{prod: producer, topic: topic}
	return &p, nil
}

// Produce produces telematic data to kafka
func Produce(producer *Producer, item *storage.Car) error {
	mes, err := json.Marshal(item)
	if err != nil {
		log.Printf("Could not convert data to json: %s", err)
		return err
	}
	err = producer.prod.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &producer.topic, Partition: kafka.PartitionAny},
		Value:          mes,
	}, nil)
	if err != nil {
		log.Println("Could not produce data", err)
		return err
	}
	return nil
}
