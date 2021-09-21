package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Username   string         `json:"username" validate:"required"`
	Password   string         `json:"password" validate:"required,min=8"`
	Phone      string         `json:"phone" validate:"required"`
	FirstName  string         `json:"first_name" validate:"required"`
	MiddleName string         `json:"middle_name"`
	LastName   string         `json:"last_name" validate:"required"`
	Email      string         `gorm:"unique" json:"email" validate:"required,email"`
	RoleName   string         `json:"role_name" validate:"required"`
	CreatedAt  time.Time      `json:"created_at,omitempty"`
	UpdatedAt  time.Time      `json:"updated_at,omitempty"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

type UserCompanyRole struct {
	UserID    uuid.UUID `json:"user_id,omitempty"`
	User      *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	CompanyID uuid.UUID `json:"company_id,omitempty"`
	Company   *Company  `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	RoleName  string    `json:"role_name"`
	Role      *Role     `json:"role,omitempty" gorm:"foreignKey:RoleName"`
}
