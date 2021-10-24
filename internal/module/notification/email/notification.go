package email

import (
	"context"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	storage "template/internal/adapter/storage/persistence"
	"template/internal/constant"
	"template/internal/constant/model"
	"template/internal/module"
	"time"
)

//Service defines all necessary service for the domain sms
type service struct {
	emailPersist   storage.EmailPersistence
	validate       *validator.Validate
	trans          ut.Translator
	contextTimeout time.Duration
}

//Initialize  creates a new object with UseCase type
func Initialize(em storage.EmailPersistence, validate *validator.Validate, trans ut.Translator, timeout time.Duration) module.EmailUsecase {
	return &service{
		emailPersist:   em,
		validate:       validate,
		trans:          trans,
		contextTimeout: timeout,
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
