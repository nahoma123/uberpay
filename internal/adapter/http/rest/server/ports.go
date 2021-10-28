package server

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(c *gin.Context)
}

//SmsHandler contains all handler interfaces
type SmsHandler interface {
	SmsMessageMiddleWare(c *gin.Context)
	SendSmsMessage(c *gin.Context)
	GetCountUnreadSMsMessages(c *gin.Context)
}

// UserHandler contans a function of handlers for the domian file
type UserHandler interface {
	UserByID(c *gin.Context)
	Users(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	StoreUser(c *gin.Context)
	AddUserToCompany(c *gin.Context)
	RemoveUserFromCompany(c *gin.Context)
	GetCompanyUsers(c *gin.Context)
	GetCompanyUserByID(c *gin.Context)
	UserMiddleWare(c *gin.Context)
	RegisterUser(c *gin.Context)
}

// PolicyHandler contains a function of handlers for the domain file
type PolicyHandler interface {
	PolicyMiddleWare(c *gin.Context)
	StorePolicy(c *gin.Context)
	RemovePolicy(c *gin.Context)
	Policies(c *gin.Context)
	UpdatePolicy(c *gin.Context)
	GetCompanyPolicyByID(c *gin.Context)
	GetAllCompaniesPolicy(c *gin.Context)
}

//NotificationHandler contains all handler interfaces
type NotificationHandler interface {
	NotificationMiddleWare(c *gin.Context)
	GetNotifications(c *gin.Context)
	PushNotification(c *gin.Context)
	DeleteNotification(c *gin.Context)
	GetCountUnreadPushNotificationMessages(c *gin.Context)
}

//EmailHandler contains all email handler interfaces
type EmailHandler interface {
	EmailMessageMiddleWare(c *gin.Context)
	SendEmailMessage(c *gin.Context)
	GetCountUnreadEmailMessages(c *gin.Context)
}
type CompanyHandler interface {
	CompanyByID(c *gin.Context)
	Companies(c *gin.Context)
	StoreCompany(c *gin.Context)
	UpdateCompany(c *gin.Context)
	DeleteCompany(c *gin.Context)
	StoreCompanyImage(ctx *gin.Context)
	UpdateCompanyImage(ctx *gin.Context)
	CompanyImages(ctx *gin.Context)
	SaveFile(f *multipart.FileHeader, format, path string, rwidth, rheiht uint) error
}
