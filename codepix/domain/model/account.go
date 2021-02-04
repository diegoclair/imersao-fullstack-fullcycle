package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/twinj/uuid"
)

//AccountRepositoryInterface is the interface for AccountModel
type AccountRepositoryInterface interface {
	AddAccount(account *Account) error
	FindAccountByID(id string) (*Account, error)
}

//Account entity model
type Account struct {
	Base      `valid:"required"`
	OwnerName string `gorm:"column:owner_name;type:varchar(255);not null" valid:"notnull"`
	Bank      *Bank  `valid:"-"`
	BankID    string `gorm:"column:bank_id;type:uuid;not null" valid:"-"`
	Number    string `json:"number" gorm:"type:varchar(20)" valid:"notnull"`
	PixKeys   []*Pix `gorm:"ForeignKey:AccountID" valid:"-"`
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
