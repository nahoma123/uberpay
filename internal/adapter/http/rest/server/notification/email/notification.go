package email

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"template/internal/adapter/http/rest/server"
	"template/internal/constant"
	custErr "template/internal/constant/errors"
	"template/internal/constant/model"
	"template/internal/module"

	"github.com/gin-gonic/gin"
	gomail "gopkg.in/mail.v2"
)

//emailHandler implements notification service and golang validator object
type emailHandler struct {
	notificationUseCase module.EmailUsecase
	m                   *gomail.Message
}

//NewEmailHandler  initializes notification services and golang validator
func NewEmailHandler(notificationUseCase module.EmailUsecase, m *gomail.Message) server.EmailHandler {
	return &emailHandler{
		notificationUseCase: notificationUseCase,
		m:                   m,
	}
}

//MiddleWareValidateEmailMessage binds pushed notification data as json
func (n emailHandler) EmailMessageMiddleWare(c *gin.Context) {
	email := model.EmailNotification{}
	err := c.Bind(&email)
	if err != nil {
		errValue := custErr.ErrorModel{
			ErrorCode:        custErr.ErrCodes[custErr.ErrorUnableToBindJsonToStruct],
			ErrorDescription: custErr.Descriptions[custErr.ErrorUnableToBindJsonToStruct],
			ErrorMessage:     custErr.ErrorUnableToBindJsonToStruct.Error(),
		}
		constant.ResponseJson(c, errValue, custErr.StatusCodes[custErr.ErrorUnableToBindJsonToStruct])
		return
	}
	c.Set("x-email", email)
	c.Next()
}

//SendEmailMessage send email message via valid email
func (n emailHandler) SendEmailMessage(c *gin.Context) {
	ctx := c.Request.Context()
	email := c.MustGet("x-email").(model.EmailNotification)
	// TODO:01 email notification code put here
	err := n.SendEmail(ctx, email, n.m)
	fmt.Println("send email error ", err)
	if err != nil {

		if strings.Contains(err.Error(), os.Getenv("ErrSecretKey")) {
			e := strings.Replace(err.Error(), os.Getenv("ErrSecretKey"), "", 1)
			errValue := custErr.ErrorModel{
				ErrorCode:        custErr.ErrCodes[custErr.ErrInvalidField],
				ErrorDescription: custErr.Descriptions[custErr.ErrInvalidField],
				ErrorMessage:     e,
			}
			constant.ResponseJson(c, errValue, http.StatusBadRequest)
			return
		} else if errors.Is(err, custErr.ErrorUnableToConvert) {
			errValue := custErr.ConvertionError()
			constant.ResponseJson(c, errValue, custErr.ErrCodes[custErr.ErrorUnableToConvert])
			return
		}
		errValue := custErr.ErrorModel{
			ErrorCode:        custErr.ErrCodes[custErr.ErrUnableToSendEmailMessage],
			ErrorMessage:     err.Error(),
			ErrorDescription: custErr.Descriptions[custErr.ErrUnableToSendEmailMessage],
		}
		fmt.Println("errValue ", errValue)
		constant.ResponseJson(c, errValue, custErr.StatusCodes[custErr.ErrUnableToSendEmailMessage])
		return
	}
	// TODO:02 Email notification stored in the database storage
	Data, err := n.notificationUseCase.SendEmailMessage(ctx, email)
	if err != nil {
		if strings.Contains(err.Error(), os.Getenv("ErrSecretKey")) {
			e := strings.Replace(err.Error(), os.Getenv("ErrSecretKey"), "", 1)
			errValue := custErr.ErrorModel{
				ErrorCode:        custErr.ErrCodes[custErr.ErrInvalidField],
				ErrorDescription: custErr.Descriptions[custErr.ErrInvalidField],
				ErrorMessage:     e,
			}
			constant.ResponseJson(c, errValue, http.StatusBadRequest)
			return
		}
		e := custErr.NewErrorResponse(err)
		constant.ResponseJson(c, e, http.StatusBadRequest)
		return
	}
	constant.ResponseJson(c, *Data, http.StatusOK)
}

//GetCountUnreadEmailMessages return the number of unread message
func (n emailHandler) GetCountUnreadEmailMessages(c *gin.Context) {
	ctx := c.Request.Context()
	count := n.notificationUseCase.GetCountUnreadEmailMessages(ctx)
	constant.ResponseJson(c, map[string]interface{}{"count": count}, http.StatusOK)
	return
}

//SendEmail sends email message via SMTP server
func (n emailHandler) SendEmail(ctx context.Context, email model.EmailNotification, m *gomail.Message) error {
	err := n.notificationUseCase.ValidSendEmail(ctx, email)
	fmt.Println("validation error ", err)
	if err != nil {
		return err
	}
	m.SetHeader("From", email.From)
	m.SetHeader("To", email.To)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/plain", email.Body)
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return custErr.ErrorUnableToConvert
	}
	d := gomail.NewDialer(os.Getenv("SMTP_SERVER"), port, email.From, os.Getenv("SMTP_PASSWORD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
