package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

//Base entity model to be reused in others structs
type Base struct {
	ID        string    `json:"id,omitempty" valid:"uuid"`
	CreatedAt time.Time `json:"created_at,omitempty" valid:"-" mapper:"-"`
	UpdatedAt time.Time `json:"updated_at,omitempty" valid:"-" mapper:"-"`
}
