package publisher

import (
	"github.com/appleboy/go-fcm"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/satori/go.uuid"
	"net/http"
	"os"
	"strconv"
	"template/internal/constant"
	"template/internal/constant/errors"
	"template/internal/constant/model"
	"template/internal/module/notification/publisher"
)

//NotificationHandler contains all handler interfaces
type NotificationHandler interface {
	MiddleWareValidateNotification(c *gin.Context)
	GetNotifications(c *gin.Context)
	PushNotification(c *gin.Context)
	DeleteNotification(c *gin.Context)
	GetCountUnreadPushNotificationMessages(c *gin.Context)
}

//notificationHandler implements notification servicea and goalang validator object
type notificationHandler struct {
	notificationUseCase publisher.Usecase
	validate            *validator.Validate
}

//NewNotificationHandler  initializes notification services and golang validator
func NewNotificationHandler(notfCase publisher.Usecase, valid *validator.Validate) NotificationHandler {
	return &notificationHandler{
		notificationUseCase: notfCase,
		validate:            valid,
	}
}

//MiddleWareValidateNotification binds pushed notification data as json
func (n notificationHandler) MiddleWareValidateNotification(c *gin.Context) {
	notification := model.PushedNotification{}
	err := c.Bind(&notification)
	notification.Token=os.Getenv("TOKEN")
	notification.ApiKey=os.Getenv("APIKEY")
	if err != nil {
		errValue := errors.ErrorModel{
			ErrorCode:        strconv.Itoa(errors.StatusCodes[errors.ErrInvalidRequest]),
			ErrorDescription: errors.Descriptions[errors.ErrInvalidRequest],
			ErrorMessage:     errors.ErrInvalidRequest.Error(),
		}
		constant.ResponseJson(c, errValue, errors.StatusCodes[errors.ErrInvalidRequest])
		return
	}
	errV := constant.StructValidator(notification, n.validate)
	if errV != nil {
		constant.ResponseJson(c, errV, errors.StatusCodes[errors.ErrorUnableToBindJsonToStruct])
		return
	}
	c.Set("x-notification", notification)
	c.Next()
}

//GetNotifications return all notification
func (n notificationHandler) GetNotifications(c *gin.Context) {
	successData, errData := n.notificationUseCase.Notifications()
	if errData != nil {
		code, err := strconv.Atoi(errData.ErrorCode)
		if err != nil {
			errValue := errors.ErrorModel{
				ErrorCode:        strconv.Itoa(errors.StatusCodes[errors.ErrorUnableToConvert]),
				ErrorDescription: errors.Descriptions[errors.ErrorUnableToConvert],
				ErrorMessage:     errors.ErrorUnableToConvert.Error(),
			}
			constant.ResponseJson(c, errValue, errors.StatusCodes[errors.ErrorUnableToConvert])
			return
		}
		constant.ResponseJson(c, *errData, code)
		return
	}
	constant.ResponseJson(c, *successData, successData.Code)

}

//PushNotification pushes message via valid device token
func (n notificationHandler) PushNotification(c *gin.Context) {
	notification := c.MustGet("x-notification").(model.PushedNotification)
	// TODO:01 push notification code put here
	data := notification
	msg := &fcm.Message{
		To:           data.Data,
		Data:         map[string]interface{}{"greet": data.Data, "api_key": data.ApiKey},
		Notification: &fcm.Notification{Title: data.Title, Body: data.Body},
	}
	//create clients from the fcm instance of api key
	client, clientErr := NewClientNotification(msg)
	if clientErr != nil {
		code, err := strconv.Atoi(clientErr.ErrorCode)
		if err != nil {
			errValue := errors.ErrorModel{
				ErrorCode:        strconv.Itoa(errors.StatusCodes[errors.ErrorUnableToConvert]),
				ErrorDescription: errors.Descriptions[errors.ErrorUnableToConvert],
				ErrorMessage:     errors.ErrorUnableToConvert.Error(),
			}
			constant.ResponseJson(c, errValue, errors.StatusCodes[errors.ErrorUnableToConvert])
			return
		}
		constant.ResponseJson(c, *clientErr, code)
		return
	}
	constant.ResponseJson(c, *client, client.Success)
	// TODO:02 store push notification here
	Data, err := n.notificationUseCase.PushSingleNotification(notification)
	if err != nil {
		code, _ := strconv.Atoi(err.ErrorCode)
		constant.ResponseJson(c, *err, code)
		return
	}
	constant.ResponseJson(c, *Data, Data.Code)
}

