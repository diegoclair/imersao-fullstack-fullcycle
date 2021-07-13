package entity

import (
	"errors"
	"time"

	"github.com/diegoclair/imersao/codepix/domain"
	"github.com/diegoclair/imersao/codepix/util/validate"
	"github.com/paemuri/brdoc"

	"github.com/twinj/uuid"
)

//Pix entity model
type Pix struct {
	Base      `validate:"required"`
	KeyType   string `json:"key_type" gorm:"type:varchar(20)" validate:"required"`
	Key       string `json:"key" gorm:"type:varchar(255)" validate:"required"`
	AccountID string `gorm:"column:account_id;type:uuid;not null" validate:"omitempty,uuid"`
	Account   *Account
	Status    string `json:"status" gorm:"type:varchar(20)" validate:"required"`
}

func (pix *Pix) isValid() error {

	if pix.KeyType != domain.PixKeytypeEmail && pix.KeyType != domain.PixKeytypeCPF {
		return errors.New("Invalid type of key")
	}

	if pix.Status != domain.PixStatusActive && pix.Status != domain.PixStatusInactive {
		return errors.New("Invalid status")
	}

	if pix.KeyType == domain.PixKeytypeCPF && !brdoc.IsCPF(pix.Key) {
		return errors.New("Invalid cpf")
	}

	return validate.Struct(pix, "Pix")
}

//NewPix return a new Pix model
func NewPix(keyType string, account *Account, key string) (*Pix, error) {
	pix := Pix{
		KeyType:   keyType,
		Key:       key,
		Account:   account,
		AccountID: account.ID,
		Status:    domain.PixStatusActive,
	}

	pix.ID = uuid.NewV4().String()
	pix.CreatedAt = time.Now()

	err := pix.isValid()
	if err != nil {
		return nil, err
	}
	return &pix, nil
}
