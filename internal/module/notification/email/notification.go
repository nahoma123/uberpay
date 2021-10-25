package email

import (
	"context"
	storage "ride_plus/internal/adapter/storage/persistence"
	"ride_plus/internal/constant"
	"ride_plus/internal/constant/model"
	"ride_plus/internal/module"
	"time"

	utils "ride_plus/internal/constant/model/init"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

//Service defines all necessary service for the domain sms
type service struct {
	emailPersist   storage.EmailPersistence
	validate       *validator.Validate
	trans          ut.Translator
	contextTimeout time.Duration
}

//Initialize  creates a new object with UseCase type
func Initialize(emailPersist storage.EmailPersistence, utils utils.Utils) module.EmailUsecase {
	return &service{
		emailPersist:   emailPersist,
		validate:       utils.GoValidator,
		trans:          utils.Translator,
		contextTimeout: utils.Timeout,
	}
}

//SendEmailMessage send email message  to one or more users
func (s service) SendEmailMessage(c context.Context, email model.EmailNotification) (*model.EmailNotification, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	errV := constant.StructValidator(email, s.validate, s.trans)
	if errV != nil {
		return nil, errV
	}
	return s.emailPersist.SendEmailMessage(ctx, email)

}

//GetCountUnreadEmailMessages returns count of unread Email notification message
func (s service) GetCountUnreadEmailMessages(c context.Context) int64 {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.emailPersist.GetCountUnreadEmailMessages(ctx)
}
func (s service) ValidSendEmail(ctx context.Context, email model.EmailNotification) error {
	return constant.StructValidator(email, s.validate, s.trans)
}
