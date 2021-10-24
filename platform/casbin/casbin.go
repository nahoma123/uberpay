package casbin

import (
	"context"
	"github.com/casbin/casbin/v2"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"log"
	"template/internal/constant"
	"template/internal/constant/errors"
	"template/internal/constant/model"
)

type CasbinAuth interface {
	AddPolicy(context.Context, model.Policy) error
	RemovePolicy(context.Context, model.Policy) error
	Policies(ctx context.Context) []model.Policy
	UpdatePolicy(ctx context.Context, parm model.PolicyUpdate) error
	GetAllRoles(c context.Context) []string
	GetCompanyPolicyByID(c context.Context, companyid string) []model.Policy
	GetAllCompaniesPolicy(c context.Context) []model.Policy
}
type casbinAuthorizer struct {
	e        *casbin.Enforcer
	validate *validator.Validate
	trans    ut.Translator
}

// NewEnforcer creates an enforcer via file or DB.
func NewEnforcer(e *casbin.Enforcer, validate *validator.Validate, trans ut.Translator) CasbinAuth {
	e.EnableAutoSave(true)
	e.LoadPolicy()
	return &casbinAuthorizer{
		e:        e,
		validate: validate,
		trans:    trans,
	}
}

//UpdatePolicy updates the policy
func (r *casbinAuthorizer) UpdatePolicy(ctx context.Context, parm model.PolicyUpdate) error {
	errV := constant.StructValidator(parm, r.validate, r.trans)
	if errV != nil {
		return errV
	}
	old := parm.Old
	new := parm.New
	isUpdated, err := r.e.UpdatePolicy([]string{old.Subject, old.Object, old.Action}, []string{new.Subject, new.Object, new.Action})
	if err != nil {
		return err
	}
	if isUpdated {
		err = r.e.SavePolicy()
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.ErrPermissionAlreadyDefined
	}
}

//AddPolicy adds new policy to casbin_rule table
func (r *casbinAuthorizer) AddPolicy(c context.Context, cas model.Policy) error {
	errV := constant.StructValidator(cas, r.validate, r.trans)
	if errV != nil {
		return errV
	}
	success, err := r.e.AddPolicy(cas.Subject, cas.Object, cas.Action)
	if err != nil {
		return err
	}
	if success {
		err = r.e.SavePolicy()
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.ErrPermissionAlreadyDefined
	}
}

//RemovePolicy removes policy from casbin rule table
func (r *casbinAuthorizer) RemovePolicy(c context.Context, cas model.Policy) error {
	success, err := r.e.RemovePolicy(cas.Subject, cas.Object, cas.Action)
	if err != nil {
		return err
	}
	if success {
		err = r.e.SavePolicy()
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	} else {
		return errors.ErrPermissionPermissionNotFound
	}
}

//Policies fetches all policies from casbin_rule table
func (r *casbinAuthorizer) Policies(c context.Context) []model.Policy {
	policies := r.e.GetPolicy()
	var permissions []model.Policy
	for i := 0; i < len(policies); i++ {
		permissions = append(permissions, model.Policy{
			Subject: policies[i][0],
			Object:  policies[i][1],
			Action:  policies[i][2],
		})
	}
	if permissions == nil {
		return []model.Policy{}
	}
	return permissions
}
func (r *casbinAuthorizer) GetAllRoles(c context.Context) []string {
	return r.e.GetAllRoles()
}
func (r *casbinAuthorizer) GetAllCompaniesPolicy(c context.Context) []model.Policy {
	AllPolicies := r.e.GetPolicy()
	var policies []model.Policy
	for _, policy := range AllPolicies {
		p := model.Policy{}
		if policy[3] != "*" {
			p.Subject = policy[0]
			p.Object = policy[1]
			p.Action = policy[2]
			p.CompanyID = policy[3]
		} else {
			continue
		}
		policies = append(policies, p)
	}
	return policies
}
func (r *casbinAuthorizer) GetCompanyPolicyByID(c context.Context, companyid string) []model.Policy {
	AllPolicies := r.e.GetPolicy()
	var policies []model.Policy
	for _, policy := range AllPolicies {
		p := model.Policy{}
		if policy[3] == companyid {
			p.Subject = policy[0]
			p.Object = policy[1]
			p.Action = policy[2]
			p.CompanyID = policy[3]
		} else {
			continue
		}
		policies = append(policies, p)
	}
	return policies
}
