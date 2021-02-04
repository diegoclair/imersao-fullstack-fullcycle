package model

import (
	"errors"
	"time"

	"github.com/Nhanderu/brdoc"
	"github.com/asaskevich/govalidator"
	"github.com/twinj/uuid"
)

// Pix key kind values
const (
	PixKeyKindEmail string = "email"
	PixKeyKindCPF   string = "cpf"
)

// Status standard values
const (
	PixKeyStatusActive   string = "active"
	PixKeyStatusInactive string = "inactive"
)

//PixKeyRepositoryInterface is the interface for PixKeyModel
type PixKeyRepositoryInterface interface {
	AddPixKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByID(key, kind string) (*PixKey, error)
}

//PixKey entity model
type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key"  valid:"notnull"`
	AccountID string   `gorm:"column:account_id;type:uuid;not null" valid:"-"`
	Account   *Account `json:"account"  valid:"-"`
	Status    string   `json:"status"  valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {

	if pixKey.Kind != PixKeyKindEmail && pixKey.Kind != PixKeyKindCPF {
		return errors.New("Invalid type of key")
	}

	if pixKey.Status != PixKeyStatusActive && pixKey.Status != PixKeyStatusInactive {
		return errors.New("Invalid status")
	}

	if pixKey.Kind == PixKeyKindCPF && !brdoc.IsCPF(pixKey.Key) {
		return errors.New("Invalid cpf")
	}

	_, err := govalidator.ValidateStruct(pixKey)
	if err != nil {
		return err
	}
	return nil
}

//NewPixKey return a new PixKey model
func NewPixKey(kind string, account *Account, key string) (*PixKey, error) {
	pixKey := PixKey{
		Kind:    kind,
		Key:     key,
		Account: account,
		Status:  PixKeyStatusActive,
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}
	return &pixKey, nil
}
