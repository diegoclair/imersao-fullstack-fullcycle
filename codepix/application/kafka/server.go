package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/diegoclair/imersao/codepix/application/factory"
	"github.com/diegoclair/imersao/codepix/application/kafka/server"
)

// StartKafkaServer to start/setup a kafka server
func StartKafkaServer() {

	factory := factory.GetDomainServices()

	producer := server.NewKafkaProducer(factory.Cfg)

	deliveryChan := make(chan kafka.Event)
	//server.Publish("Olá Kafka", "teste", producer, deliveryChan)
	go server.DeliveryReport(deliveryChan)

	kafkaProcessor := server.NewKafkaProcessor(producer, deliveryChan, factory)
	kafkaProcessor.Consume()
}
