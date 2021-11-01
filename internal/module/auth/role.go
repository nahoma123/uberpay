package auth

import (
	"context"
	storage "ride_plus/internal/adapter/storage/persistence"
	"ride_plus/internal/constant"
	"ride_plus/internal/constant/errors"
	model "ride_plus/internal/constant/model/dbmodel"
	utils "ride_plus/internal/constant/model/init"
	"time"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type RoleUseCase interface {
	Role(c context.Context, name string) (*model.Role, error)
	Roles(c context.Context) ([]model.Role, error)
	DeleteRole(c context.Context, name string) error
	StoreRole(c context.Context, role model.Role) (*model.Role, *errors.ErrorModel)
}
type roleService struct {
	rolePersistence storage.RolePersistence
	validate        *validator.Validate
	trans           ut.Translator
	contextTimeout  time.Duration
}

func RoleInitialize(rolePersistence storage.RolePersistence, utils utils.Utils) RoleUseCase {
	return &roleService{
		rolePersistence: rolePersistence,
		validate:        utils.GoValidator,
		trans:           utils.Translator,
		contextTimeout:  utils.Timeout,
	}
}
func (s roleService) Roles(c context.Context) ([]model.Role, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	roles, err := s.rolePersistence.Roles(ctx)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
func (s roleService) Role(c context.Context, name string) (*model.Role, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	role, err := s.rolePersistence.Role(ctx, name)
	if err != nil {
		return nil, err
	}
	return role, nil
}
func (s roleService) DeleteRole(c context.Context, name string) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	err := s.rolePersistence.DeleteRole(ctx, name)
	if err != nil {
		return err
	}
	return nil
}

func (s roleService) StoreRole(c context.Context, r model.Role) (*model.Role, *errors.ErrorModel) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	// verify input from transport layer
	errV := constant.VerifyInput(r, s.validate, s.trans)
	if errV != nil {
		return nil, errV
	}
	role, err := s.rolePersistence.StoreRole(ctx, r)
	if err != nil {
		return nil, err
	}
	return role, nil
}
