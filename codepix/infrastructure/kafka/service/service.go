package service

import (
	"github.com/diegoclair/imersao/codepix/contract"
	"github.com/diegoclair/imersao/codepix/infrastructure/config"
)

type kafkaService struct {
	kfk contract.KafkaService
	cfg *config.EnvironmentVariables
}

// NewKafkaManager return a kafka service instance
func NewKafkaManager(cfg *config.EnvironmentVariables) contract.KafkaManager {
	kfk := new(kafkaService)
	return &kafkaService{
		cfg: cfg,
		kfk: kfk.Kafka(),
	}
}

func (k *kafkaService) Kafka() contract.KafkaService {
	return newKafkaService(k.cfg)
}
