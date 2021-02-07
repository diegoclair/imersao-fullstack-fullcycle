package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/diegoclair/imersao/codepix/application/factory"
	"github.com/diegoclair/imersao/codepix/application/kafka/server"
	"github.com/diegoclair/imersao/codepix/infrastructure/kafka/service"
)

// StartKafkaServer to start/setup a kafka server
func StartKafkaServer() {

	factory := factory.GetDomainServices()

	kfk := service.NewKafkaManager(factory.Cfg)

	deliveryChan := make(chan kafka.Event)

	go kfk.Kafka().DeliveryReport(deliveryChan)

	kafkaProcessor := server.NewKafkaProcessor(kfk, deliveryChan, factory)
	kafkaProcessor.Consume()
}
