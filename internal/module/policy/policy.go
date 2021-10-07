package policy

import (
	"context"
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"template/internal/adapter/storage/persistence"
	"template/internal/constant"
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

func PolicyInitialize(permPersistence storage.PermissionPersistence, validate *validator.Validate, trans ut.Translator, timeout time.Duration) module.PolicyUseCase {
	return &service{
		permPersistence: permPersistence,
		trans:           trans,
		validate:        validate,
		contextTimeout:  timeout,
	}
}
func (s service) CompanyPolicy(c context.Context, u_id uuid.UUID) ([]model.Policy, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	policies :=[]model.Policy{}
	AllPolicy, err := s.permPersistence.Policies(ctx)
	if err != nil {
		return nil, err
	}

	for _, policy := range AllPolicy {
		p:=model.Policy{}
		if policy.V3 == u_id.String() {
			p.Subject=policy.V0
			p.Object=policy.V1
			p.Action=policy.V2
			p.CompanyID=policy.V3
		}
		policies=append(policies,p)
	}
	return policies, nil
}
func (s service) CompanyPolicies(c context.Context) ([]model.Policy, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	policies :=[]model.Policy{}
	AllPolicy, err := s.permPersistence.Policies(ctx)
	if err != nil {
		return nil, err
	}

	for _, policy := range AllPolicy {
		p:=model.Policy{}
		if policy.V3 !="*" {
			p.Subject=policy.V0
			p.Object=policy.V1
			p.Action=policy.V2
			p.CompanyID=policy.V3
		}
		policies=append(policies,p)
	}
	return policies, nil
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
