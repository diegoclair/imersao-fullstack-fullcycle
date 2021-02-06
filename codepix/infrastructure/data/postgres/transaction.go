package postgres

import (
	"fmt"

	"github.com/diegoclair/imersao/codepix/domain/entity"
	"github.com/jinzhu/gorm"
)

// transactionRepo repository
type transactionRepo struct {
	db *gorm.DB
}

//newTransactionRepo return the implementation of transactionRepo interface
func newTransactionRepo(db *gorm.DB) *transactionRepo {
	return &transactionRepo{
		db: db,
	}
}

func (r transactionRepo) Register(transaction *entity.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r transactionRepo) Save(transaction *entity.Transaction) error {
	return r.db.Save(transaction).Error
}

func (r transactionRepo) FindByID(id string) (*entity.Transaction, error) {
	var transaction entity.Transaction

	r.db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("No transaction was not found")
	}
	return &transaction, nil
}
