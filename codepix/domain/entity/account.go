package entity

import (
	"fmt"
	"strconv"
	"time"

	"github.com/diegoclair/go_utils-lib/v2/validstruct"
	"github.com/twinj/uuid"
)

//Account entity model
type Account struct {
	Base      `validate:"required"`
	OwnerName string `gorm:"column:owner_name;type:varchar(255);not null" validate:"required"`
	Bank      *Bank  `validate:"-"`
	BankID    string `gorm:"column:bank_id;type:uuid;not null"`
	Number    string `json:"number" gorm:"type:varchar(20)" validate:"required"`
	PixKeys   []*Pix `gorm:"ForeignKey:AccountID"`
}

func (account *Account) isValid() error {
	err := validstruct.ValidateStruct(account)
	if err != nil {
		validationErrors := err.Causes().([]string)
		fmt.Println("Error to validate account entity struct")
		for i := range validationErrors {
			fmt.Println(strconv.Itoa(i+1) + " - " + validationErrors[i])
		}

		return fmt.Errorf(fmt.Sprintf("%v", validationErrors))
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
