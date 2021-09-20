package publisher

import (
	"gorm.io/gorm"
	"template/internal/constant/model"
)

//NotificationPersistence contains all services for notification interface
type NotificationPersistence interface {
	Notifications() ([]model.PushedNotification, error)
	NotificationByID(parm model.PushedNotification) (*model.PushedNotification, error)
	PushSingleNotification(activity model.PushedNotification) (*model.PushedNotification, error)
	DeleteNotification(param model.PushedNotification) error
	GetCountUnreadPushNotificationMessages() int64
	MigrateNotification() error
}

//notificationPersistence creates gorm pointer struct
type notificationPersistence struct {
	conn *gorm.DB
}
//NotificationInit creates notificationPersistence object and implements NotificationPersistence
func NotificationInit(conn *gorm.DB) NotificationPersistence {
	return &notificationPersistence{
		conn: conn,
	}
}