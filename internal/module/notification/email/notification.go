package email

import (
	"net/http"
	"template/internal/constant"
	"template/internal/constant/errors"
	"template/internal/constant/model"
)

//SendEmailMessage send email message  to one or more users
func (s service) SendEmailMessage(email model.EmailNotification) (*constant.SuccessData, *errors.ErrorModel) {
	if email.From == "" {
		errorData := errors.NewErrorResponse(errors.ErrorInvalidSenderAddress)
		return nil, &errorData
	}
	if len(email.To) == 0 {
		errorData := errors.NewErrorResponse(errors.ErrorInvalidRecieverAddress)
		return nil, &errorData
	}
	if email.Body == "" {
		errorData := errors.NewErrorResponse(errors.ErrorInvalidBody)
		return nil, &errorData
	}
	emailnotification, err := s.emailPersistance.SendEmailMessage(email)
	if err != nil {
		errorData := errors.NewErrorResponse(errors.ErrUnableToSendEmailMessage)
		return nil, &errorData
	}
	return &constant.SuccessData{
		Code: http.StatusOK,
		Data: emailnotification,
	}, nil
}

//GetCountUnreadEmailMessages returns count of unread Email notification message
func (s service) GetCountUnreadEmailMessages() int64 {
	count := s.emailPersistance.GetCountUnreadEmailMessages()
	return count
}
