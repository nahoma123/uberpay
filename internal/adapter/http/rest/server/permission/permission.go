package permission

import (
	"net/http"
	errModel "template/internal/constant/errors"
	"template/internal/constant/model"
	"template/platform/casbin"

	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
	ut "github.com/go-playground/universal-translator"
)

// PermissionHandler contans a function of handlers for the domian file
type PermissionHandler interface {
	MiddleWareValidatePermission(c *gin.Context)
	StorePersmision(c *gin.Context)
	// Persmision(c *gin.Context)
	DeletePersmision(c *gin.Context)
	Persmisions(c *gin.Context)
}

// userHandler defines all the things neccessary for users handlers
type permissionHandler struct {
	casbinAuth casbin.CasbinAuth
	validate   *validator.Validate
	trans       ut.Translator
}

//PermissionInit initializes a user handler for the domin permission
func PermissionInit(
	casbinAuth casbin.CasbinAuth,
	validate   *validator.Validate,
	 trans ut.Translator,

	 ) PermissionHandler {
	return &permissionHandler{
		casbinAuth,
		validate,
		trans,

	}
}
func (n permissionHandler) MiddleWareValidatePermission(c *gin.Context) {
	permX := model.Permision{}
	err := c.Bind(&permX)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errModel.NewErrorResponse(err)})
		return
	}
	valErr := n.validate.Struct(permX)

	if valErr != nil {
		errs := valErr.(validator.ValidationErrors)
		valErr := errs.Translate(n.trans)
		c.JSON(http.StatusBadRequest, gin.H{"errors":valErr})
		return
	}

	c.Set("x-permission", permX)
	c.Next()
}
// Persmisions gets a list of permissions
func (uh permissionHandler) Persmisions(c *gin.Context) {
	permissions := uh.casbinAuth.Policies()
	c.JSON(http.StatusOK, gin.H{"permissions": permissions})
}
func (n permissionHandler) StorePersmision(c *gin.Context) {
	addP := c.MustGet("x-permission").(model.Permision)

	err := n.casbinAuth.AddPolicy(addP)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    c.JSON(200,addP)
}

// DeletePermission deletes user by id
func (uh permissionHandler) DeletePersmision(c *gin.Context) {
	addP := c.MustGet("x-permission").(model.Permision)

	err := uh.casbinAuth.RemovePolicy(addP)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    c.JSON(200,addP)
}
