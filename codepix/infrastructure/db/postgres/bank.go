package postgres

import (
	"fmt"

	"github.com/diegoclair/imersao/codepix/domain/model"
	"github.com/jinzhu/gorm"
)

// bankRepo repository
type bankRepo struct {
	db *gorm.DB
}

//NewBankRepo return the implementation of bankRepo interface
func NewBankRepo(db *gorm.DB) *bankRepo {
	return &bankRepo{
		db: db,
	}
}

func (r bankRepo) AddBank(bank *model.Bank) error {
	return r.db.Create(bank).Error
}

func (r bankRepo) FindBankByID(id string) (*model.Bank, error) {
	var bank model.Bank

	r.db.First(&bank, "id = ?", id)
	if bank.ID == "" {
		return nil, fmt.Errorf("no bank was found")
	}
	return &bank, nil
}