//DeleteNotification removes  specific notification message identified by id notification
func (n notificationHandler) DeleteNotification(c *gin.Context) {
	id := c.Param("id")
	u_id, err := uuid.FromString(id)
	if err != nil {
		errValue := errors.ErrorModel{
			ErrorCode:        strconv.Itoa(errors.StatusCodes[errors.ErrorUnableToConvert]),
			ErrorDescription: errors.Descriptions[errors.ErrorUnableToConvert],
			ErrorMessage:     errors.ErrorUnableToConvert.Error(),
		}
		constant.ResponseJson(c, errValue, errors.StatusCodes[errors.ErrorUnableToConvert])
		return
	}

	err = constant.ValidateVariable(u_id, n.validate)
	if err != nil {
		errValue := errors.ErrorModel{
			ErrorCode:        strconv.Itoa(errors.StatusCodes[errors.ErrInvalidVariable]),
			ErrorDescription: errors.Descriptions[errors.ErrInvalidVariable],
			ErrorMessage:     errors.ErrInvalidVariable.Error(),
		}
		constant.ResponseJson(c, errValue, errors.StatusCodes[errors.ErrInvalidVariable])
		return
	}

	successData, errData := n.notificationUseCase.DeleteNotification(model.PushedNotification{ID: u_id})
	if errData != nil {
		code, err := strconv.Atoi(errData.ErrorCode)
		if err != nil {
			errValue := errors.ErrorModel{
				ErrorCode:        strconv.Itoa(errors.StatusCodes[errors.ErrorUnableToConvert]),
				ErrorDescription: errors.Descriptions[errors.ErrorUnableToConvert],
				ErrorMessage:     errors.ErrorUnableToConvert.Error(),
			}
			constant.ResponseJson(c, errValue, code)
			return
		}
		return
	}
	constant.ResponseJson(c, *successData, successData.Code)
}
func (n notificationHandler) GetCountUnreadPushNotificationMessages(c *gin.Context) {
	count := n.notificationUseCase.GetCountUnreadPushNotificationMessages()
	constant.ResponseJson(c, map[string]interface{}{"count": count}, http.StatusOK)

}

//NewClientNotification send push notification using firebase cloudy message apikey and  device  valid token
func NewClientNotification(msg *fcm.Message) (*fcm.Response, *errors.ErrorModel) {
	// Create a FCM client to send the message.
	client, err := fcm.NewClient(msg.Data["api_key"].(string))
	if err != nil {
		errValue := errors.ErrorModel{
			ErrorCode:        strconv.Itoa(errors.ErrCodes[errors.ErrInvalidClient]),
			ErrorDescription: errors.Descriptions[errors.ErrInvalidClient],
			ErrorMessage:     errors.ErrInvalidClient.Error(),
		}
		return nil, &errValue
	}
	// Send the message and receive the response without retries.
	fcmResponse, err := client.Send(msg)
	if err != nil {
		errValue := errors.ErrorModel{
			ErrorCode:        strconv.Itoa(errors.ErrCodes[errors.ErrUnauthorizedClient]),
			ErrorDescription: errors.Descriptions[errors.ErrUnauthorizedClient],
			ErrorMessage:     errors.ErrUnauthorizedClient.Error(),
		}
		return nil, &errValue
	}
	return fcmResponse, nil
}
