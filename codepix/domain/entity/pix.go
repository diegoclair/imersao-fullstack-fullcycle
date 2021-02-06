package entity

import (
	"errors"
	"time"

	"github.com/Nhanderu/brdoc"
	"github.com/asaskevich/govalidator"
	"github.com/twinj/uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// Pix kind values
const (
	PixKindEmail string = "email"
	PixKindCPF   string = "cpf"
)

// Status standard values
const (
	PixStatusActive   string = "active"
	PixStatusInactive string = "inactive"
)

//Pix entity model
type Pix struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key"  valid:"notnull"`
	AccountID string   `gorm:"column:account_id;type:uuid;not null" valid:"-"`
	Account   *Account `json:"account"  valid:"-"`
	Status    string   `json:"status"  valid:"notnull"`
}

func (pix *Pix) isValid() error {

	if pix.Kind != PixKindEmail && pix.Kind != PixKindCPF {
		return errors.New("Invalid type of key")
	}

	if pix.Status != PixStatusActive && pix.Status != PixStatusInactive {
		return errors.New("Invalid status")
	}

	if pix.Kind == PixKindCPF && !brdoc.IsCPF(pix.Key) {
		return errors.New("Invalid cpf")
	}

	_, err := govalidator.ValidateStruct(pix)
	if err != nil {
		return err
	}
	return nil
}

//NewPix return a new Pix model
func NewPix(kind string, account *Account, key string) (*Pix, error) {
	pix := Pix{
		Kind:    kind,
		Key:     key,
		Account: account,
		Status:  PixStatusActive,
	}

	pix.ID = uuid.NewV4().String()
	pix.CreatedAt = time.Now()

	err := pix.isValid()
	if err != nil {
		return nil, err
	}
	return &pix, nil
}
