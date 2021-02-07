package server

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/diegoclair/imersao/codepix/infrastructure/config"
)

// NewKafkaProducer return an instace of a producer
func NewKafkaProducer(cfg *config.EnvironmentVariables) *kafka.Producer {

	//the port 9092 is defined in docker compose
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.BootstrapServers,
	}

	p, err := kafka.NewProducer(configMap)
	if err != nil {
		panic(err)
	}

	return p
}

//Publish to publish a message into a kafka topic
func Publish(msg, topic string, producer *kafka.Producer, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(msg),
	}

	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

//DeliveryReport listen a chan to report if the message was delivered or if got a error to delivery
func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch event := e.(type) {
		case *kafka.Message:
			if event.TopicPartition.Error != nil {
				fmt.Println("Delivery Faield", event.TopicPartition)
			} else {
				fmt.Println("Delivered message to:", event.TopicPartition)
			}
		}
	}
}
