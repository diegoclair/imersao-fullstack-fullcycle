package postgres

import (
	"fmt"

	"github.com/diegoclair/imersao/codepix/domain/entity"
	"github.com/jinzhu/gorm"
)

// pixRepo repository
type pixRepo struct {
	db *gorm.DB
}

//newPixRepo return the implementation of pixRepo interface
func newPixRepo(db *gorm.DB) *pixRepo {
	return &pixRepo{
		db: db,
	}
}

func (r pixRepo) AddPixKey(pixKey *entity.Pix) (*entity.Pix, error) {
	err := r.db.Create(pixKey).Error
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}

func (r pixRepo) FindPixByKey(key, kind string) (*entity.Pix, error) {
	var pix entity.Pix

	r.db.Preload("Account.Bank").First(&pix, "kind = ? and key = ?", kind, key)
	if pix.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}
	return &pix, nil
}
