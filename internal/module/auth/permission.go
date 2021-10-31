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

func (srv permissionservice) AddBulkPermissions(prms []model.RolePermission) {
}

func (srv permissionservice) AddPermission(prm model.RolePermission) error {
	fmt.Println(prm.Role, prm.CompanyId, prm.Object, prm.Action)
	_, err := srv.enforcer.AddPolicy(prm.Role, prm.CompanyId, prm.Object, prm.Action)
	return err
}

func (srv permissionservice) MigratePermissionsToCasbin() error {
	return nil
}

func (srv permissionservice) AddRole(rl model.UserRole) error {
	_, err := srv.enforcer.AddRoleForUserInDomain(rl.UserId, rl.Role, rl.CompanyId)
	permissionsMap, err := srv.enforcer.GetRolesForUser(rl.UserId, rl.CompanyId)
	fmt.Println("ERR-", err)
	fmt.Println("permissionsMap", permissionsMap)
	return nil
}

func (srv permissionservice) GetUserPermissionsInCompany(userId uuid.UUID, prm model.RolePermission) []model.RolePermission {
	srv.enforcer.AddPolicy()
	permissions := []model.RolePermission{}
	if prm.CompanyId == "" {
		prm.CompanyId = "*"
	}

	permissionsMap := srv.enforcer.GetPermissionsForUserInDomain(userId.String(), prm.CompanyId)
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
