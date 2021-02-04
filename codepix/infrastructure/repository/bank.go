package repository

import (
	"fmt"

	"github.com/diegoclair/imersao-fullstack-fullcycle/domain/model"
	"github.com/jinzhu/gorm"
)

// 	AddBank(bank *Bank) error

// bankReposirotyDB repository
type bankReposirotyDB struct {
	DB *gorm.DB
}

func (r bankReposirotyDB) AddBank(bank *model.Bank) error {
	return r.DB.Create(bank).Error
}

func (r bankReposirotyDB) FindBankByID(id string) (*model.Bank, error) {
	var bank model.Bank

	r.DB.First(&bank, "id = ?", id)
	if bank.ID == "" {
		return nil, fmt.Errorf("no bank was found")
	}
	return &bank, nil
}
