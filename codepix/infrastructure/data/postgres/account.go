package postgres

import (
	"fmt"

	"github.com/diegoclair/imersao/codepix/domain/entity"
	"github.com/jinzhu/gorm"
)

// accountRepo repository
type accountRepo struct {
	db *gorm.DB
}

//newAccountRepo return the implementation of accountRepo interface
func newAccountRepo(db *gorm.DB) *accountRepo {
	return &accountRepo{
		db: db,
	}
}

func (r accountRepo) AddAccount(account *entity.Account) error {
	return r.db.Create(account).Error
}

func (r accountRepo) FindAccountByID(id string) (*entity.Account, error) {
	var account entity.Account

	r.db.Preload("Bank").First(&account, "id = ?", id)
	if account.ID == "" {
		return nil, fmt.Errorf("no account was found")
	}
	return &account, nil
}
