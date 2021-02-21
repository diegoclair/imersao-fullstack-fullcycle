package model

import (
	"encoding/json"

	"github.com/diegoclair/imersao/codepix/util/validate"
)

//Transaction view model
type Transaction struct {
	ID           string  `json:"id" validate:"required,uuid4"`
	AccountID    string  `json:"accountID" validate:"required,uuid4"`
	Amount       float64 `json:"amount" validate:"required,numeric"`
	PixKeyTo     string  `json:"pixKeyTo" validate:"required"`
	PixKeyTypeTo string  `json:"pixKeyTypeTo" validate:"required"`
	Description  string  `json:"description"`
	Status       string  `json:"status"`
	Error        string  `json:"error"`
}

func (t *Transaction) isValid() error {
	return validate.Struct(t, "TransactionModel")
}

//ParseJSON to parse a transaction message bytes received from kafka to json
func (t *Transaction) ParseJSON(data []byte) error {

	err := json.Unmarshal(data, t)
	if err != nil {
		return err
	}
	return t.isValid()
}

//NewTransaction return a new transaction instance
func NewTransaction() *Transaction {
	return &Transaction{}
}

//ToJSON to parse a transaction message bytes received from kafka to json
func (t *Transaction) ToJSON() ([]byte, error) {

	err := t.isValid()
	if err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
