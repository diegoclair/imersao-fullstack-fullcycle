package usecase

import (
	"log"

	"github.com/diegoclair/imersao-fullstack-fullcycle/domain/model"
)

type transactionUseCase struct {
	transactionRepository model.TransactionRepositoryInterface
	accountRepository     model.AccountRepositoryInterface
	pixKeyRepository      model.PixKeyRepositoryInterface
}

func (a *transactionUseCase) Register(accountIDFrom, pixKeyTo, pixKeyKindTo, description string, amount float64) (*model.Transaction, error) {

	account, err := a.accountRepository.FindAccountByID(accountIDFrom)
	if err != nil {
		return nil, err
	}

	pixKey, err := a.pixKeyRepository.FindKeyByID(pixKeyTo, pixKeyKindTo)
	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, amount, pixKey, description)
	if err != nil {
		return nil, err
	}

	err = a.transactionRepository.Register(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (a *transactionUseCase) Confirm(transactionID string) (*model.Transaction, error) {

	transaction, err := a.transactionRepository.FindByID(transactionID)
	if err != nil {
		log.Println("transactionUseCase.Confirm.FindByID", err)
		return nil, err
	}

	err = transaction.Confirm()
	if err != nil {
		log.Println("transactionUseCase.Confirm.Confirm", err)
		return nil, err
	}

	err = a.transactionRepository.Save(transaction)
	if err != nil {
		log.Println("transactionUseCase.Confirm.Save", err)
		return nil, err
	}

	return transaction, nil
}

func (a *transactionUseCase) Complete(transactionID string) (*model.Transaction, error) {

	transaction, err := a.transactionRepository.FindByID(transactionID)
	if err != nil {
		log.Println("transactionUseCase.Complete.FindByID", err)
		return nil, err
	}

	err = transaction.Complete()
	if err != nil {
		log.Println("transactionUseCase.Complete.Complete", err)
		return nil, err
	}

	err = a.transactionRepository.Save(transaction)
	if err != nil {
		log.Println("transactionUseCase.Complete.Save", err)
		return nil, err
	}

	return transaction, nil
}

func (a *transactionUseCase) Error(transactionID, reason string) (*model.Transaction, error) {

	transaction, err := a.transactionRepository.FindByID(transactionID)
	if err != nil {
		log.Println("transactionUseCase.Error.FindByID", err)
		return nil, err
	}

	err = transaction.Cancel(reason)
	if err != nil {
		log.Println("transactionUseCase.Error.Cancel", err)
		return nil, err
	}

	err = a.transactionRepository.Save(transaction)
	if err != nil {
		log.Println("transactionUseCase.Error.Save", err)
		return nil, err
	}

	return transaction, nil
}
