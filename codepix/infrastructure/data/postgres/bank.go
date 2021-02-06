package postgres

import (
	"fmt"

	"github.com/diegoclair/imersao/codepix/domain/entity"
	"github.com/jinzhu/gorm"
)

// bankRepo repository
type bankRepo struct {
	db *gorm.DB
}

//newBankRepo return the implementation of bankRepo interface
func newBankRepo(db *gorm.DB) *bankRepo {
	return &bankRepo{
		db: db,
	}
}

func (r bankRepo) AddBank(bank *entity.Bank) error {
	return r.db.Create(bank).Error
}

func (r bankRepo) FindBankByID(id string) (*entity.Bank, error) {
	var bank entity.Bank

	r.db.First(&bank, "id = ?", id)
	if bank.ID == "" {
		return nil, fmt.Errorf("no bank was found")
	}
	return &bank, nil
}
