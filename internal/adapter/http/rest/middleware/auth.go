package middleware

import (
	"fmt"
	"log"
	"net/http"
	"ride_plus/internal/constant"
	"ride_plus/internal/module"
	"strings"

	utils "ride_plus/internal/constant/model/init"

	appErr "ride_plus/internal/constant/errors"

	"github.com/casbin/casbin/v2"

	"github.com/gin-gonic/gin"
)

var actions = map[string]string{
	"GET":    "read",
	"POST":   "create",
	"PUT":    "update",
	"DELETE": "delete",
	"PATCH":  "update",
}

type AuthMiddleware interface {
	Authorizer(e *casbin.Enforcer) gin.HandlerFunc
	ExtractToken(r *http.Request) string
}

type authMiddleWare struct {
	authUseCase module.LoginUseCase
	utils       utils.Utils
}

func NewAuthMiddleware(authUseCase module.LoginUseCase, utils utils.Utils) AuthMiddleware {
	return &authMiddleWare{
		authUseCase: authUseCase,
		utils:       utils,
	}
}

//Authorizer is a middleware for authorization
func (n *authMiddleWare) Authorizer(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := "anonymous"
		token := n.ExtractToken(c.Request)
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

func (n *authMiddleWare) ExtractToken(r *http.Request) string {
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
