package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type PushedNotification struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	ApiKey    string    `json:"api_key" validate:"required"`
	Token     string    `json:"token" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	Body      string    `json:"body" validate:"required"`
	Data      string    `json:"data" validate:"required"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
}
type SMS struct {
	SmsID         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Password      string    `json:"password"`
	User          string    `json:"user"`
	SenderId      string    `json:"sender_id"`
	ApiGateWay    string    `json:"api_gate_way"`
	CallBackUrl   string    `json:"call_back_url"`
	Body          string    `json:"body" form:"body" binding:"required"`
	ReceiverPhone string    `json:"receiver_phone" form:"receiver_phone" binding:"required"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
type EmailNotification struct {
	ID        uuid.UUID `json:"email_message_id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Body      string    `json:"body" validator:"required"`
	From      string    `json:"from" validator:"required"`
	To        string    `json:"to"`
	Subject   string    `json:"subject"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpDatedAt time.Time `json:"up_dated_at"`
}

type Notification struct {
	UserID         uuid.UUID `json:"user_id"`
	NotificationID uuid.UUID `json:"notification_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Status         string    `json:"status"`
}
