package email

import (
	"context"
	"gorm.io/gorm"
	storage "template/internal/adapter/storage/persistence"
	"template/internal/constant/errors"
	"template/internal/constant/model"
)

//emailPersistence creates gorm pointer struct
type emailPersistence struct {
	conn *gorm.DB
}

//EmailInit creates emailPersistence object and implements EmailPersistence
func EmailInit(conn *gorm.DB) storage.EmailPersistence {
	return &emailPersistence{
		conn: conn,
	}
}

//SendEmailMessage send email message to the specified email
func (e emailPersistence) SendEmailMessage(ctx context.Context, sms model.EmailNotification) (*model.EmailNotification, error) {
	conn := e.conn.WithContext(ctx)
	err := conn.Model(&model.EmailNotification{}).Create(&sms).Error
	if err != nil {
		return nil, errors.ErrUnableToSave
	}
	return &sms, nil
}

//MigrateEmail create migration of models
func (e emailPersistence) MigrateEmail(ctx context.Context) error {
	conn := e.conn.WithContext(ctx)
	err := conn.Migrator().AutoMigrate(&model.EmailNotification{})
	if err != nil {
		return errors.ErrUnableToMigrate
	}
	return nil
}

//GetCountUnreadEmailMessages fetches all unread email message
func (e emailPersistence) GetCountUnreadEmailMessages(ctx context.Context) int64 {
	var count int64
	conn := e.conn.WithContext(ctx)
	conn.Model(&model.EmailNotification{}).Where("status = ?", "unread").Count(&count)
	return count
}
