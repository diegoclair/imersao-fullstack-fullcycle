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

func (r PixKeyReposirotyDB) AddPixKey(pixKey *model.Pix) (*model.Pix, error) {
	err := r.DB.Create(pixKey).Error
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}

func (r PixKeyReposirotyDB) FindKeyByID(key, kind string) (*model.Pix, error) {
	var pix model.Pix

	r.DB.Preload("Account.Bank").First(&pix, "kind = ? and key = ?", kind, key)
	if pix.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}
	return &pix, nil
}
