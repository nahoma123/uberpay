package auth

import (
	"fmt"
	"ride_plus/internal/module"
	"time"

	model "ride_plus/internal/constant/model/dbmodel"
	utils "ride_plus/internal/constant/model/init"

	"github.com/casbin/casbin/v2"
)

type permissionservice struct {
	contextTimeout time.Duration
	enforcer       *casbin.Enforcer
}

func InitializePermission(utils utils.Utils) module.PermissionUseCase {
	return &permissionservice{
		contextTimeout: utils.Timeout,
		enforcer:       utils.Enforcer,
	}
}

func (srv permissionservice) AddBulkPermissions(prms []model.Permission) {
}

func (srv permissionservice) AddPermission(prm model.Permission) error {
	_, err := srv.enforcer.AddPolicy(prm.UserId, prm.CompanyId, prm.Object, prm.Action)
	return err
}

func (srv permissionservice) MigratePermissionsToCasbin() error {
	return nil
}

func (srv permissionservice) GetPermissions(prm model.Permission) []model.Permission {
	fmt.Println(srv.enforcer.GetPermissionsForUserInDomain(prm.UserId, prm.CompanyId))
	return nil
}

func (src permissionservice) IsAuthorized(prm model.Permission) (bool, error) {
	return false, nil
}
