package role

import (
	"context"
	storage "ride_plus/internal/adapter/storage/persistence"
	"ride_plus/internal/constant"
	"ride_plus/internal/constant/model"
	utils "ride_plus/internal/constant/model/init"
	"time"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type UseCase interface {
	Role(c context.Context, name string) (*model.Role, error)
	Roles(c context.Context) ([]model.Role, error)
	DeleteRole(c context.Context, name string) error
	StoreRole(c context.Context, role model.Role) (*model.Role, error)
}
type service struct {
	rolePersistence storage.RolePersistence
	validate        *validator.Validate
	trans           ut.Translator
	contextTimeout  time.Duration
}

func RoleInitialize(rolePersistence storage.RolePersistence, utils utils.Utils) UseCase {
	return &service{
		rolePersistence: rolePersistence,
		validate:        utils.GoValidator,
		trans:           utils.Translator,
		contextTimeout:  utils.Timeout,
	}
}
func (s service) Roles(c context.Context) ([]model.Role, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	roles, err := s.rolePersistence.Roles(ctx)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
func (s service) Role(c context.Context, name string) (*model.Role, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	role, err := s.rolePersistence.Role(ctx, name)
	if err != nil {
		return nil, err
	}
	return role, nil
}
func (s service) DeleteRole(c context.Context, name string) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	err := s.rolePersistence.DeleteRole(ctx, name)
	if err != nil {
		return err
	}
	return nil
}
func (s service) StoreRole(c context.Context, r model.Role) (*model.Role, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	errV := constant.StructValidator(r, s.validate, s.trans)
	if errV != nil {
		return nil, errV
	}
	role, err := s.rolePersistence.StoreRole(ctx, r)
	if err != nil {
		return nil, err
	}
	return role, nil
}
