package storage

import "template/internal/constant/model"

type RolePersistence interface {
	Role(name string) (*model.Role, error)
	Roles() ([]model.Role, error)
	UpdateRole(role model.Role) (*model.Role, error)
	DeleteRole(name string) error
	StoreRole(role model.Role) (*model.Role, error)
	RoleExists(name string) (bool, error)
	MigrateRole() error
}
