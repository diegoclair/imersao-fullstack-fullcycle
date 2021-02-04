package repository

import (
	"fmt"

	"github.com/diegoclair/imersao-fullstack-fullcycle/domain/model"
	"github.com/jinzhu/gorm"
)

// accountReposirotyDB repository
type accountReposirotyDB struct {
	DB *gorm.DB
}

func (r accountReposirotyDB) AddAccount(account *model.Account) error {
	return r.DB.Create(account).Error
}

func (r accountReposirotyDB) FindAccountByID(id string) (*model.Account, error) {
	var account model.Account

	r.DB.Preload("Bank").First(&account, "id = ?", id)
	if account.ID == "" {
		return nil, fmt.Errorf("no account was found")
	}
	return &account, nil
}
