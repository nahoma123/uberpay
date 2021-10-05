package policy

import (
	"context"
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"strings"
	"template/internal/adapter/storage/persistence"
	"template/internal/constant"
	"template/internal/constant/errors"
	"template/internal/constant/model"
	"template/internal/module"
	"time"
)


type service struct {
	permPersistence storage.PermissionPersistence
	validate        *validator.Validate
	trans           ut.Translator
	contextTimeout  time.Duration
}

func PermissionsInitialize(permPersistence storage.PermissionPersistence, validate *validator.Validate, trans ut.Translator, timeout time.Duration) module.PolicyUseCase {
	return &service{
		permPersistence: permPersistence,
		trans:           trans,
		validate:        validate,
		contextTimeout:  timeout,
	}
}
func (s service) CompanyPolicy(c context.Context, u_id uuid.UUID) (*model.CasbinRule, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	permissions, err := s.permPersistence.Policies(ctx)
	if err != nil {
		return nil, err
	}
	for _, permission := range permissions {
		url_path_array := strings.Split(permission.V1, "/")
		c_id, err := uuid.FromString(url_path_array[2])
		if err != nil {
			continue
		}
		if c_id == u_id {
			return &permission, nil
		}
	}
	return nil, errors.ErrRecordNotFound
}
func (s service) CompanyPolicies(c context.Context) ([]model.CasbinRule, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	cr := []model.CasbinRule{}
	permissions, err := s.permPersistence.Policies(ctx)
	if err != nil {
		return nil, err
	}
	for _, permission := range permissions {
		url_path_array := strings.Split(permission.V1, "/")
		_, err = uuid.FromString(url_path_array[2])
		if err != nil {
			continue
		}
		cr = append(cr, permission)
	}
	return cr, nil
}
func (s service) Policies(c context.Context) ([]model.CasbinRule, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.permPersistence.Policies(ctx)
}
func (s service) Policy(c context.Context, id uint) (*model.CasbinRule, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.permPersistence.Policy(ctx, id)
}
func (s service) DeletePolicy(c context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.permPersistence.RemovePolicy(ctx, id)
}
func (s service) UpdatePolicy(c context.Context, p model.CasbinRule) (*model.CasbinRule, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	errV := constant.StructValidator(p, s.validate, s.trans)
	if errV != nil {
		return nil, errV
	}
	return s.permPersistence.UpdatePolicy(ctx, p)
}
func (s service) StorePolicy(c context.Context, p model.CasbinRule) (*model.CasbinRule, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	errV := constant.StructValidator(p, s.validate, s.trans)
	fmt.Println("error validation ", errV)
	if errV != nil {
		return nil, errV
	}
	return s.permPersistence.StorePolicy(ctx, p)
}
