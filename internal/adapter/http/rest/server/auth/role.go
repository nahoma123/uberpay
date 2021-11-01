package auth

import (
	errs "errors"
	"fmt"
	"net/http"
	"ride_plus/internal/constant/errors"
	appErr "ride_plus/internal/constant/errors"
	model "ride_plus/internal/constant/model/dbmodel"
	utils "ride_plus/internal/constant/model/init"
	"ride_plus/internal/constant/rest"
	"ride_plus/internal/module/auth"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type RolesHandler interface {
	RoleMiddleWare(c *gin.Context)
	GetRoles(c *gin.Context)
	GetRoleByName(c *gin.Context)
	AddRole(c *gin.Context)
	DeleteRole(c *gin.Context)
}
type rolesHandler struct {
	roleUseCase auth.RoleUseCase
	validate    *validator.Validate
	trans       ut.Translator
}

func NewRoleHandler(useCase auth.RoleUseCase, utils utils.Utils) RolesHandler {
	return &rolesHandler{roleUseCase: useCase, trans: utils.Translator, validate: utils.GoValidator}
}
func (n rolesHandler) RoleMiddleWare(c *gin.Context) {
	roleX := model.Role{}
	err := c.Bind(&roleX)
	if err != nil {
		errValue := errors.ErrorModel{
			ErrorCode:        errors.ErrCodes[errors.ErrInvalidRequest],
			ErrorDescription: errors.Descriptions[errors.ErrInvalidRequest],
			ErrorMessage:     errors.ErrInvalidRequest.Error(),
		}
		rest.ErrorResponseJson(c, errValue, errors.StatusCodes[errors.ErrInvalidRequest])
		return
	}
	c.Set("x-role", roleX)
	c.Next()
}

func (n rolesHandler) GetRoles(c *gin.Context) {
	ctx := c.Request.Context()
	roles, err := n.roleUseCase.Roles(ctx)
	if err != nil {
		ee := errors.ServiceError(err)
		rest.ErrorResponseJson(c, ee, http.StatusBadRequest)
		return
	}
	rest.ErrorResponseJson(c, roles, http.StatusBadRequest)
}

func (n rolesHandler) GetRoleByName(c *gin.Context) {
	rolename := c.Param("name")
	ctx := c.Request.Context()
	r, err := n.roleUseCase.Role(ctx, rolename)
	if err != nil {
		ee := errors.ServiceError(err)
		rest.ErrorResponseJson(c, ee, http.StatusOK)
		return
	}
	c.JSON(200, r)
}

func (n rolesHandler) AddRole(c *gin.Context) {
	rl := c.MustGet("x-role").(model.Role)
	ctx := c.Request.Context()
	fmt.Println("store role ")
	r, err := n.roleUseCase.StoreRole(ctx, rl)
	fmt.Println("error handler ", err)
	if err != nil {
		rest.ErrorResponseJson(c, err, appErr.StatusCodes[errs.New(err.ErrorMessage)])
		return
	}
	rest.ErrorResponseJson(c, r, http.StatusOK)

}

func (n rolesHandler) DeleteRole(c *gin.Context) {
	ctx := c.Request.Context()
	rolename := c.Param("name")
	err := n.roleUseCase.DeleteRole(ctx, rolename)
	if err != nil {
		ee := errors.ServiceError(err)
		rest.ErrorResponseJson(c, ee, http.StatusBadRequest)
		return
	}
	rest.ErrorResponseJson(c, "Deleted Successfully", http.StatusOK)

}
