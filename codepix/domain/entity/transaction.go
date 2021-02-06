package entity

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/diegoclair/go_utils-lib/v2/validstruct"
	"github.com/twinj/uuid"
)

//Transaction Standard values
const (
	TransactionPending   string = "pending"
	TransactionCompleted string = "completed"
	TransactionError     string = "error"
	TransactionConfirmed string = "confirmed"
)

//Transaction entity model
type Transaction struct {
	Base              `validate:"required"`
	AccountFrom       *Account `validate:"required,dive,required"`
	AccountFromID     string   `gorm:"column:account_from_id;type:uuid;" validate:"required,gt=0"`
	Amount            float64  `json:"amount" gorm:"type:float" validate:"required"`
	PixTo             *Pix     `validate:"required,dive,required"`
	PixKeyToID        string   `gorm:"column:pix_key_to_id;type:uuid;" validate:"required,uuid4"`
	Status            string   `json:"status" gorm:"type:varchar(20)" validate:"required"`
	Description       string   `json:"description" gorm:"type:varchar(255)" validate:"required"`
	CancelDescription string   `json:"cancel_description" gorm:"type:varchar(255)" `
}

func (t *Transaction) isValid() error {

	if t.Status != TransactionCompleted &&
		t.Status != TransactionConfirmed &&
		t.Status != TransactionError &&
		t.Status != TransactionPending {
		return errors.New("Invalid Status for the transaction")
	}

	if t.PixTo.AccountID == t.AccountFromID {
		return errors.New("The source and destination account cannot be the same")
	}

	err := validstruct.ValidateStruct(t)
	if err != nil {
		validationErrors := err.Causes().([]string)
		fmt.Println("Error to validate transaction entity struct")
		for i := range validationErrors {
			fmt.Println(strconv.Itoa(i+1) + " - " + validationErrors[i])
		}

		return fmt.Errorf(fmt.Sprintf("%v", validationErrors))
	}
	return nil
}

//NewTransaction return a new transaction model
func NewTransaction(accountFrom *Account, amount float64, pixTo *Pix, description, id string) (*Transaction, error) {

	transaction := Transaction{
		AccountFrom:   accountFrom,
		AccountFromID: accountFrom.ID,
		Amount:        amount,
		PixTo:         pixTo,
		PixKeyToID:    pixTo.ID,
		Status:        TransactionPending,
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
func (t *Transaction) Confirm() error {
	return t.setStatus(TransactionCompleted)
}

//Complete to set completed to transaction status
func (t *Transaction) Complete() error {
	return t.setStatus(TransactionCompleted)
}

//Cancel to set canceled to transaction status
func (t *Transaction) Cancel(reason string) error {
	t.CancelDescription = reason
	return t.setStatus(TransactionError)
}

func (t *Transaction) setStatus(status string) error {
	t.Status = status
	t.UpdatedAt = time.Now()
	return t.isValid()
}
