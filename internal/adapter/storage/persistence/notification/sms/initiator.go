package sms

import (
    "gorm.io/gorm"
	"template/internal/constant/model"
)
//SmsPersistence contains all services for notification interface
type SmsPersistence interface {
	SendSmsMessage(sms model.SMS) (interface{}, error)
	GetCountUnreadSmsMessages()(int64)
	MigrateSms() error
}
//smsPersistence creates gorm pointer struct
type smsPersistence struct {
	conn *gorm.DB
}
//SmsInit creates smsPersistence object and implements SmsPersistence
func SmsInit(conn *gorm.DB) SmsPersistence {
	return &smsPersistence{conn: conn,}
}