package entity

import (
	"time"

	"github.com/diegoclair/imersao/codepix/infrastructure/validate"
	"github.com/twinj/uuid"
)

//Account entity model
type Account struct {
	Base      `validate:"required"`
	OwnerName string `gorm:"column:owner_name;type:varchar(255);not null" validate:"required"`
	Bank      *Bank
	BankID    string `gorm:"column:bank_id;type:uuid;not null"`
	Number    string `json:"number" gorm:"type:varchar(20)" validate:"required"`
	PixKeys   []*Pix `gorm:"ForeignKey:AccountID"`
}

func (account *Account) isValid() error {
	return validate.Struct(account, "Account")
}

//NewAccount return a new Account model
func NewAccount(bank *Bank, ownerName, number string) (*Account, error) {
	account := Account{
		OwnerName: ownerName,
		Bank:      bank,
		Number:    number,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, err
	}
	return &account, nil
}
