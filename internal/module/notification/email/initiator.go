package email

import (
	"template/internal/adapter/storage/persistence/notification/email"
	"template/internal/constant"
	"template/internal/constant/errors"
	"template/internal/constant/model"
)

// Usecase interface contains function of business logic port for domain PushedNotification
type Usecase interface {
	SendEmailMessage(sms model.EmailNotification) (*constant.SuccessData, *errors.ErrorModel)
	GetCountUnreadEmailMessages()(int64)
}
//service defines all necessary service for the domain Email Messaging
type service struct {
	emailPersistance  email.EmailPersistence
}
// InitializeEmail creates a new object with UseCase type and implements services
func InitializeEmail(emailP  email.EmailPersistence) Usecase {
	return &service{
		emailPersistance: emailP,
	}
}
