package publisher

import (
	"fmt"
	"net/http"
	"os"
	"ride_plus/internal/adapter/http/rest/server"
	"ride_plus/internal/constant/errors"
	model "ride_plus/internal/constant/model/dbmodel"
	utils "ride_plus/internal/constant/model/init"
	"ride_plus/internal/constant/rest"
	"ride_plus/internal/module"
	"strings"

	"github.com/appleboy/go-fcm"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)

//notificationHandler implements notification servicea and goalang validator object
type notificationHandler struct {
	notificationUseCase module.NotificationUsecase
	validate            *validator.Validate
}

//NewNotificationHandler  initializes notification services and golang validator
func NewNotificationHandler(notificationUseCase module.NotificationUsecase, utils utils.Utils) server.NotificationHandler {
	return &notificationHandler{
		notificationUseCase: notificationUseCase,
		validate:            utils.GoValidator,
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
		rest.ErrorResponseJson(c, errValue, errors.StatusCodes[errors.ErrInvalidRequest])
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
		err := errors.ServiceError(errData)
		code := err.ErrorCode
		rest.ErrorResponseJson(c, err, code)
		return
	}
	rest.ErrorResponseJson(c, notification, http.StatusOK)
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
		rest.ErrorResponseJson(c, clientErr, code)
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
			rest.ErrorResponseJson(c, errValue, http.StatusBadRequest)
			return
		}
		err := errors.ServiceError(err)
		rest.ErrorResponseJson(c, err, http.StatusBadRequest)
		return
	}
	rest.ErrorResponseJson(c, *Data, http.StatusOK)
	return
}

//DeleteNotification removes  specific notification message identified by id notification
func (n notificationHandler) DeleteNotification(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	u_id, err := uuid.FromString(id)
	if err != nil {
		errValue := errors.ConvertionError()
		rest.ErrorResponseJson(c, errValue, http.StatusBadRequest)
		return
	}
	err = n.notificationUseCase.DeleteNotification(ctx, model.PushedNotification{ID: u_id})
	if err != nil {
		e := errors.ServiceError(err)
		code := e.ErrorCode
		if err != nil {
			errValue := errors.ConvertionError()
			rest.ErrorResponseJson(c, errValue, http.StatusBadRequest)
			return
		}
		rest.ErrorResponseJson(c, e, code)
		return
	}
	rest.ErrorResponseJson(c, nil, http.StatusOK)
	return
}
func (n notificationHandler) GetCountUnreadPushNotificationMessages(c *gin.Context) {
	ctx := c.Request.Context()
	count := n.notificationUseCase.GetCountUnreadPushNotificationMessages(ctx)
	rest.ErrorResponseJson(c, map[string]interface{}{"count": count}, http.StatusOK)
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
