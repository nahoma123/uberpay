package sms

import (
	"context"
	storage "ride_plus/internal/adapter/storage/persistence"
	"ride_plus/internal/constant/errors"
	"ride_plus/internal/constant/model"

	"gorm.io/gorm"
)

//smsPersistence creates gorm pointer struct
type smsPersistence struct {
	conn *gorm.DB
}

//SmsInit creates smsPersistence object and implements SmsPersistence
func SmsInit(conn *gorm.DB) storage.SmsPersistence {
	return &smsPersistence{conn: conn}
}

//SendSmsMessage send sms message text to user phone
func (s smsPersistence) SendSmsMessage(ctx context.Context, sms model.SMS) (*model.SMS, error) {
	conn := s.conn.WithContext(ctx)
	err := conn.Model(&model.SMS{}).Create(&sms).Error
	if err != nil {
		return nil, errors.ErrUnableToSave
	}
	return &sms, nil
}

//MigrateSms create migration of models
func (s smsPersistence) MigrateSms(ctx context.Context) error {
	conn := s.conn.WithContext(ctx)
	err := conn.Migrator().AutoMigrate(&model.SMS{})
	if err != nil {
		return errors.ErrUnableToMigrate
	}
	return nil
}

//GetCountUnreadSmsMessages fetches all unread sms messages
func (s smsPersistence) GetCountUnreadSmsMessages(ctx context.Context) int64 {
	var count int64
	conn := s.conn.WithContext(ctx)
	conn.Model(&model.SMS{}).Where("status = ?", "unread").Count(&count)
	return count
}
