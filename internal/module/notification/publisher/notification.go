package publisher

import (
	"fmt"
	"net/http"
	"template/internal/constant"
	"template/internal/constant/errors"
	"template/internal/constant/model"
)

//Notifications returns all pushed notifications
func (s service) Notifications() (*constant.SuccessData, *errors.ErrorModel) {
	data, err := s.notificationPersistance.Notifications()
	if err != nil {
		errorData := errors.NewErrorResponse(err)
		return nil, &errorData
	}
	return &constant.SuccessData{
		Code: http.StatusOK,
		Data: data,
	}, nil

}

//PushSingleNotification creates a notification and send via valid token and api key
func (s service) PushSingleNotification(notification model.PushedNotification) (*constant.SuccessData, *errors.ErrorModel) {

	if notification.ApiKey == "" {
		errorData := errors.NewErrorResponse(errors.ErrInvalidAPIKey)
		return nil, &errorData
	}
	if notification.Token == "" {
		errorData := errors.NewErrorResponse(errors.ErrInvalidToken)
		return nil, &errorData
	}
	_, err := s.notificationPersistance.NotificationByID(notification)
	if err != nil {
		errorData := errors.NewErrorResponse(errors.ErrDataAlreayExist)
		return nil, &errorData
	}
	newnotification, err := s.notificationPersistance.PushSingleNotification(notification)
	fmt.Println("error ", err)
	if err != nil {
		errorData := errors.NewErrorResponse(err)
		return nil, &errorData
	}
	return &constant.SuccessData{
		Code: http.StatusOK,
		Data: newnotification,
	}, nil
}

//DeleteNotification removes a pushed notification from the resource
func (s service) DeleteNotification(param model.PushedNotification) (*constant.SuccessData, *errors.ErrorModel) {
	_, err := s.notificationPersistance.NotificationByID(param)
	if err != nil {
		errorData := errors.NewErrorResponse(err)
		return nil, &errorData
	}
	err = s.notificationPersistance.DeleteNotification(param)
	if err != nil {
		errorData := errors.NewErrorResponse(err)
		return nil, &errorData
	}
	return &constant.SuccessData{
		Code: http.StatusOK,
		Data: "PushedNotification Deleted",
	}, nil

}

//GetCountUnreadPushNotificationMessages returns count of unread pushed notification message
func (s service) GetCountUnreadPushNotificationMessages() int64 {
	count := s.notificationPersistance.GetCountUnreadPushNotificationMessages()
	return count
}
