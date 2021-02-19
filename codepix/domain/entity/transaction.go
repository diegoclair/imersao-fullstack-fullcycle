package entity

import (
	"errors"
	"time"

	"github.com/diegoclair/imersao/codepix/domain"
	"github.com/diegoclair/imersao/codepix/infrastructure/validate"
	"github.com/twinj/uuid"
)

//Transaction entity model
type Transaction struct {
	Base              `validate:"required"`
	AccountFrom       *Account `validate:"required,dive,required"`
	AccountFromID     string   `gorm:"column:account_from_id;type:uuid;" validate:"required,uuid4,gt=0"`
	Amount            float64  `json:"amount" gorm:"type:float" validate:"required"`
	PixTo             *Pix
	PixKeyToID        string `gorm:"column:pix_key_to_id;type:uuid;" validate:"required,uuid4"`
	Status            string `json:"status" gorm:"type:varchar(20)" validate:"required"`
	Description       string `json:"description" gorm:"type:varchar(255)" validate:"required"`
	CancelDescription string `json:"cancel_description" gorm:"type:varchar(255)" `
}

func (t *Transaction) isValid() error {

	if t.Status != domain.TransactionCompleted &&
		t.Status != domain.TransactionConfirmed &&
		t.Status != domain.TransactionError &&
		t.Status != domain.TransactionPending {
		return errors.New("Invalid Status for the transaction")
	}

	if t.PixTo.AccountID == t.AccountFromID {
		return errors.New("The source and destination account cannot be the same")
	}

	return validate.Struct(t, "Transaction")
}

//NewTransaction return a new transaction model
func NewTransaction(accountFrom *Account, amount float64, pixTo *Pix, description, id string) (*Transaction, error) {

	transaction := Transaction{
		AccountFrom:   accountFrom,
		AccountFromID: accountFrom.ID,
		Amount:        amount,
		PixTo:         pixTo,
		PixKeyToID:    pixTo.ID,
		Status:        domain.TransactionPending,
		Description:   description,
	}

	transaction.ID = id
	if transaction.ID == "" {
		transaction.ID = uuid.NewV4().String()
	}
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

//Confirm to set confirmed to transaction status
func (t *Transaction) Confirm() {
	t.setStatus(domain.TransactionCompleted)
}

//Complete to set completed to transaction status
func (t *Transaction) Complete() {
	t.setStatus(domain.TransactionCompleted)
}

//Cancel to set canceled to transaction status
func (t *Transaction) Cancel(reason string) {
	t.CancelDescription = reason
	t.setStatus(domain.TransactionError)
}

func (t *Transaction) setStatus(status string) {
	t.Status = status
	t.UpdatedAt = time.Now()
}
