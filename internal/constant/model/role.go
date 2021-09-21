package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	Name      string         `gorm:"primary_key" json:"name"  validate:"required"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
