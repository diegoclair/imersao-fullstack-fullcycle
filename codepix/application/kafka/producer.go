package kafka

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// NewKafkaProducer return an instace of a producer
func NewKafkaProducer() *ckafka.Producer {

	//the port 9092 is defined in docker compose
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}

	p, err := ckafka.NewProducer(configMap)
	if err != nil {
		panic(err)
	}

	return p
}

//Publish to publish a message into a kafka topic
func Publish(msg, topic string, producer *ckafka.Producer, deliveryChan chan ckafka.Event) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
	}

	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

//DeliveryReport listen a chan to report if the message was delivered or if got a error to delivery
func DeliveryReport(deliveryChan chan ckafka.Event) {
	for e := range deliveryChan {
		switch event := e.(type) {
		case *ckafka.Message:
			if event.TopicPartition.Error != nil {
				fmt.Println("Delivery Faield", event.TopicPartition)
			} else {
				fmt.Println("Delivered message to:", event.TopicPartition)
			}
		}
	}
}
