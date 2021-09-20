package email

import (
	"gorm.io/gorm"
	"template/internal/constant/model"
)

//EmailPersistence contains all services for notification interface
type EmailPersistence interface {
	SendEmailMessage(sms model.EmailNotification) (interface{}, error)
	GetCountUnreadEmailMessages() int64
	MigrateEmail() error
}

//emailPersistence creates gorm pointer struct
type emailPersistence struct {
	conn *gorm.DB
}
//EmailInit creates emailPersistence object and implements EmailPersistence
func EmailInit(conn *gorm.DB) EmailPersistence {
	return &emailPersistence{
		conn: conn,
	}
}
