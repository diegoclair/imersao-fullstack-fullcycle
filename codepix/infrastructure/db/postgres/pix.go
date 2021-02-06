package postgres

import (
	"fmt"

	"github.com/diegoclair/imersao/codepix/domain/model"
	"github.com/jinzhu/gorm"
)

// pixRepo repository
type pixRepo struct {
	db *gorm.DB
}

//NewPixRepo return the implementation of pixRepo interface
func NewPixRepo(db *gorm.DB) *pixRepo {
	return &pixRepo{
		db: db,
	}
}

func (r pixRepo) AddPixKey(pixKey *model.Pix) (*model.Pix, error) {
	err := r.db.Create(pixKey).Error
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}

func (r pixRepo) FindPixKeyByID(key, kind string) (*model.Pix, error) {
	var pix model.Pix

	r.db.Preload("Account.Bank").First(&pix, "kind = ? and key = ?", kind, key)
	if pix.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}
	return &pix, nil
}
