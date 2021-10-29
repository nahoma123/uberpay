package auth

import (
	"fmt"
	"ride_plus/internal/module"
	"time"

	model "ride_plus/internal/constant/model/dbmodel"
	utils "ride_plus/internal/constant/model/init"

	"github.com/casbin/casbin/v2"
	uuid "github.com/satori/go.uuid"
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

func (srv permissionservice) GetUserPermissions(prm model.Permission) []model.Permission {
	fmt.Println(srv.enforcer.GetPermissionsForUserInDomain(prm.UserId, prm.CompanyId))
	return nil
}

func (src permissionservice) IsAuthorized(userId uuid.UUID, companyId uuid.UUID, obj string, action string) (bool, error) {
	return false, nil
}
