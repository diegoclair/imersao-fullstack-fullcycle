package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/twinj/uuid"
)

//Transaction Standard values
const (
	TransactionPending   string = "pending"
	TransactionCompleted string = "completed"
	TransactionError     string = "errpr"
	TransactionConfirmed string = "confirmed"
)

//TransactionRepository is the interface for TransactionModel
type TransactionRepository interface {
	RegisterTransaction(transaction *Transaction) error
	Save(transaction *Transaction) error
	FindByID(id string) (*Transaction, error)
}

//Transaction entity model
type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	Amount            float64  `json:"amount" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	Status            string   `json:"status" valid:"notnull"`
	Description       string   `json:"description" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" valid:"-"`
}

func (t *Transaction) isValid() error {

	if t.Amount <= 0 {
		return errors.New("The amout must be greater than 0")
	}

	if t.Status != TransactionCompleted &&
		t.Status != TransactionConfirmed &&
		t.Status != TransactionError &&
		t.Status != TransactionPending {
		return errors.New("Invalid Status for the transaction")
	}

	if t.PixKeyTo.AccountID == t.AccountFrom.ID {
		return errors.New("The source and destination account cannot be the same")
	}

	_, err := govalidator.ValidateStruct(t)
	if err != nil {
		return err
	}
	return nil
}

//NewTransaction return a new transaction model
func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom: accountFrom,
		Amount:      amount,
		PixKeyTo:    pixKeyTo,
		Status:      TransactionPending,
		Description: description,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

//Confirm to set confirmed to transaction status
func (t *Transaction) Confirm() error {
	return t.setStatus(TransactionCompleted)
}

//Complete to set completed to transaction status
func (t *Transaction) Complete() error {
	return t.setStatus(TransactionCompleted)
}

//Cancel to set canceled to transaction status
func (t *Transaction) Cancel(description string) error {
	t.Description = description
	return t.setStatus(TransactionCompleted)
}

func (t *Transaction) setStatus(status string) error {
	t.Status = status
	t.UpdatedAt = time.Now()
	return t.isValid()
}
