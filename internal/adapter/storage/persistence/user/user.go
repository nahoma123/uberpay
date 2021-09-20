package user

import (
	"template/internal/constant/model"
	uuid "github.com/satori/go.uuid"
)

type UserStorage interface {
	User(param model.User) (*model.User, error)
	UserCompanyRole(param model.UserCompanyRole) (*model.UserCompanyRole, error)
	CreateUser(companyID uuid.UUID,user *model.User) (*model.User, error)
	CreateSystemUser(user *model.User) (*model.User, error)
	GetUserById(id uuid.UUID) (*model.User, error)
	DeleteUser(id uuid.UUID) error
	GetUsers() ([]model.User, error)
}
