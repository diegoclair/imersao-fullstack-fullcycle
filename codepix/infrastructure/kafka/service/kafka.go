package service

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/diegoclair/imersao/codepix/infrastructure/config"
)

//KafkaService holds kafka service functions
type KafkaService struct {
	cfg *config.EnvironmentVariables
}

//newKafkaService returns a kafka service instance
func newKafkaService(cfg *config.EnvironmentVariables) *KafkaService {
	return &KafkaService{
		cfg: cfg,
	}
}

//NewConsumer returns a consumer that can be used to consume topics on kafka
func (s *KafkaService) NewConsumer() (*kafka.Consumer, error) {

	configMap := &kafka.ConfigMap{
		"bootstrap.servers": s.cfg.Kafka.BootstrapServers,
		"group.id":          s.cfg.Kafka.ConsumerGroupID,
		"auto.offset.reset": "earliest",
	}

	return kafka.NewConsumer(configMap)
}

// NewProducer return an instace of a producer
func (s *KafkaService) NewProducer() (*kafka.Producer, error) {

	configMap := &kafka.ConfigMap{
		"bootstrap.servers": s.cfg.Kafka.BootstrapServers,
	}
	return kafka.NewProducer(configMap)
}

//Publish to publish a message into a kafka topic
func (s *KafkaService) Publish(msg, topic string, producer *kafka.Producer, deliveryChan chan kafka.Event) error {

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
func (s *KafkaService) DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch event := e.(type) {
		case *kafka.Message:
			if event.TopicPartition.Error != nil {
				log.Println("Delivery Faield", event.TopicPartition)
			} else {
				log.Println("Delivered message to:", event.TopicPartition)
			}
		}
	}
}
