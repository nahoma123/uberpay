package publisher

import (
	"fmt"
	"github.com/appleboy/go-fcm"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/satori/go.uuid"
	"net/http"
	"os"
	"strings"
	"template/internal/adapter/http/rest/server"
	"template/internal/constant"
	"template/internal/constant/errors"
	"template/internal/constant/model"
	"template/internal/module"
)



//notificationHandler implements notification servicea and goalang validator object
type notificationHandler struct {
	notificationUseCase module.NotificationUsecase
	validate            *validator.Validate
}

//NewNotificationHandler  initializes notification services and golang validator
func NewNotificationHandler(notfCase module.NotificationUsecase, valid *validator.Validate) server.NotificationHandler {
	return &notificationHandler{
		notificationUseCase: notfCase,
		validate:            valid,
	}
}

//MiddleWareValidateNotification binds pushed notification data as json
func (n notificationHandler) NotificationMiddleWare(c *gin.Context) {
	notification := model.PushedNotification{}
	err := c.Bind(&notification)
	notification.ApiKey = os.Getenv("APIKEY")
	notification.Token = os.Getenv("TOKEN")
	fmt.Println("error middle ware  ", err)
	if err != nil {
		errValue := errors.ErrorModel{
			ErrorCode:        errors.ErrCodes[errors.ErrInvalidRequest],
			ErrorDescription: errors.Descriptions[errors.ErrInvalidRequest],
			ErrorMessage:     errors.ErrInvalidRequest.Error(),
		}
		constant.ResponseJson(c, errValue, errors.StatusCodes[errors.ErrInvalidRequest])
		return
	}
	c.Set("x-notification", notification)
	c.Next()
}

//GetNotifications return all notification
func (n notificationHandler) GetNotifications(c *gin.Context) {
	ctx := c.Request.Context()
	notification, errData := n.notificationUseCase.Notifications(ctx)
	if errData != nil {
		err := errors.NewErrorResponse(errData)
		code := err.ErrorCode
		constant.ResponseJson(c, err, code)
		return
	}
	constant.ResponseJson(c, notification, http.StatusOK)
	return
}

//PushNotification pushes message via valid device token
func (n notificationHandler) PushNotification(c *gin.Context) {
	ctx := c.Request.Context()
	notification := c.MustGet("x-notification").(model.PushedNotification)
	// TODO:01 push notification code put here
	data := notification
	msg := &fcm.Message{
		To:           data.Data,
		Data:         map[string]interface{}{"greet": data.Data, "api_key": data.ApiKey},
		Notification: &fcm.Notification{Title: data.Title, Body: data.Body},
	}
	//create clients from the fcm instance of api key
	_, clientErr := NewClientNotification(msg)
	if clientErr != nil {
		code := http.StatusUnauthorized
		constant.ResponseJson(c, clientErr, code)
		return
	}

	// TODO:02 store push notification here
	Data, err := n.notificationUseCase.PushSingleNotification(ctx, notification)
	if err != nil {
		if strings.Contains(err.Error(), os.Getenv("ErrSecretKey")) {
			e := strings.Replace(err.Error(), os.Getenv("ErrSecretKey"), "", 1)
			errValue := errors.ErrorModel{
				ErrorCode:        errors.ErrCodes[errors.ErrInvalidField],
				ErrorDescription: errors.Descriptions[errors.ErrInvalidField],
				ErrorMessage:     e,
			}
			constant.ResponseJson(c, errValue, http.StatusBadRequest)
			return
		}
		err := errors.NewErrorResponse(err)
		constant.ResponseJson(c, err, http.StatusBadRequest)
		return
	}
	constant.ResponseJson(c, *Data, http.StatusOK)
	return
}

//DeleteNotification removes  specific notification message identified by id notification
func (n notificationHandler) DeleteNotification(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	u_id, err := uuid.FromString(id)
	if err != nil {
		errValue := errors.ConvertionError()
		constant.ResponseJson(c, errValue, http.StatusBadRequest)
		return
	}
	err = n.notificationUseCase.DeleteNotification(ctx, model.PushedNotification{ID: u_id})
	if err != nil {
		e := errors.NewErrorResponse(err)
		code := e.ErrorCode
		if err != nil {
			errValue := errors.ConvertionError()
			constant.ResponseJson(c, errValue, http.StatusBadRequest)
			return
		}
		constant.ResponseJson(c, e, code)
		return
	}
	constant.ResponseJson(c, nil, http.StatusOK)
	return
}
func (n notificationHandler) GetCountUnreadPushNotificationMessages(c *gin.Context) {
	ctx := c.Request.Context()
	count := n.notificationUseCase.GetCountUnreadPushNotificationMessages(ctx)
	constant.ResponseJson(c, map[string]interface{}{"count": count}, http.StatusOK)
	return
}

//NewClientNotification send push notification using firebase cloudy message apikey and  device  valid token
func NewClientNotification(msg *fcm.Message) (*fcm.Response, *errors.ErrorModel) {
	// Create a FCM client to send the message.
	client, err := fcm.NewClient(msg.Data["api_key"].(string))
	if err != nil {
		errValue := errors.ErrorModel{
			ErrorCode:        errors.ErrCodes[errors.ErrInvalidClient],
			ErrorDescription: errors.Descriptions[errors.ErrInvalidClient],
			ErrorMessage:     errors.ErrInvalidClient.Error(),
		}
		return nil, &errValue
	}
	// Send the message and receive the response without retries.
	fcmResponse, err := client.Send(msg)
	if err != nil {
		errValue := errors.ErrorModel{
			ErrorCode:        errors.ErrCodes[errors.ErrUnauthorizedClient],
			ErrorDescription: errors.Descriptions[errors.ErrUnauthorizedClient],
			ErrorMessage:     errors.ErrUnauthorizedClient.Error(),
		}
		return nil, &errValue
	}
	return fcmResponse, nil
}
