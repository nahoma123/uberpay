package policy

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"template/internal/adapter/http/rest/server"
	"template/internal/constant"
	custErr "template/internal/constant/errors"
	"template/internal/constant/model"
	"template/platform/casbin"
)


// policyHandler defines all the things necessary for users handlers
type policyHandler struct {
	casbinAuth casbin.CasbinAuth
}

//PolicyInit initializes a user handler for the domain policy
func PolicyInit(casbinAuth casbin.CasbinAuth) server.PolicyHandler {
	return &policyHandler{
		casbinAuth,
	}
}
func (n policyHandler) PolicyMiddleWare(c *gin.Context) {
	permX := model.Policy{}
	err := c.Bind(&permX)
	if err != nil {
		nrr := custErr.ErrorModel{
			ErrorCode:        custErr.ErrCodes[custErr.ErrorUnableToBindJsonToStruct],
			ErrorMessage:     custErr.Descriptions[custErr.ErrorUnableToBindJsonToStruct],
			ErrorDescription: err.Error(),
		}
		constant.ResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	c.Set("x-policy", permX)
	c.Next()
}

//Policies gets a list of policies
func (uh policyHandler) Policies(c *gin.Context) {
	ctx := c.Request.Context()
	permissions := uh.casbinAuth.Policies(ctx)
	constant.ResponseJson(c, permissions, http.StatusOK)
}
func (uh policyHandler) StorePolicy(c *gin.Context) {
	ctx := c.Request.Context()
	addP := c.MustGet("x-policy").(model.Policy)
	err := uh.casbinAuth.AddPolicy(ctx, addP)

	if err != nil {
		n := custErr.ErrorModel{
			ErrorCode:        custErr.ErrCodes[custErr.ErrUnableToSave],
			ErrorMessage:     custErr.Descriptions[custErr.ErrUnableToSave],
			ErrorDescription: err.Error(),
		}
		constant.ResponseJson(c, n, http.StatusBadRequest)
		return
	}
	constant.ResponseJson(c, addP, http.StatusOK)
}

//RemovePolicy removes policy
func (uh policyHandler) RemovePolicy(c *gin.Context) {
	ctx := c.Request.Context()
	addP := c.MustGet("x-policy").(model.Policy)
	err := uh.casbinAuth.RemovePolicy(ctx, addP)
	if err != nil {
		n := custErr.ErrorModel{
			ErrorCode:        custErr.ErrCodes[custErr.ErrUnableToDelete],
			ErrorMessage:     custErr.Descriptions[custErr.ErrUnableToDelete],
			ErrorDescription: err.Error(),
		}
		constant.ResponseJson(c, n, http.StatusBadRequest)
		return
	}
	constant.ResponseJson(c, addP, http.StatusOK)
}

//UpdatePolicy updates policies new by old
func (uh policyHandler) UpdatePolicy(c *gin.Context) {
	ctx := c.Request.Context()
	permU := model.PolicyUpdate{}
	err := c.Bind(&permU)
	if err != nil {
		constant.ResponseJson(c, custErr.NewErrorResponse(err), http.StatusBadRequest)
		return
	}
	err = uh.casbinAuth.UpdatePolicy(ctx, permU)
	if err != nil {
		nrr := custErr.ErrorModel{
			ErrorCode:        custErr.ErrCodes[custErr.ErrUnableToSave],
			ErrorMessage:     "Update failed",
			ErrorDescription: err.Error(),
		}
		constant.ResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	constant.ResponseJson(c, "Update successful", http.StatusOK)
}
