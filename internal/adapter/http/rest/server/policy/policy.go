package policy

import (
	"net/http"
	"os"
	"ride_plus/internal/adapter/http/rest/server"
	"ride_plus/internal/constant"
	custErr "ride_plus/internal/constant/errors"
	"ride_plus/internal/constant/model"
	"ride_plus/platform/casbin"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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
func (ph policyHandler) PolicyMiddleWare(c *gin.Context) {
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
func (ph policyHandler) Policies(c *gin.Context) {
	ctx := c.Request.Context()
	permissions := ph.casbinAuth.Policies(ctx)
	constant.ResponseJson(c, permissions, http.StatusOK)
}
func (ph policyHandler) StorePolicy(c *gin.Context) {
	ctx := c.Request.Context()
	addP := c.MustGet("x-policy").(model.Policy)
	err := ph.casbinAuth.AddPolicy(ctx, addP)
	if err != nil {
		if strings.Contains(err.Error(), os.Getenv("ErrSecretKey")) {
			e := strings.Replace(err.Error(), os.Getenv("ErrSecretKey"), "", 1)
			errValue := custErr.ErrorModel{
				ErrorCode:        custErr.ErrCodes[custErr.ErrInvalidField],
				ErrorDescription: custErr.Descriptions[custErr.ErrInvalidField],
				ErrorMessage:     e,
			}
			constant.ResponseJson(c, errValue, http.StatusBadRequest)
			return
		}
		constant.ResponseJson(c, custErr.NewErrorResponse(err), custErr.ErrCodes[err])
		return
	}
	constant.ResponseJson(c, addP, http.StatusOK)
}

//RemovePolicy removes policy
func (ph policyHandler) RemovePolicy(c *gin.Context) {
	ctx := c.Request.Context()
	addP := c.MustGet("x-policy").(model.Policy)
	err := ph.casbinAuth.RemovePolicy(ctx, addP)
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
func (ph policyHandler) UpdatePolicy(c *gin.Context) {
	ctx := c.Request.Context()
	permU := model.PolicyUpdate{}
	err := c.Bind(&permU)
	if err != nil {
		constant.ResponseJson(c, custErr.NewErrorResponse(err), http.StatusBadRequest)
		return
	}
	err = ph.casbinAuth.UpdatePolicy(ctx, permU)
	if err != nil {
		if strings.Contains(err.Error(), os.Getenv("ErrSecretKey")) {
			e := strings.Replace(err.Error(), os.Getenv("ErrSecretKey"), "", 1)
			errValue := custErr.ErrorModel{
				ErrorCode:        custErr.ErrCodes[custErr.ErrInvalidField],
				ErrorDescription: custErr.Descriptions[custErr.ErrInvalidField],
				ErrorMessage:     e,
			}
			constant.ResponseJson(c, errValue, http.StatusBadRequest)
			return
		}
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

//GetCompanyPolicyByID gets all company policy identified by id
func (ph policyHandler) GetCompanyPolicyByID(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := uuid.FromString(c.Param("company-id"))
	if err != nil {
		constant.ResponseJson(c, custErr.ConvertionError(), http.StatusBadRequest)
		return
	}
	policies := ph.casbinAuth.GetCompanyPolicyByID(ctx, id.String())
	constant.ResponseJson(c, policies, http.StatusOK)

}

//GetAllCompaniesPolicy gets all company policy
func (ph policyHandler) GetAllCompaniesPolicy(c *gin.Context) {
	ctx := c.Request.Context()
	policies := ph.casbinAuth.GetAllCompaniesPolicy(ctx)
	constant.ResponseJson(c, policies, http.StatusOK)
}
