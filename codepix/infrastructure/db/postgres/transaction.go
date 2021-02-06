package postgres

import (
	"fmt"

	"github.com/diegoclair/imersao-fullstack-fullcycle/codepix/domain/model"
	"github.com/jinzhu/gorm"
)

// transactionRepo repository
type transactionRepo struct {
	db *gorm.DB
}

//NewTransactionRepo return the implementation of transactionRepo interface
func NewTransactionRepo(db *gorm.DB) *transactionRepo {
	return &transactionRepo{
		db: db,
	}
}

func (r transactionRepo) Register(transaction *model.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r transactionRepo) Save(transaction *model.Transaction) error {
	return r.db.Save(transaction).Error
}

func (r transactionRepo) FindByID(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	r.db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("No transaction was not found")
	}
	return &transaction, nil
}
