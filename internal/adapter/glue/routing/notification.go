package routing

import (
	"github.com/gin-gonic/gin"
	"template/internal/adapter/http/rest/server/notification/email"
	"template/internal/adapter/http/rest/server/notification/publisher"
	"template/internal/adapter/http/rest/server/notification/sms"
)

// PublisherRoutes registers Push Notification routes
func PublisherRoutes(gp *gin.RouterGroup, pub publisher.NotificationHandler) {
	gp.GET("/notifications", pub.GetNotifications)
	gp.GET("/notifications/unread/publish", pub.GetCountUnreadPushNotificationMessages)
	gp.POST("/notifications", pub.PushNotification)
	gp.DELETE("/notifications/:id", pub.DeleteNotification)
}

//EmailRoutes registers Email Message routes
func EmailRoutes(gp *gin.RouterGroup, email email.EmailHandler) {
	gp.GET("/notification/unread/email", email.GetCountUnreadEmailMessages)
	gp.POST("/notification/email", email.SendEmailMessage)
}

//SmsRoutes registers Sms Message routes
func SmsRoutes(gp *gin.RouterGroup, s sms.SmsHandler) {
	gp.POST("/notification/sms", s.SendSmsMessage)
	gp.GET("/notification/unread/sms", s.GetCountUnreadSMsMessages)
}
