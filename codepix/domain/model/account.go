package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/twinj/uuid"
)

//AccountRepository is the interface for AccountModel
type AccountRepository interface {
	AddAccount(account *Account) error
	FindAccountByID(id, kind string) (*Account, error)
}

//Account entity model
type Account struct {
	Base      `valid:"required"`
	OwnerName string    `json:"owner_name" valid:"notnull"`
	Bank      *Bank     `valid:"-"`
	Number    string    `json:"number" valid:"notnull"`
	PixKeys   []*PixKey `valid:"-"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}
	return nil
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
