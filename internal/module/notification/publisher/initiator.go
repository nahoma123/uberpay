package publisher

import (
	"template/internal/adapter/storage/persistence/notification/publisher"
	"template/internal/constant"
	"template/internal/constant/errors"
	"template/internal/constant/model"
)

// Usecase interface contains function of business logic port for domain PushedNotification
type Usecase interface {
	Notifications() (*constant.SuccessData,*errors.ErrorModel)
	PushSingleNotification(notification model.PushedNotification) (*constant.SuccessData, *errors.ErrorModel)
	DeleteNotification(param model.PushedNotification) (*constant.SuccessData, *errors.ErrorModel)
	GetCountUnreadPushNotificationMessages()(int64)
}
//service defines all necessary service for the domain PushedNotification
type service struct {
	notificationPersistance publisher.NotificationPersistence
}
// InitializePublisher creates a new object with UseCase type and implements services
func InitializePublisher(notificationPersistance publisher.NotificationPersistence) Usecase {
	return &service{
		notificationPersistance: notificationPersistance,
	}
}
