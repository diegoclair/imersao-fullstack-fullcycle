package contract

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaManager holds the methods that interact with kafka plataform
type KafkaManager interface {
	Kafka() KafkaService
}

// KafkaService defines the interface for kafka service
type KafkaService interface {
	NewConsumer() (*kafka.Consumer, error)
	NewProducer() (*kafka.Producer, error)
	Publish(msg, topic string, producer *kafka.Producer, deliveryChan chan kafka.Event) error
	DeliveryReport(deliveryChan chan kafka.Event)
}
