package publisher

import (
	"context"
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	storage "template/internal/adapter/storage/persistence"
	"template/internal/constant"
	"template/internal/constant/model"
	"template/internal/module"
	"time"
)
//Service defines all necessary service for the domain sms
type service struct {
	notifyPersist   storage.NotificationPersistence
	validate       *validator.Validate
	trans          ut.Translator
	contextTimeout time.Duration
}

//Initialize  creates a new object with UseCase type
func Initialize(notp storage.NotificationPersistence, validate *validator.Validate, trans ut.Translator, timeout time.Duration) module.NotificationUsecase {
	return &service{
		notifyPersist: notp,
		validate:       validate,
		trans:          trans,
		contextTimeout: timeout,
	}
}
//Notifications returns all pushed notifications
func (s service) Notifications(c context.Context) ([]model.PushedNotification, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	data, err := s.notifyPersist.Notifications(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil

}

//PushSingleNotification creates a notification and send via valid token and api key
func (s service) PushSingleNotification(c context.Context, notification model.PushedNotification) (*model.PushedNotification, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	errV := constant.StructValidator(notification, s.validate, s.trans)
	fmt.Println("val error ", errV)
	if errV != nil {
		return nil, errV
	}
	return s.notifyPersist.PushSingleNotification(ctx, notification)

}

//DeleteNotification removes a pushed notification from the resource
func (s service) DeleteNotification(c context.Context, param model.PushedNotification) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	_, err := s.notifyPersist.NotificationByID(ctx, param)
	if err != nil {
		return err
	}
	return s.notifyPersist.DeleteNotification(ctx, param)

}

//GetCountUnreadPushNotificationMessages returns count of unread pushed notification message
func (s service) GetCountUnreadPushNotificationMessages(c context.Context) int64 {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.notifyPersist.GetCountUnreadPushNotificationMessages(ctx)

}
