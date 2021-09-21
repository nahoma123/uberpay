package permission

import (
	"errors"
	"net/http"
	// "strconv"
	errModel "template/internal/constant/errors"
	"template/internal/constant/model"
	"template/platform/casbin"

	"github.com/gin-gonic/gin"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"
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
	trans       ut.Translator
}

//PermissionInit initializes a user handler for the domin permission
func PermissionInit(
	casbinAuth casbin.CasbinAuth,
	 trans ut.Translator,

	 ) PermissionHandler {
	return &permissionHandler{
		casbinAuth,
		trans,

	}
}
func (n permissionHandler) MiddleWareValidatePermission(c *gin.Context) {
	permX := model.Permision{}
	err := c.Bind(&permX)
	if err != nil {

		var verr validator.ValidationErrors

		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": verr.Translate(n.trans)})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": errModel.NewErrorResponse(errModel.ErrUnknown)})
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
		c.JSON(errModel.GetStatusCode(err),err)

		return
	}
    c.JSON(200,addP)
}

// Persmision gets a permission by id
// func (uh permissionHandler) Persmision(c *gin.Context) {

// 	ID := c.Param("id")

// 	id, err := strconv.Atoi(ID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"errors": errModel.NewErrorResponse(err)})
// 		return
// 	}

// 	perm, err := uh.permisionUsecase.Persmision(uint(id))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"errors": errModel.NewErrorResponse(err)})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"permission": perm})

// }

// DeletePermission deletes user by id
func (uh permissionHandler) DeletePersmision(c *gin.Context) {
	addP := c.MustGet("x-permission").(model.Permision)

	err := uh.casbinAuth.RemovePolicy(addP)

	if err != nil {
		c.JSON(errModel.GetStatusCode(err),err)

		return
	}
    c.JSON(200,addP)
}
