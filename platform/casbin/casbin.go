package casbin

import (
	"log"
	"template/internal/constant/model"
	"template/internal/constant/errors"


	"github.com/casbin/casbin/v2"
)

type CasbinAuth interface {
	AddPolicy(model.Permision) error
	RemovePolicy(model.Permision)error
	Policies() ([]model.Permision)
}

type casbinAuthorizer struct {
	e *casbin.Enforcer
}

// NewCasbin is for initialize the handler
func NewCasbin(e *casbin.Enforcer) CasbinAuth {
	e.EnableAutoSave(true)
	e.LoadPolicy()
	return &casbinAuthorizer{
		e,
	}
}
func (r *casbinAuthorizer) AddPolicy( cas model.Permision) error{
    success,err:=r.e.AddPolicy(cas.Subject,cas.Object,cas.Action)
	if err!=nil {
		log.Println("--err adding")
		// log.Println(err)
       return err
	}
	if success{
		log.Println("success")
		err=r.e.SavePolicy()
		if err!=nil {
			log.Println("--error saving")
			log.Println(err)
			return err
		}
		return nil
	}else{
		log.Println("not success")
		return errors.ErrPermissionAlreadyDefined
	}
}
func (r *casbinAuthorizer) RemovePolicy( cas model.Permision) error{
    success,err:=r.e.RemovePolicy(cas.Subject,cas.Object,cas.Action)
	if err!=nil {
		log.Println("--err adding")
		// log.Println(err)
       return err
	}
	if success{
		log.Println("success")
		err=r.e.SavePolicy()
		if err!=nil {
			log.Println("--error saving")
			log.Println(err)
			return err
		}
		return nil
	}else{
		log.Println("not success")
		return errors.ErrPermissionPermissionNotFound
	}
}
func (r *casbinAuthorizer)	Policies() ([]model.Permision){
      policies:=r.e.GetPolicy()
	  var permissions []model.Permision
	  for i := 0; i < len(policies); i++ {
        permissions = append(permissions, model.Permision{
			Subject: policies[i][0],
			Object: policies[i][1],
			Action: policies[i][2],
		})
	  }
	  if permissions==nil {
			return []model.Permision{}
	  }
	  return permissions;
}
