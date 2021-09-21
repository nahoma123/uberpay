package role

import (
	"template/internal/adapter/storage/persistence"
	"template/internal/constant/model"
	err "template/internal/constant/errors"
)

type UseCase interface {
	Role(name string) (*model.Role, error)
	Roles() ([]model.Role, error)
	// UpdateRole(role model.Role) (*model.Role, error)
	DeleteRole(name string) error
	StoreRole(role model.Role) (*model.Role, error)
}
type service struct {
	rolePersistence storage.RolePersistence
}

func RoleInitialize(rolePersistence storage.RolePersistence) UseCase {
	return &service{
		rolePersistence: rolePersistence,
	}
}

func (s service) Roles() ([]model.Role, error) {

	roles, err := s.rolePersistence.Roles()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (s service) Role(name string) (*model.Role, error) {
	r, err := s.rolePersistence.Role(name)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s service) DeleteRole(name string) error {

	err := s.rolePersistence.DeleteRole(name)
	if err != nil {

		return err
	}
	return nil
}

//TODO define error
func (s service) StoreRole(role model.Role) (*model.Role, error) {
	if role.Name == "" {
		return nil, err.ErrRoleNameISEmpty
	}

	r, err := s.rolePersistence.StoreRole(role)
	if err != nil {
		return nil, err
	}
	return r, nil
}
