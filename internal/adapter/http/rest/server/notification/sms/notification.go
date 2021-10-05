package sms

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"os"
	"strings"
	"template/internal/adapter/http/rest/server"
	"template/internal/constant"
	"template/internal/constant/errors"
	"template/internal/constant/model"
	"template/internal/module"
)



//smsHandler implements sms servicea and golang validator object
type smsHandler struct {
	smsUseCase  module.SmsUsecase
	validate   *validator.Validate
}

//NewSmsHandler  initializes notification services and golang validator
func NewSmsHandler(s  module.SmsUsecase, valid *validator.Validate) server.SmsHandler {
	return &smsHandler{smsUseCase: s, validate: valid}
}

//MiddleWareValidateSmsMessage binds sms data SMS struct
func (n smsHandler) MiddleWareValidateSmsMessage(c *gin.Context) {
	sms := model.SMS{}
	err := c.Bind(&sms)
	sms.User = os.Getenv("SMS_USER")
	sms.SenderId = os.Getenv("SMS_SENDER")
	sms.ApiGateWay = os.Getenv("SMS_API_GATE_WAY")
	sms.CallBackUrl = os.Getenv("SMS_CALLBACK_URL")
	sms.Password = os.Getenv("SMS_PASSWORD")
	if err != nil {
		errValue := errors.ErrorModel{
			ErrorCode:        errors.ErrCodes[errors.ErrInvalidRequest],
			ErrorDescription: errors.Descriptions[errors.ErrInvalidRequest],
			ErrorMessage:     errors.ErrInvalidRequest.Error(),
		}
		constant.ResponseJson(c, errValue, errors.StatusCodes[errors.ErrInvalidRequest])
		return
	}
	c.Set("x-sms", sms)
	c.Next()
}

//SendSmsMessage  sends sms message to a user via phone number
func (n smsHandler) SendSmsMessage(c *gin.Context) {
	ctx := c.Request.Context()
	sms := c.MustGet("x-sms").(model.SMS)
	// TODO:01 sms notification code put here
	_, err := SendSmsMessage(sms)
	fmt.Println("error sms ", err)
	if err != nil {
		errValue := errors.ErrorModel{
			ErrorCode:        errors.ErrCodes[errors.ErrUnableToSendSmsMessage],
			ErrorDescription: errors.Descriptions[errors.ErrUnableToSendSmsMessage],
			ErrorMessage:     errors.ErrUnableToSendSmsMessage.Error(),
		}
		constant.ResponseJson(c, errValue, errors.StatusCodes[errors.ErrUnableToSendSmsMessage])
		return
	}
	// TODO:02 sms notification data store in the database put here
	data, err := n.smsUseCase.SendSmsMessage(ctx, sms)
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
		constant.ResponseJson(c, errors.NewErrorResponse(err), errors.ErrCodes[err])
		return
	}
	constant.ResponseJson(c, *data, http.StatusOK)
	return
}

//GetCountUnreadSMsMessages counts unread sms message
func (n smsHandler) GetCountUnreadSMsMessages(c *gin.Context) {
	ctx := c.Request.Context()
	count := n.smsUseCase.GetCountUnreadSmsMessages(ctx)
	constant.ResponseJson(c, map[string]interface{}{"count": count}, http.StatusOK)
}

//SendSmsMessage sends sms message via phone number
func SendSmsMessage(sms model.SMS) (interface{}, error) {
	reqString := fmt.Sprintf(`
		{
			"type":"text",
			"auth":{"username":"%s", "password":"%s"},
			"sender":"%s",
			"receiver":"%s",
			"dcs":"GSM",
			"text":"%s",
			"dlrMask":3,
			"dlrUrl":"%s"
        }
	`, sms.User, sms.Password, sms.SenderId, sms.ReceiverPhone, sms.Body, sms.CallBackUrl)
	requestBody := strings.NewReader(reqString)
	// post some data
	res, err := http.Post(sms.ApiGateWay, "application/json; charset=UTF-8", requestBody)
	fmt.Println("error2 sms ", err)

	if err != nil {
		return nil, errors.ErrUnableToSendSmsMessage
	}
	if !(res.StatusCode >= 200 && res.StatusCode < 300) {
		return nil, errors.ErrUnableToSendSmsMessage
	}
	// read response data
	var smsResponseData interface{}
	err = json.NewDecoder(res.Body).Decode(&smsResponseData)
	if err != nil {
		return nil, errors.ErrUnableToSendSmsMessage
	}
	err = res.Body.Close()
	if err != nil {
		return nil, errors.ErrUnableToSendSmsMessage
	}
	return smsResponseData, nil
}
