package authorization

import (
	"context"
	"log"
	"ride_plus/internal/constant/errors"
	model "ride_plus/internal/constant/model/dbmodels"

	"github.com/casbin/casbin/v2"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

type CasbinAuth interface {
	AddPolicy(context.Context, model.Policy) error
	RemovePolicy(context.Context, model.Policy) error
	Policies(ctx context.Context) []model.Policy
	UpdatePolicy(ctx context.Context, parm model.PolicyUpdate) error
	GetAllRoles(c context.Context) []string
	GetCompanyPolicyByID(c context.Context, companyid string) []model.Policy
	GetAllCompaniesPolicy(c context.Context) []model.Policy
	IsAuthorized(c context.Context, userId string, companyId string, role string, subject string, object string) (bool, error)
}
type casbinAuthorizer struct {
	enforcer *casbin.Enforcer
}

// NewEnforcer creates an enforcer via file or DB.
func NewEnforcer(conn *gorm.DB, model string) CasbinAuth {
	adapter, err := gormadapter.NewAdapterByDB(conn)
	if err != nil {
		log.Fatal("error ", err)
	}

	enforcer, err := casbin.NewEnforcer(model, adapter)
	if err != nil {
		log.Fatal("error ", err)
	}

	enforcer.EnableAutoSave(true)
	enforcer.LoadPolicy()
	return &casbinAuthorizer{
		enforcer: enforcer,
	}
}

func (r *casbinAuthorizer) IsAuthorized(c context.Context, userId string, companyId string, role string, subject string, object string) (bool, error) {
	return r.enforcer.Enforce(userId, companyId, object, subject)
}

//UpdatePolicy updates the policy
func (r *casbinAuthorizer) UpdatePolicy(ctx context.Context, parm model.PolicyUpdate) error {
	old := parm.Old
	new := parm.New
	isUpdated, err := r.enforcer.UpdatePolicy([]string{old.Subject, old.Object, old.Action}, []string{new.Subject, new.Object, new.Action})
	if err != nil {
		return err
	}
	if isUpdated {
		err = r.enforcer.SavePolicy()
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

	success, err := r.enforcer.AddPolicy(cas.Subject, cas.Object, cas.Action)
	if err != nil {
		return err
	}
	if success {
		err = r.enforcer.SavePolicy()
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
	success, err := r.enforcer.RemovePolicy(cas.Subject, cas.Object, cas.Action)
	if err != nil {
		return err
	}
	if success {
		err = r.enforcer.SavePolicy()
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
	policies := r.enforcer.GetPolicy()
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
	return r.enforcer.GetAllRoles()
}
func (r *casbinAuthorizer) GetAllCompaniesPolicy(c context.Context) []model.Policy {
	AllPolicies := r.enforcer.GetPolicy()
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
	AllPolicies := r.enforcer.GetPolicy()
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
