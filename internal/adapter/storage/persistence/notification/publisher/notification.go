package publisher

import (
	"context"
	storage "ride_plus/internal/adapter/storage/persistence"
	"ride_plus/internal/constant/errors"
	model "ride_plus/internal/constant/model/dbmodels"

	"gorm.io/gorm"
)

//notificationPersistence creates gorm pointer struct
type notificationPersistence struct {
	conn *gorm.DB
}

//NotificationInit creates notificationPersistence object and implements NotificationPersistence
func NotificationInit(conn *gorm.DB) storage.NotificationPersistence {
	return &notificationPersistence{
		conn: conn,
	}
}

//Notifications fetches all available notification
func (n notificationPersistence) Notifications(ctx context.Context) ([]model.PushedNotification, error) {
	conn := n.conn.WithContext(ctx)
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
func (n notificationPersistence) NotificationByID(ctx context.Context, parm model.PushedNotification) (*model.PushedNotification, error) {
	conn := n.conn.WithContext(ctx)
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
func (n notificationPersistence) PushSingleNotification(ctx context.Context, notification model.PushedNotification) (*model.PushedNotification, error) {
	conn := n.conn.WithContext(ctx)
	err := conn.Model(&model.PushedNotification{}).Create(&notification).Error
	if err != nil {
		if err == gorm.ErrRegistered {
			return nil, errors.ErrorUnableToCreate
		}
		return nil, errors.ErrInvalidRequest
	}
	return &notification, nil
}

//DeleteNotification removes notification from the resource center (storage)
func (n notificationPersistence) DeleteNotification(ctx context.Context, param model.PushedNotification) error {
	conn := n.conn.WithContext(ctx)
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
func (n notificationPersistence) MigrateNotification(ctx context.Context) error {
	conn := n.conn.WithContext(ctx)
	err := conn.Migrator().AutoMigrate(&model.PushedNotification{})
	if err != nil {
		return err
	}
	return nil
}

//GetCountUnreadPushNotificationMessages  gets number of unread pushed notification
func (n notificationPersistence) GetCountUnreadPushNotificationMessages(ctx context.Context) int64 {
	var count int64
	conn := n.conn.WithContext(ctx)
	conn.Model(&model.PushedNotification{}).Where("status = ?", "unread").Count(&count)
	return count
}
