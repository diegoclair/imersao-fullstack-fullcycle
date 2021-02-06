package entity

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Nhanderu/brdoc"
	"github.com/diegoclair/go_utils-lib/v2/validstruct"
	"github.com/twinj/uuid"
)

// Pix kind values
const (
	PixKeytypeEmail string = "email"
	PixKeytypeCPF   string = "cpf"
)

// Status standard values
const (
	PixStatusActive   string = "active"
	PixStatusInactive string = "inactive"
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

	if pix.Keytype != PixKeytypeEmail && pix.Keytype != PixKeytypeCPF {
		return errors.New("Invalid type of key")
	}

	if pix.Status != PixStatusActive && pix.Status != PixStatusInactive {
		return errors.New("Invalid status")
	}

	if pix.Keytype == PixKeytypeCPF && !brdoc.IsCPF(pix.Key) {
		return errors.New("Invalid cpf")
	}

	err := validstruct.ValidateStruct(pix)
	if err != nil {
		validationErrors := err.Causes().([]string)
		fmt.Println("Error to validate pix entity struct")
		for i := range validationErrors {
			fmt.Println(strconv.Itoa(i+1) + " - " + validationErrors[i])
		}

		return fmt.Errorf(fmt.Sprintf("%v", validationErrors))
	}
	return nil
}

//NewPix return a new Pix model
func NewPix(keyType string, account *Account, key string) (*Pix, error) {
	pix := Pix{
		Keytype:   keyType,
		Key:       key,
		Account:   account,
		AccountID: account.ID,
		Status:    PixStatusActive,
	}

	pix.ID = uuid.NewV4().String()
	pix.CreatedAt = time.Now()

	err := pix.isValid()
	if err != nil {
		return nil, err
	}
	return &pix, nil
}
