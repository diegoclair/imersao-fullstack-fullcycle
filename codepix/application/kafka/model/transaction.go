package model

import (
	"fmt"

	"github.com/diegoclair/go_utils-lib/v2/validstruct"
)

type Transaction struct {
	ID           string  `json:"id" validate:"required,uuid4"`
	AccountID    string  `json:"accountId" validate:"required,uuid4"`
	Amount       float64 `json:"amount" validate:"required,numeric"`
	PixKeyTo     string  `json:"pixKeyTo" validate:"required"`
	PixKeyTypeTo string  `json:"pixKeyTypeTo" validate:"required"`
	Description  string  `json:"description"`
	Status       string  `json:"status" validate:"required"`
	Error        string  `json:"error"`
}

func (t *Transaction) isValid() error {
	err := validstruct.ValidateStruct(t)
	if err != nil {
		fmt.Errorf("Error to validate transaction struct: %v", err)
		return err
	}
	return nil
}
