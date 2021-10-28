package routing

import (
	"ride_plus/internal/adapter/http/rest/middleware"
	"ride_plus/internal/adapter/http/rest/server"

	"github.com/gin-gonic/gin"
)

// PublisherRoutes registers Push Notification routes
func PublisherRoutes(gp *gin.RouterGroup, authMiddleware middleware.AuthMiddleware, pub server.NotificationHandler) {
	gp.GET("/notifications", pub.GetNotifications)
	gp.GET("/notifications/unread/publish", pub.GetCountUnreadPushNotificationMessages)
	gp.POST("/notifications", pub.NotificationMiddleWare, pub.PushNotification)
	gp.DELETE("/notifications/:id", pub.DeleteNotification)
}

//EmailRoutes registers Email Message routes
func EmailRoutes(gp *gin.RouterGroup, authMiddleware middleware.AuthMiddleware, email server.EmailHandler) {
	gp.GET("/notifications/unread/email", email.GetCountUnreadEmailMessages)
	gp.POST("/notifications/email", email.EmailMessageMiddleWare, email.SendEmailMessage)
}

//SmsRoutes registers Sms Message routes
func SmsRoutes(gp *gin.RouterGroup, authMiddleware middleware.AuthMiddleware, s server.SmsHandler) {
	gp.POST("/notifications/sms", s.SmsMessageMiddleWare, s.SendSmsMessage)
	gp.GET("/notifications/unread/sms", s.GetCountUnreadSMsMessages)
}
