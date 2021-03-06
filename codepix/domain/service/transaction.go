package service

import (
	"log"

	"github.com/diegoclair/imersao/codepix/contract"
	"github.com/diegoclair/imersao/codepix/domain/entity"
)

type transactionService struct {
	svc *Service
}

//newTransactionService return a new instance of the service
func newTransactionService(svc *Service) contract.TransactionService {
	return &transactionService{
		svc: svc,
	}
}

func (s *transactionService) Register(accountIDFrom, pixKeyTo, pixKeyKindTo, description, id string, amount float64) (*entity.Transaction, error) {

	account, err := s.svc.dm.Postgres().Account().FindAccountByID(accountIDFrom)
	if err != nil {
		return nil, err
	}

	pixKey, err := s.svc.dm.Postgres().Pix().FindPixByKey(pixKeyTo, pixKeyKindTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(account, amount, pixKey, description, id)
	if err != nil {
		return nil, err
	}

	err = s.svc.dm.Postgres().Transaction().Register(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *transactionService) Confirm(transactionID string) (*entity.Transaction, error) {

	transaction, err := s.svc.dm.Postgres().Transaction().FindByID(transactionID)
	if err != nil {
		log.Println("transactionUseCase.Confirm.FindByID", err)
		return nil, err
	}

	transaction.Confirm()

	err = s.svc.dm.Postgres().Transaction().Save(transaction)
	if err != nil {
		log.Println("transactionUseCase.Confirm.Save", err)
		return nil, err
	}

	return transaction, nil
}

func (s *transactionService) Complete(transactionID string) (*entity.Transaction, error) {

	transaction, err := s.svc.dm.Postgres().Transaction().FindByID(transactionID)
	if err != nil {
		log.Println("transactionUseCase.Complete.FindByID", err)
		return nil, err
	}

	transaction.Complete()

	err = s.svc.dm.Postgres().Transaction().Save(transaction)
	if err != nil {
		log.Println("transactionUseCase.Complete.Save", err)
		return nil, err
	}

	return transaction, nil
}

func (s *transactionService) Error(transactionID, reason string) (*entity.Transaction, error) {

	transaction, err := s.svc.dm.Postgres().Transaction().FindByID(transactionID)
	if err != nil {
		log.Println("transactionUseCase.Error.FindByID", err)
		return nil, err
	}

	transaction.Cancel(reason)

	err = s.svc.dm.Postgres().Transaction().Save(transaction)
	if err != nil {
		log.Println("transactionUseCase.Error.Save", err)
		return nil, err
	}

	return transaction, nil
}
