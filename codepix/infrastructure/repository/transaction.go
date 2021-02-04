package repository

import (
	"fmt"

	"github.com/diegoclair/imersao-fullstack-fullcycle/domain/model"
	"github.com/jinzhu/gorm"
)

// transactionReposirotyDB repository
type transactionReposirotyDB struct {
	DB *gorm.DB
}

func (r transactionReposirotyDB) Register(transaction *model.Transaction) error {
	return r.DB.Create(transaction).Error
}

func (r transactionReposirotyDB) Save(transaction *model.Transaction) error {
	return r.DB.Save(transaction).Error
}

func (r transactionReposirotyDB) FindByID(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	r.DB.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("No transaction was not found")
	}
	return &transaction, nil
}
