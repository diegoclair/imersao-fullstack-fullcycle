package service

import (
	"github.com/diegoclair/imersao/codepix/domain/contract"
	"github.com/diegoclair/imersao/codepix/infrastructure/config"
)

type kafkaService struct {
	cfg *config.EnvironmentVariables
}

// NewKafkaManager return a kafka service instance
func NewKafkaManager(cfg *config.EnvironmentVariables) contract.KafkaManager {
	return &kafkaService{
		cfg: cfg,
	}
}

func (k *kafkaService) Kafka() contract.KafkaService {
	return newKafkaService(k.cfg)
}
