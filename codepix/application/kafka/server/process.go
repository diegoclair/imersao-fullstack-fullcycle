package server

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/diegoclair/imersao/codepix/application/factory"
	"github.com/diegoclair/imersao/codepix/application/kafka/model"
	"github.com/diegoclair/imersao/codepix/domain"
	"github.com/diegoclair/imersao/codepix/domain/contract"
	"github.com/diegoclair/imersao/codepix/domain/entity"
	"github.com/diegoclair/imersao/codepix/infrastructure/config"
)

type KafkaProcessor struct {
	Producer     *kafka.Producer
	DeliveryChan chan kafka.Event

	cfg                *config.EnvironmentVariables
	pixService         contract.PixService
	transactionService contract.TransactionService
}

func NewKafkaProcessor(producer *kafka.Producer, deliveryChan chan kafka.Event, factory *factory.Services) *KafkaProcessor {
	return &KafkaProcessor{
		Producer:     producer,
		DeliveryChan: deliveryChan,

		cfg:                factory.Cfg,
		pixService:         factory.PixService,
		transactionService: factory.TransactionService,
	}
}

func (k *KafkaProcessor) Consume() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": k.cfg.Kafka.BootstrapServers,
		"group.id":          k.cfg.Kafka.ConsumerGroupID,
		"auto.offset.reset": "earliest",
	}

	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		panic(err)
	}

	topics := []string{k.cfg.Kafka.TransactionTopic, k.cfg.Kafka.TransactionConfirmationTopic}
	c.SubscribeTopics(topics, nil)
	log.Println("kafka consumer has been started")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			err = k.processMessage(msg)
			if err != nil {
				log.Fatalf("Error to process message: %v", err)
			}
		}
	}
}

func (k *KafkaProcessor) processMessage(msg *kafka.Message) error {

	transaction, err := k.parseMessageToTransaction(msg)
	if err != nil {
		return err
	}

	switch topic := *msg.TopicPartition.Topic; topic {
	case k.cfg.Kafka.TransactionTopic:
		return k.processTransactionMessage(transaction)
	case k.cfg.Kafka.TransactionConfirmationTopic:
		return k.processTransactionConfirmationMessage(transaction)
	default:
		return fmt.Errorf("not a valid topic - message: %v", string(msg.Value))
	}
}

func (k *KafkaProcessor) parseMessageToTransaction(msg *kafka.Message) (*model.Transaction, error) {

	transaction := model.NewTransaction()
	err := transaction.ParseJSON(msg.Value)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (k *KafkaProcessor) processTransactionMessage(transaction *model.Transaction) error {

	createdTransaction, err := k.saveTransaction(transaction)
	if err != nil {
		return err
	}

	topic := "bank" + createdTransaction.PixTo.Account.Bank.Code
	transaction.ID = createdTransaction.ID
	transaction.Status = domain.TransactionPending

	return k.sendToDestinationBank(topic, transaction)
}

func (k *KafkaProcessor) saveTransaction(transaction *model.Transaction) (*entity.Transaction, error) {

	return k.transactionService.Register(
		transaction.AccountID,
		transaction.PixKeyTo,
		transaction.PixKeyTypeTo,
		transaction.Description,
		transaction.ID,
		transaction.Amount,
	)
}

func (k *KafkaProcessor) processTransactionConfirmationMessage(transaction *model.Transaction) error {

	if transaction.Status == domain.TransactionConfirmed {
		confirmedTransaction, err := k.transactionService.Confirm(transaction.ID)
		if err != nil {
			return err
		}

		topic := "bank" + confirmedTransaction.AccountFrom.Bank.Code
		return k.sendToSourceBank(topic, transaction)
	}
	if transaction.Status == domain.TransactionCompleted {
		_, err := k.transactionService.Complete(transaction.ID)
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (k *KafkaProcessor) sendToSourceBank(topic string, transaction *model.Transaction) error {
	return k.sendToBank(topic, transaction)
}

func (k *KafkaProcessor) sendToDestinationBank(topic string, transaction *model.Transaction) error {
	return k.sendToBank(topic, transaction)
}

func (k *KafkaProcessor) sendToBank(topic string, transaction *model.Transaction) error {

	transactionJSON, err := transaction.ToJSON()
	if err != nil {
		return err
	}

	return Publish(string(transactionJSON), topic, k.Producer, k.DeliveryChan)
}