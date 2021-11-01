package auth

import (
	"fmt"
	"ride_plus/internal/module"
	"time"

	"ride_plus/internal/constant/errors"
	appErr "ride_plus/internal/constant/errors"
	model "ride_plus/internal/constant/model/dbmodel"
	utils "ride_plus/internal/constant/model/init"
	"ride_plus/internal/constant/permission"

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

func (srv permissionservice) AddBulkPermissions(prms []model.RolePermission) {
}

func (srv permissionservice) AddPermission(prm model.RolePermission) error {
	if permission.PermissionObjects[prm.Name] != "" {
		prm := model.RolePermission{
			Role:   prm.Role,
			Tenant: prm.Tenant,
			Object: permission.PermissionObjects[prm.Name],
			Action: permission.PermissionActions[prm.Name],
		}
		fmt.Println(prm)
	}

	_, err := srv.enforcer.AddPolicy(prm.Role, prm.Tenant, prm.Object, prm.Action)
	return err
}

func (srv permissionservice) MigratePermissionsToCasbin() error {
	return nil
}

func (srv permissionservice) AddRole(rl model.UserRole) (*model.UserRole, *errors.ErrorModel) {
	// _, err := srv.enforcer.AddRoleForUserInDomain(rl.UserId, rl.Role, rl.Tenant)
	// permissionsMap, err := srv.enforcer.GetRolesForUser(rl.UserId, rl.Tenant)
	// fmt.Println("ERR-", err)
	// fmt.Println("permissionsMap", permissionsMap)
	// return nil

	_, err := srv.enforcer.AddRoleForUserInDomain(rl.UserId, rl.Role, rl.Tenant)
	if err != nil {
		return nil, appErr.ServiceError(appErr.ErrorUnableToBindJsonToStruct)
	}
	return &rl, nil
}

func (srv permissionservice) GetUserPermissionsInCompany(userId uuid.UUID, prm model.RolePermission) []model.RolePermission {
	srv.enforcer.AddPolicy()
	permissions := []model.RolePermission{}
	if prm.Tenant == "" {
		prm.Tenant = "*"
	}

	permissionsMap := srv.enforcer.GetPermissionsForUserInDomain(userId.String(), prm.Tenant)
	// for i, _ := range permi                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                ssionsMap {
	// 	permissions[i] = model.RolePermission{}
	// }
	fmt.Println("permissionsMap", permissionsMap)
	return permissions
}

func (srv permissionservice) GetAllUserPermissions(userId uuid.UUID) []model.RolePermission {
	// srv.enforcer.AddPolicy()
	permissions := []model.RolePermission{}

	permissionsMap := srv.enforcer.GetPermissionsForUser(userId.String())
	for i, prm := range permissionsMap {
		permission := model.RolePermission{}
		permission.ID = prm[0]
		permission.UserId = prm[0]
		permission.Name = prm[2]
		permission.Action = fmt.Sprintf("%s, %s", prm[2], prm[3])
		permission.Object = prm[2]
		permissions = append(permissions, permission)
		fmt.Println(i, prm)
	}

	fmt.Println("permissionsMap", permissions)
	return permissions
}

func (src permissionservice) IsAuthorized(prm model.RolePermission) (bool, error) {
	return false, nil
}
