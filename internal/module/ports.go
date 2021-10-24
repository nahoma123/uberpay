package module

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"template/internal/constant/model"
)

// UserUsecase interface contains function of business logic for domian USer
type UserUsecase interface {
	UserByID(ctx context.Context, param model.User) (*model.User, error)
	Users(ctx context.Context) ([]model.User, error)
	UpdateUser(ctx context.Context, user model.User) (*model.User, error)
	DeleteUser(ctx context.Context, param model.User) error
	StoreUser(ctx context.Context, user model.User) (*model.User, error)
	UserExists(ctx context.Context, param model.User) (bool, error)
	PhoneExists(ctx context.Context, phone string) (bool, error)
	EmailExists(ctx context.Context, email string) (bool, error)
	AddUserToCompany(ctx context.Context, parm model.CompanyUser) error
	RemoveUser(ctx context.Context, parm model.CompanyUser) error
	GetCompanyUsers(ctx context.Context, companyID uuid.UUID) ([]model.CompanyUser, error)
	GetCompanyUserByID(ctx context.Context, user_id uuid.UUID) (*model.CompanyUser, error)
}
type PolicyUseCase interface {
	Policy(c context.Context, id uint) (*model.CasbinRule, error)
	Policies(c context.Context) ([]model.CasbinRule, error)
	UpdatePolicy(c context.Context, rule model.CasbinRule) (*model.CasbinRule, error)
	DeletePolicy(c context.Context, id uint) error
	StorePolicy(c context.Context, role model.CasbinRule) (*model.CasbinRule, error)
	CompanyPolicy(c context.Context, u_id uuid.UUID) ([]model.Policy, error)
	CompanyPolicies(ctx context.Context) ([]model.Policy, error)
}

//SmsUsecase interface contains function of business logic port for domain PushedNotification
type SmsUsecase interface {
	SendSmsMessage(c context.Context, sms model.SMS) (*model.SMS, error)
	GetCountUnreadSmsMessages(c context.Context) int64
}

// NotificationUsecase interface contains function of business logic port for domain PushedNotification
type NotificationUsecase interface {
	Notifications(c context.Context) ([]model.PushedNotification, error)
	PushSingleNotification(c context.Context, notification model.PushedNotification) (*model.PushedNotification, error)
	DeleteNotification(c context.Context, param model.PushedNotification) error
	GetCountUnreadPushNotificationMessages(c context.Context) int64
}

// EmailUsecase interface contains function of business logic port for domain PushedNotification
type EmailUsecase interface {
	SendEmailMessage(c context.Context, sms model.EmailNotification) (*model.EmailNotification, error)
	ValidSendEmail(ctx context.Context, email model.EmailNotification) error
	GetCountUnreadEmailMessages(c context.Context) int64
}

// CompanyUsecase interface contains function of business logic for domain company
type CompanyUsecase interface {
	CompanyByID(ctx context.Context, param model.Company) (*model.Company, error)
	Companies(ctx context.Context) ([]model.Company, error)
	StoreCompany(ctx context.Context, param model.Company) (*model.Company, error)
	UpdateCompany(ctx context.Context, param model.Company) (*model.Company, error)
	DeleteCompany(ctx context.Context, param model.Company) error
	StoreCompanyImage(ctx context.Context, images model.CompanyImage) (*model.CompanyImage, error)
	UpdateCompanyImage(c context.Context, param model.CompanyImage) (*model.CompanyImage, error)
	CompanyImages(ctx context.Context) ([]model.CompanyImage, error)
	CompanyExists(ctx context.Context, param model.Company) (bool, error)
}

type LoginUseCase interface {
	Login(c context.Context, username, password string) (*model.LoginResponse, error)
	GetClaims(token string) (*model.UserClaims, error)
}
