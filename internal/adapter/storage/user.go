package storage

import (
	"template/internal/constant/model"

	uuid "github.com/satori/go.uuid"
)

type UserStorage interface {
	CreateUser(*model.User) (*model.User, error)
	GetUserById(id uuid.UUID) (*model.User, error)
	DeleteUser(id uuid.UUID) error
}
