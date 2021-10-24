package sms

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
	smsPersist     storage.SmsPersistence
	validate       *validator.Validate
	trans          ut.Translator
	contextTimeout time.Duration
}

//Initialize  creates a new object with UseCase type
func Initialize(smsPersist storage.SmsPersistence, validate *validator.Validate, trans ut.Translator, timeout time.Duration) module.SmsUsecase {
	return &service{
		smsPersist:     smsPersist,
		validate:       validate,
		trans:          trans,
		contextTimeout: timeout,
	}
}

//SendSmsMessage send sms message via phone numbers
func (s service) SendSmsMessage(c context.Context, sms model.SMS) (*model.SMS, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	errV := constant.StructValidator(sms, s.validate, s.trans)

	if errV != nil {
		return nil, errV
	}
	return s.smsPersist.SendSmsMessage(ctx, sms)

}

//GetCountUnreadSmsMessages returns count of unread SMS notification message
func (s service) GetCountUnreadSmsMessages(c context.Context) int64 {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.smsPersist.GetCountUnreadSmsMessages(ctx)

}
