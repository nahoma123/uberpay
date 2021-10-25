package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type PushedNotification struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	ApiKey    string    `json:"api_key" validate:"required"`
	Token     string    `json:"token" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	Body      string    `json:"body" validate:"required"`
	Data      string    `json:"data" validate:"required"`
	Status    string    `json:"status" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
}
type SMS struct {
	SmsID         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Password      string    `json:"password" validate:"required"`
	User          string    `json:"user" validate:"required"`
	SenderId      string    `json:"sender_id" validate:"required"`
	ApiGateWay    string    `json:"api_gate_way" validate:"required"`
	CallBackUrl   string    `json:"call_back_url" validate:"required"`
	Body          string    `json:"body" form:"body" binding:"required" validate:"required"`
	ReceiverPhone string    `json:"receiver_phone" form:"receiver_phone" binding:"required" validate:"required"`
	Status        string    `json:"status" validate:"required"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
type EmailNotification struct {
	ID        uuid.UUID `json:"email_message_id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Body      string    `json:"body"    validator:"required"`
	From      string    `json:"from"  validate:"required,email"`
	To        string    `json:"to"      validate:"required,email"`
	Subject   string    `json:"subject" validate:"required"`
	Status    string    `json:"status"  validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpDatedAt time.Time `json:"up_dated_at"`
}
