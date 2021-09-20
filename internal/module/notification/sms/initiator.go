package sms

import (
	"template/internal/adapter/storage/persistence/notification/sms"
	"template/internal/constant"
	"template/internal/constant/errors"
	"template/internal/constant/model"
)

// Usecase interface contains function of business logic port for domain PushedNotification
type Usecase interface {
	SendSmsMessage(sms model.SMS) (*constant.SuccessData, *errors.ErrorModel)
	GetCountUnreadSmsMessages() int64
}

//service defines all necessary service for the domain Usecase
type service struct {
	smsPersistance sms.SmsPersistence
}

// InitializeSms creates a new object with UseCase type and implements sms services
func InitializeSms(smsP sms.SmsPersistence) Usecase {
	return &service{
		smsPersistance: smsP,
	}
}
