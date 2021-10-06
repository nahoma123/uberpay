package routing

import (
	"github.com/gin-gonic/gin"
	"template/internal/adapter/http/rest/server"
)

// PublisherRoutes registers Push Notification routes
func PublisherRoutes(gp *gin.RouterGroup, pub server.NotificationHandler) {
	gp.GET("/notifications", pub.GetNotifications)
	gp.GET("/notifications/unread/publish", pub.GetCountUnreadPushNotificationMessages)
	gp.POST("/notifications", pub.NotificationMiddleWare, pub.PushNotification)
	gp.DELETE("/notifications/:id", pub.DeleteNotification)
}

//EmailRoutes registers Email Message routes
func EmailRoutes(gp *gin.RouterGroup, email server.EmailHandler) {
	gp.GET("/notifications/unread/email", email.GetCountUnreadEmailMessages)
	gp.POST("/notifications/email", email.EmailMessageMiddleWare, email.SendEmailMessage)
}

//SmsRoutes registers Sms Message routes
func SmsRoutes(gp *gin.RouterGroup, s server.SmsHandler) {
	gp.POST("/notifications/sms", s.SmsMessageMiddleWare, s.SendSmsMessage)
	gp.GET("/notifications/unread/sms", s.GetCountUnreadSMsMessages)
}
