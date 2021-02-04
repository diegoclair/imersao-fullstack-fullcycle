package repository

import (
	"fmt"

	"github.com/diegoclair/imersao-fullstack-fullcycle/domain/model"
	"github.com/jinzhu/gorm"
)

// PixKeyReposirotyDB repository
type PixKeyReposirotyDB struct {
	DB *gorm.DB
}

func (r PixKeyReposirotyDB) AddPixKey(pixKey *model.PixKey) (*model.PixKey, error) {
	err := r.DB.Create(pixKey).Error
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}

func (r PixKeyReposirotyDB) FindKeyByID(key, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	r.DB.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)
	if pixKey.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}
	return &pixKey, nil
}
