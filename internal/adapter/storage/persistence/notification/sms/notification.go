package sms

import (
	"template/internal/constant/model"
)

//SendSmsMessage send sms message text to user phone
func (s smsPersistence) SendSmsMessage(sms model.SMS) (interface{}, error) {
	conn := s.conn
	err := conn.Model(&model.SMS{}).Create(&sms).Error
	if err != nil {
		return nil, err
	}
	return sms, nil
}

//MigrateSms create migration of models
func (s smsPersistence) MigrateSms() error {
	db := s.conn
	err := db.Migrator().AutoMigrate(&model.SMS{})
	if err != nil {
		return err
	}
	return nil
}

//GetCountUnreadSmsMessages fetches all unread sms messages
func (s smsPersistence) GetCountUnreadSmsMessages() int64 {
	var count int64
	db := s.conn
	db.Model(&model.SMS{}).Where("status = ?", "unread").Count(&count)
	return count

}
