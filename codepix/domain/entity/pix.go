package entity

import (
	"errors"
	"time"

	"github.com/Nhanderu/brdoc"
	"github.com/diegoclair/imersao/codepix/domain"
	"github.com/diegoclair/imersao/codepix/infrastructure/validate"

	"github.com/twinj/uuid"
)

//Pix entity model
type Pix struct {
	Base      `validate:"required"`
	Keytype   string   `json:"key_type" validate:"required"`
	Key       string   `json:"key"  validate:"required"`
	AccountID string   `gorm:"column:account_id;type:uuid;not null"`
	Account   *Account `json:"account"`
	Status    string   `json:"status"  validate:"required"`
}

func (pix *Pix) isValid() error {

	if pix.Keytype != domain.PixKeytypeEmail && pix.Keytype != domain.PixKeytypeCPF {
		return errors.New("Invalid type of key")
	}

	if pix.Status != domain.PixStatusActive && pix.Status != domain.PixStatusInactive {
		return errors.New("Invalid status")
	}

	if pix.Keytype == domain.PixKeytypeCPF && !brdoc.IsCPF(pix.Key) {
		return errors.New("Invalid cpf")
	}

	return validate.Struct(pix, "Pix")
}

//NewPix return a new Pix model
func NewPix(keyType string, account *Account, key string) (*Pix, error) {
	pix := Pix{
		Keytype:   keyType,
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
