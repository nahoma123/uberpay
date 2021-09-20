package role

import (
	errs "errors"
	"net/http"
	"template/internal/constant/errors"
	"template/internal/constant/model"
	"template/internal/module/role"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
)

type RolesHandler interface {
	MiddleWareValidateRole(c *gin.Context)
	GetRoles(c *gin.Context)
	GetRoleByName(c *gin.Context)
	AddRole(c *gin.Context)
	DeleteRole(c *gin.Context)
}
type rolesHandler struct {
	roleUseCase        role.UseCase
	validate   *validator.Validate
	trans       ut.Translator

}

func NewRoleHandler(useCase role.UseCase,trans ut.Translator,validate   *validator.Validate) RolesHandler {
	return &rolesHandler{roleUseCase: useCase,trans: trans,validate: validate}
}
func (n rolesHandler) MiddleWareValidateRole(c *gin.Context) {
	roleX := model.Role{}
	err := c.Bind(&roleX)
	if err != nil {

		var verr validator.ValidationErrors

		if errs.As(err, &verr) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": verr.Translate(n.trans)})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors":errors.NewErrorResponse(err)})
		return
	}
	valErr := n.validate.Struct(roleX)

	if valErr != nil {
		errs := valErr.(validator.ValidationErrors)
		valErr := errs.Translate(n.trans)
		c.JSON(http.StatusBadRequest, gin.H{"errors":valErr})
		return
	}


	c.Set("x-role", roleX)
	c.Next()
}

func (n rolesHandler) GetRoles(c *gin.Context) {
	roles, err := n.roleUseCase.Roles()

	if err != nil {
		c.JSON(errors.GetStatusCode(err),err)
		return
	}
	c.JSON(200,roles)

}

func (n rolesHandler) GetRoleByName(c *gin.Context) {
	rolename := c.Param("name")

	r, err := n.roleUseCase.Role(rolename)

	if err != nil {
		c.JSON(errors.GetStatusCode(err),err)
		return
	}
	c.JSON(200,r)

}

func (n rolesHandler) AddRole(c *gin.Context) {
	rl := c.MustGet("x-role").(model.Role)

	r, err := n.roleUseCase.StoreRole(rl)

	if err != nil {
				c.JSON(errors.GetStatusCode(err),err)

		return
	}

		c.JSON(200,r)


}

func (n rolesHandler) DeleteRole(c *gin.Context) {

	rolename := c.Param("name")
	err := n.roleUseCase.DeleteRole(rolename)

	if err != nil {
				c.JSON(errors.GetStatusCode(err),err)

		return
	}
	c.JSON(200,"User Delted Successfully")
}
