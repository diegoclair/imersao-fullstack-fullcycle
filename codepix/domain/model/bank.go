package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/twinj/uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

//BankRepository is the interface for BankModel
type BankRepository interface {
	AddBank(bank *Bank) error
}

//Bank entity model
type Bank struct {
	Base     `valid:"required"`
	Code     string     `json:"code,omitempty" valid:"notnull"`
	Name     string     `json:"name,omitempty" valid:"notnull"`
	Accounts *[]Account `valid:"-"`
}

func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)
	if err != nil {
		return err
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
