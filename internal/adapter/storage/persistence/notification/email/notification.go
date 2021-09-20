package email

import (
	"template/internal/constant/model"
)
//SendEmailMessage send email message to the specified email
func (e emailPersistence) SendEmailMessage(sms model.EmailNotification) (interface{}, error) {
	conn := e.conn
	err := conn.Model(&model.EmailNotification{}).Create(&sms).Error
	if err != nil {
		return nil, err
	}
	return sms, nil
}
//MigrateEmail create migration of models
func (e emailPersistence) MigrateEmail() error {
	db := e.conn
	err := db.Migrator().AutoMigrate(&model.EmailNotification{})
	if err != nil {
		return err
	}
	return nil
}
//GetCountUnreadEmailMessages fetches all unread email message
func (e emailPersistence) GetCountUnreadEmailMessages() int64 {
	var count int64
	db := e.conn
	db.Model(&model.EmailNotification{}).Where("status = ?", "unread").Count(&count)
	return count
}

