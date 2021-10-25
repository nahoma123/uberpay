package initiator

import (
	email3 "template/internal/adapter/http/rest/server/notification/email"
	publisher3 "template/internal/adapter/http/rest/server/notification/publisher"
	sms3 "template/internal/adapter/http/rest/server/notification/sms"
	"template/internal/adapter/storage/persistence/notification/email"
	"template/internal/adapter/storage/persistence/notification/publisher"
	"template/internal/adapter/storage/persistence/notification/sms"
	utils "template/internal/constant/model/init"
	routing2 "template/internal/glue/routing"
	email2 "template/internal/module/notification/email"
	publisher2 "template/internal/module/notification/publisher"
	sms2 "template/internal/module/notification/sms"

	"github.com/gin-gonic/gin"
	gomail "gopkg.in/mail.v2"
)

// initialize notification domain
func InitNotification(utils utils.Utils, router *gin.RouterGroup) {
	//notification persistence
	emailPersistent := email.EmailInit(utils.Conn)
	smsPersistent := sms.SmsInit(utils.Conn)
	publisherPersistent := publisher.NotificationInit(utils.Conn)

	m := gomail.NewMessage()

	//notification services
	emailUsecase := email2.Initialize(emailPersistent, utils)
	smsUsecase := sms2.Initialize(smsPersistent, utils)
	publisherUsecase := publisher2.Initialize(publisherPersistent, utils)

	emailHandler := email3.NewEmailHandler(emailUsecase, m)
	smsHandler := sms3.NewSmsHandler(smsUsecase, utils)
	publisherHandler := publisher3.NewNotificationHandler(publisherUsecase, utils)

	//notification
	routing2.EmailRoutes(router, emailHandler)
	routing2.SmsRoutes(router, smsHandler)
	routing2.PublisherRoutes(router, publisherHandler)
}
