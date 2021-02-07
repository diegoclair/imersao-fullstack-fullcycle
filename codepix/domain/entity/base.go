package entity

import (
	"time"
)

//Base entity model to be reused in others structs
type Base struct {
	ID        string    `json:"id" validate:"required,uuid4"`
	CreatedAt time.Time `json:"created_at" mapper:"-"`
	UpdatedAt time.Time `json:"updated_at" mapper:"-"`
}
