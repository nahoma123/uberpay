package sms

import (
	"fmt"
	"net/http"
	"template/internal/constant"
	"template/internal/constant/errors"
	"template/internal/constant/model"
)

//SendSmsMessage send sms message via phone numbers
func (s service) SendSmsMessage(sms model.SMS) (*constant.SuccessData, *errors.ErrorModel) {
	if sms.ApiGateWay == "" {
		errorData := errors.NewErrorResponse(errors.ErrInvalidAPIKey)
		return nil, &errorData
	}
	if sms.CallBackUrl == "" {
		errorData := errors.NewErrorResponse(errors.ErrorInvalidCallBackUrl)
		return nil, &errorData
	}
	newnotification, err := s.smsPersistance.SendSmsMessage(sms)
	fmt.Println("err perst ", err)
	if err != nil {
		errorData := errors.NewErrorResponse(errors.ErrUnableToSendSmsMessage)
		return nil, &errorData
	}
	return &constant.SuccessData{
		Code: http.StatusOK,
		Data: newnotification,
	}, nil

}

//GetCountUnreadSmsMessages returns count of unread SMS notification message
func (s service) GetCountUnreadSmsMessages() int64 {
	count := s.smsPersistance.GetCountUnreadSmsMessages()
	return count
}
