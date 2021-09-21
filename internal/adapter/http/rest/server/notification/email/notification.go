package email

import (
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	gomail "gopkg.in/mail.v2"
	"net/http"
	"os"
	"strconv"
	"template/internal/constant"
	"template/internal/constant/errors"
	"template/internal/constant/model"
	"template/internal/module/notification/email"
)

//EmailHandler contains all email handler interfaces
type EmailHandler interface {
	MiddleWareValidateEmailMessage(c *gin.Context)
	SendEmailMessage(c *gin.Context)
	GetCountUnreadEmailMessages(c *gin.Context)
}

//emailHandler implements notification service and golang validator object
type emailHandler struct {
	notificationUseCase email.Usecase
	validate            *validator.Validate
	m                   *gomail.Message
}

//NewEmailHandler  initializes notification services and golang validator
func NewEmailHandler(em email.Usecase, valid *validator.Validate, mes *gomail.Message) EmailHandler {
	return &emailHandler{
		notificationUseCase: em,
		validate:            valid,
		m:                   mes,
	}
}

//MiddleWareValidateEmailMessage binds pushed notification data as json
func (n emailHandler) MiddleWareValidateEmailMessage(c *gin.Context) {
	email := model.EmailNotification{}
	err := c.Bind(&email)
	if err != nil {
		errValue := errors.ErrorModel{
			ErrorCode:        strconv.Itoa(errors.StatusCodes[errors.ErrInvalidRequest]),
			ErrorDescription: errors.Descriptions[errors.ErrInvalidRequest],
			ErrorMessage:     errors.ErrInvalidRequest.Error(),
		}
		constant.ResponseJson(c, errValue, errors.StatusCodes[errors.ErrInvalidRequest])
		return
	}
	errV := constant.StructValidator(email, n.validate)
	if errV != nil {
		constant.ResponseJson(c, errV, errors.StatusCodes[errors.ErrorUnableToBindJsonToStruct])
		return
	}
	c.Set("x-email", email)
	c.Next()
}

//SendEmailMessage send email message via valid email
func (n emailHandler) SendEmailMessage(c *gin.Context) {
	email := c.MustGet("x-email").(model.EmailNotification)
	// TODO:01 email notification code put here
	err := SendEmail(email, n.m)
	fmt.Println("error ", err)
	if err != nil {
		if err == errors.ErrorUnableToConvert {
			errValue := errors.ErrorModel{
				ErrorCode:        strconv.Itoa(errors.StatusCodes[errors.ErrorUnableToConvert]),
				ErrorDescription: errors.Descriptions[errors.ErrorUnableToConvert],
				ErrorMessage:     errors.ErrorUnableToConvert.Error(),
			}
			constant.ResponseJson(c, errValue, errors.StatusCodes[errors.ErrorUnableToConvert])
			return
		}
		errValue := errors.ErrorModel{
			ErrorCode:        strconv.Itoa(errors.StatusCodes[errors.ErrUnableToSendEmailMessage]),
			ErrorDescription: errors.Descriptions[errors.ErrUnableToSendEmailMessage],
			ErrorMessage:     errors.ErrUnableToSendEmailMessage.Error(),
		}
		constant.ResponseJson(c, errValue, errors.StatusCodes[errors.ErrUnableToSendEmailMessage])
		return
	}
	// TODO:02 Email notification stored in the database storage
	Data, errData := n.notificationUseCase.SendEmailMessage(email)
	if errData != nil {
		code, _ := strconv.Atoi(errData.ErrorCode)
		constant.ResponseJson(c, *errData, code)
		return
	}
	constant.ResponseJson(c, *Data, Data.Code)
	return
}

//GetCountUnreadEmailMessages return the number of unread message
func (n emailHandler) GetCountUnreadEmailMessages(c *gin.Context) {
	count := n.notificationUseCase.GetCountUnreadEmailMessages()
	constant.ResponseJson(c, map[string]interface{}{"count": count}, http.StatusOK)

}

//SendEmail sends email message via SMTP server
func SendEmail(email model.EmailNotification, m *gomail.Message) error {
	m.SetHeader("From", email.From)
	m.SetHeader("To", email.To)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/plain", email.Body)
	//os.Setenv("SMTP_PORT", "587")
	//os.Setenv("SMTP_PASSWORD", "yideg2378")
	//os.Setenv("SMTP_SERVER", "smtp.gmail.com")
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return errors.ErrorUnableToConvert
	}
	d := gomail.NewDialer(os.Getenv("SMTP_SERVER"), port, email.From, os.Getenv("SMTP_PASSWORD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
