package storage

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"template/internal/constant/model"
)

//RolePersistence all services of role
type RolePersistence interface {
	Role(ctx context.Context, name string) (*model.Role, error)
	Roles(ctx context.Context) ([]model.Role, error)
	UpdateRole(ctx context.Context, role model.Role) (*model.Role, error)
	DeleteRole(ctx context.Context, name string) error
	StoreRole(ctx context.Context, role model.Role) (*model.Role, error)
	RoleExists(ctx context.Context, name string) (bool, error)
	MigrateRole() error
}

//PermissionPersistence contains all services for PermissionPersistence interface
type PermissionPersistence interface {
	Policy(ctx context.Context, id uint) (*model.CasbinRule, error)
	Policies(ctx context.Context) ([]model.CasbinRule, error)
	UpdatePolicy(ctx context.Context, role model.CasbinRule) (*model.CasbinRule, error)
	RemovePolicy(ctx context.Context, id uint) error
	StorePolicy(ctx context.Context, role model.CasbinRule) (*model.CasbinRule, error)
	MigratePolicy(c context.Context) error
}

//EmailPersistence contains all services for notification interface
type EmailPersistence interface {
	SendEmailMessage(ctx context.Context, sms model.EmailNotification) (*model.EmailNotification, error)
	GetCountUnreadEmailMessages(ctx context.Context) int64
	MigrateEmail(ctx context.Context) error
}

//NotificationPersistence contains all services for notification interface
type NotificationPersistence interface {
	Notifications(ctx context.Context) ([]model.PushedNotification, error)
	NotificationByID(ctx context.Context, parm model.PushedNotification) (*model.PushedNotification, error)
	PushSingleNotification(ctx context.Context, activity model.PushedNotification) (*model.PushedNotification, error)
	DeleteNotification(ctx context.Context, param model.PushedNotification) error
	GetCountUnreadPushNotificationMessages(ctx context.Context) int64
	MigrateNotification(ctx context.Context) error
}

//SmsPersistence contains all services for notification interface
type SmsPersistence interface {
	SendSmsMessage(ctx context.Context, sms model.SMS) (*model.SMS, error)
	GetCountUnreadSmsMessages(ctx context.Context) int64
	MigrateSms(ctx context.Context) error
}

//CompanyPersistence contains all services for Company interface
type CompanyPersistence interface {
	CompanyByID(ctx context.Context, param model.Company) (*model.Company, error)
	Companies(ctx context.Context) ([]model.Company, error)
	StoreCompany(ctx context.Context, param model.Company) (*model.Company, error)
	UpdateCompany(ctx context.Context, param model.Company) (*model.Company, error)
	DeleteCompany(ctx context.Context, param model.Company) error
	CompanyExists(ctx context.Context, param model.Company) (bool, error)
	MigrateCompany(ctx context.Context) error
}
type Handler interface {
	Authorizer(e *casbin.Enforcer) gin.HandlerFunc
	Login(c *gin.Context)
}