package entity

import (
	"time"

	"github.com/diegoclair/imersao/codepix/infrastructure/validate"
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
	return validate.Struct(bank, "Bank")
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
