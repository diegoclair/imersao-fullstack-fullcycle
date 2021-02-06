package postgres

import (
	"fmt"

	"github.com/diegoclair/imersao-fullstack-fullcycle/codepix/domain/model"
	"github.com/jinzhu/gorm"
)

// accountRepo repository
type accountRepo struct {
	db *gorm.DB
}

//NewAccountRepo return the implementation of accountRepo interface
func NewAccountRepo(db *gorm.DB) *accountRepo {
	return &accountRepo{
		db: db,
	}
}

func (r accountRepo) AddAccount(account *model.Account) error {
	return r.db.Create(account).Error
}

func (r accountRepo) FindAccountByID(id string) (*model.Account, error) {
	var account model.Account

	r.db.Preload("Bank").First(&account, "id = ?", id)
	if account.ID == "" {
		return nil, fmt.Errorf("no account was found")
	}
	return &account, nil
}
