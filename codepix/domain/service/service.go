package service

import "github.com/diegoclair/imersao/codepix/domain/contract"

// Service holds the domain service repositories
type Service struct {
	db contract.DataManager
}

// New returns a new domain Service instance
func New(db contract.DataManager) *Service {
	svc := new(Service)
	svc.db = db

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
