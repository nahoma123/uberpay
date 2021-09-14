package model

import (
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
}
