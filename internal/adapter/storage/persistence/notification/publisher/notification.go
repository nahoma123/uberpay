package publisher

import (
	"gorm.io/gorm"
	"template/internal/constant/errors"
	"template/internal/constant/model"
)
//Notifications fetches all available notification
func (n notificationPersistence) Notifications() ([]model.PushedNotification, error) {
	conn := n.conn
	notications := []model.PushedNotification{}

	err := conn.Model(&model.PushedNotification{}).Find(&notications).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrRecordNotFound
		}
		return nil, errors.ErrorUnableToFetch
	}
	return notications, nil
}
//NotificationByID find pushed notification identified by its id
func (n notificationPersistence) NotificationByID(parm model.PushedNotification) (*model.PushedNotification, error) {
	conn := n.conn
	notification := &model.PushedNotification{}

	err := conn.Model(&model.PushedNotification{}).Where(&parm).First(notification).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrIDNotFound
		}
		return nil, errors.ErrorUnableToFetch
	}
	return notification, nil
}
//PushSingleNotification pushes notification to a single device using firebase cloudy message of api key
func (n notificationPersistence) PushSingleNotification(notification model.PushedNotification) (*model.PushedNotification, error) {
	conn := n.conn

	err := conn.Model(&model.PushedNotification{}).Create(&notification).Error
	if err != nil {
		if err == gorm.ErrRegistered {
			return nil,errors.ErrorUnableToCreate
		}
		return nil, errors.ErrInvalidRequest
	}
	return &notification, nil
}
//DeleteNotification removes notification from the resource center (storage)
func (n notificationPersistence) DeleteNotification(param model.PushedNotification) error {
	conn := n.conn

	err := conn.Model(&model.PushedNotification{}).Where(&param).Delete(&param).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrIDNotFound
		}
		return errors.ErrUnableToDelete
	}
	return nil
}
//MigrateNotification create migration of models
func (n notificationPersistence) MigrateNotification() error {
	db := n.conn
	err := db.Migrator().AutoMigrate(&model.PushedNotification{})
	if err != nil {
		return err
	}
	return nil
}
//GetCountUnreadPushNotificationMessages  gets number of unread pushed notification
func (n notificationPersistence) GetCountUnreadPushNotificationMessages() int64 {
	var count int64
	db := n.conn
	db.Model(&model.PushedNotification{}).Where("status = ?", "unread").Count(&count)
	return count
}
