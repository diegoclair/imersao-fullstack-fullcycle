package kafka

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/diegoclair/imersao/codepix/domain/contract"
)

type KafkaProcessor struct {
	Database     contract.DataManager
	Producer     *ckafka.Producer
	DeliveryChan chan ckafka.Event
}

func NewKafkaProcessor(database contract.DataManager, producer *ckafka.Producer, deliveryChan chan ckafka.Event) *KafkaProcessor {
	return &KafkaProcessor{
		Database:     database,
		Producer:     producer,
		DeliveryChan: deliveryChan,
	}
}

func (s *KafkaProcessor) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "consumergroup",
		"auto.offset.reset": "earliest",
	}

	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		panic(err)
	}

	topics := []string{"teste"}
	c.SubscribeTopics(topics, nil)
	log.Println("kafka consumer has been started")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			s.processMessage(msg)
		}
	}
}

func (s *KafkaProcessor) processMessage(msg *ckafka.Message) {
	transactionsTopic := "transactions"
	transactionConfirmationTopic := "transaction_confirmation"

	switch topic := *msg.TopicPartition.Topic; topic {
	case transactionsTopic:
	case transactionConfirmationTopic:
	default:
		fmt.Println("not a valid topic", string(msg.Value))
	}
}

func (s *KafkaProcessor) processTransaction(msg *ckafka.Message) error {
	//transaction := model.NewTransaction()
	return nil
}
