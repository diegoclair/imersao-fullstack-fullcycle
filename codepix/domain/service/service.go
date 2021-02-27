package service

import (
	"github.com/diegoclair/imersao/codepix/contract"
	"github.com/diegoclair/imersao/codepix/util/config"
)

// Service holds the domain service repositories
type Service struct {
	dm  contract.DataManager
	cfg *config.EnvironmentVariables
}

// New returns a new domain Service instance
func New(dm contract.DataManager, cfg *config.EnvironmentVariables) *Service {
	svc := new(Service)
	svc.dm = dm
	svc.cfg = cfg

	return svc
}

//Manager defines the services aggregator interface
type Manager interface {
	PixService(svc *Service) contract.PixService
	TransactionService(svc *Service) contract.TransactionService
}

type serviceManager struct {
	svc *Service
}

// NewServiceManager return a service manager instance
func NewServiceManager() Manager {
	return &serviceManager{}
}

func (s *serviceManager) PixService(svc *Service) contract.PixService {
	return newPixService(svc)
}

func (s *serviceManager) TransactionService(svc *Service) contract.TransactionService {
	return newTransactionService(svc)
}
