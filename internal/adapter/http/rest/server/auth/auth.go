package auth

import (
	"fmt"
	"log"
	"net/http"
	"ride_plus/internal/adapter/http/rest/server"
	"ride_plus/internal/constant"
	appErr "ride_plus/internal/constant/errors"
	model "ride_plus/internal/constant/model/dbmodels"
	"ride_plus/internal/module"
	custCas "ride_plus/platform/casbin"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authUseCase module.LoginUseCase
	casbinAuth  custCas.CasbinAuth
}

func NewAuthHandler(authUseCase module.LoginUseCase, perm custCas.CasbinAuth) server.AuthHandler {
	return &authHandler{
		authUseCase: authUseCase,
		casbinAuth:  perm,
	}
}

var actions = map[string]string{
	"GET":    "read",
	"POST":   "create",
	"PUT":    "update",
	"DELETE": "delete",
	"PATCH":  "update",
}

//Authorizer is a middleware for authorization
func (n *authHandler) Authorizer(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := "anonymous"
		token := ExtractToken(c.Request)
		claims, _ := n.authUseCase.GetClaims(token)
		if claims != nil {
			role = claims.Role
			c.Set("x-userid", claims.Subject)
			c.Set("x-userrole", role)
			if claims.CompanyID != "" {
				c.Set("x-companyid", claims.CompanyID)
			}
		}
		err := e.LoadPolicy()
		if err != nil {
			log.Fatal("error ", err)
		}
		var c_id string
		if claims.CompanyID == "" {
			c_id = "*"
		} else {
			c_id = strings.TrimSpace(claims.CompanyID)
		}

		res, err := e.Enforce(role, c.Request.URL.Path, actions[c.Request.Method], c_id)
		fmt.Println("enforce error ", err, "res ", res)
		if err != nil {
			err := appErr.NewErrorResponse(appErr.ErrPermissionPermissionNotFound)
			constant.ResponseJson(c, err, appErr.StatusCodes[appErr.ErrPermissionPermissionNotFound])
			c.AbortWithStatus(appErr.StatusCodes[appErr.ErrPermissionPermissionNotFound])
			return
		}
		if res {
			c.Next()
		} else {
			err := appErr.NewErrorResponse(appErr.ErrUnauthorizedClient)
			constant.ResponseJson(c, err, appErr.StatusCodes[appErr.ErrUnauthorizedClient])
			c.AbortWithStatus(appErr.StatusCodes[appErr.ErrUnauthorizedClient])
			return
		}

	}
}
func (n authHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	authData := &model.User{}
	err := c.Bind(authData)
	if err != nil {
		constant.ResponseJson(c, appErr.NewErrorResponse(appErr.ErrorUnableToBindJsonToStruct), http.StatusBadRequest)
		return
	}
	loginResponse, err := n.authUseCase.Login(ctx, authData.Phone, authData.Password)
	if err != nil {
		constant.ResponseJson(c, appErr.NewErrorResponse(err), http.StatusUnauthorized)
		return
	}
	constant.ResponseJson(c, loginResponse, http.StatusOK)
}
func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("auth")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}