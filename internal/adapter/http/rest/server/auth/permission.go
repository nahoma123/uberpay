package auth

import (
	"errors"
	"fmt"
	"net/http"
	"ride_plus/internal/adapter/http/rest/server"
	"ride_plus/internal/module"

	appErr "ride_plus/internal/constant/errors"

	model "ride_plus/internal/constant/model/dbmodel"
	utils "ride_plus/internal/constant/model/init"
	"ride_plus/internal/constant/rest"

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
	prmRequest := c.MustGet("x-request").(model.RolePermission)
	userId := c.Param("user-id")
	fmt.Println("userId", userId)
	prmRequest.UserId = userId
	c.JSON(200, ph.permissionUseCase.AddPermission(prmRequest))
}

func (ph permissionHandler) GetUserPermissions(c *gin.Context) {
	// todo :- fill the object with company as * and permission action and object from permission name
	userId := uuid.FromStringOrNil(c.Param("userId"))
	c.JSON(200, ph.permissionUseCase.GetAllUserPermissions(userId))
}

func (ph permissionHandler) AddUserRole(c *gin.Context) {
	role := c.MustGet("x-request").(model.UserRole)
	rl, srvErr := ph.permissionUseCase.AddRole(role)
	if srvErr != nil {
		rest.ErrorResponseJson(c, srvErr, appErr.StatusCodes[errors.New(srvErr.ErrorMessage)])
		return
	}
	rest.SuccessResponseJson(c, nil, rl, http.StatusCreated)
}

func (ph permissionHandler) GetCompanyRoles(c *gin.Context) {

}
