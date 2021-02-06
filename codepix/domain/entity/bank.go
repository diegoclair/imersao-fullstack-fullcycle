package entity

import (
	"fmt"
	"time"

	"github.com/diegoclair/go_utils-lib/v2/validstruct"
	"github.com/twinj/uuid"
)

//Bank entity model
type Bank struct {
	Base     `validate:"required"`
	Code     string     `json:"code" gorm:"type:varchar(20)" validate:"required"`
	Name     string     `json:"name" gorm:"type:varchar(255)" validate:"required"`
	Accounts *[]Account `gorm:"ForeignKey:BankID"`
}

func (bank *Bank) isValid() error {
	err := validstruct.ValidateStruct(bank)
	if err != nil {
		return fmt.Errorf(err.Message())
	}
	return nil
}

//NewBank return a new Bank model
func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid()
	if err != nil {
		return nil, err
	}
	return &bank, nil
}
