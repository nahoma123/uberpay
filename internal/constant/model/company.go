package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Company struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string         `json:"name,omitempty" form:"name" validate:"required"`
	Phone     string         `json:"phone,omitempty" form:"phone" validate:"required"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
