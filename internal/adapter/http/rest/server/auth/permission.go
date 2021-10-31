package auth

import (
	"ride_plus/internal/adapter/http/rest/server"
	"ride_plus/internal/module"

	model "ride_plus/internal/constant/model/dbmodel"
	utils "ride_plus/internal/constant/model/init"
	"ride_plus/internal/constant/permission"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)

type permissionHandler struct {
	permissionUseCase module.PermissionUseCase
	validate          *validator.Validate
	trans             ut.Translator
}

func NewPermissionHandler(prmUseCase module.PermissionUseCase, utils utils.Utils) server.PermissionHandler {
	return &permissionHandler{permissionUseCase: prmUseCase, trans: utils.Translator, validate: utils.GoValidator}
}

func (ph permissionHandler) AddPermission(c *gin.Context) {
	prm := model.RolePermission{
		CompanyId: "PERSONAL",
		Object:    permission.PermissionObjects[permission.CreateCompany],
		Action:    permission.PermissionActions[permission.CreateCompany],
	}
	c.JSON(200, ph.permissionUseCase.AddPermission(prm))
}

func (ph permissionHandler) GetUserPermissions(c *gin.Context) {
	// todo :- fill the object with company as * and permission action and object from permission name
	userId := uuid.FromStringOrNil(c.Param("userId"))
	c.JSON(200, ph.permissionUseCase.GetAllUserPermissions(userId))
}
