package middleware

import (
	"log"
	"net/http"
	"ride_plus/internal/constant"
	"ride_plus/internal/module"
	"strings"

	utils "ride_plus/internal/constant/model/init"

	appErr "ride_plus/internal/constant/errors"

	permission "ride_plus/internal/constant/permission"

	"github.com/gin-gonic/gin"
)

type Status struct {
	Status string `json:"status"`
}

var actions = map[string]string{
	"GET":    "read",
	"POST":   "create",
	"PUT":    "update",
	"DELETE": "delete",
	"PATCH":  "update",
}

type AuthMiddleware interface {
	Authorizer(permission string) gin.HandlerFunc
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
func (n *authMiddleWare) Authorizer(prm string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := "anonymous"
		token := n.ExtractToken(c.Request)
		status := Status{}
		err := c.Bind(&status)
		if err != nil {
			err := appErr.NewErrorResponse(appErr.ErrorUnableToBindJsonToStruct)
			constant.ResponseJson(c, err, appErr.StatusCodes[appErr.ErrorUnableToBindJsonToStruct])
			c.AbortWithStatus(appErr.StatusCodes[appErr.ErrorUnableToBindJsonToStruct])
			return
		}

		claims, err := n.authUseCase.GetClaims(token)
		if err != nil {
			err := appErr.NewErrorResponse(appErr.ErrInvalidAccessToken)
			constant.ResponseJson(c, err, appErr.StatusCodes[appErr.ErrInvalidAccessToken])
			c.AbortWithStatus(appErr.StatusCodes[appErr.ErrInvalidAccessToken])
			return
		}

		if claims != nil {
			role = claims.Role
			c.Set("x-userid", claims.Subject)
			c.Set("x-userrole", role)
			if claims.CompanyID != "" {
				c.Set("x-companyid", claims.CompanyID)
			}
		} else {
			err := appErr.NewErrorResponse(appErr.ErrAuthorizationTokenNotProvided)
			constant.ResponseJson(c, err, appErr.StatusCodes[appErr.ErrAuthorizationTokenNotProvided])
			c.AbortWithStatus(appErr.StatusCodes[appErr.ErrAuthorizationTokenNotProvided])
			return
		}

		err = n.utils.Enforcer.LoadPolicy()
		if err != nil {
			log.Fatal("error ", err)
		}
		var c_id string

		if claims.CompanyID == "" {
			c_id = "*"
		} else {
			c_id = strings.TrimSpace(claims.CompanyID)
		}

		isAuthorized := false
		if permission.PermissionActions[prm] == permission.Create || permission.PermissionActions[prm] == permission.Update {
			if status.Status != "" {
				// if status being changed or provided then we need to ensure that user has the authority to publish.
				isAuthorized, err = n.utils.Enforcer.Enforce(claims.Subject, c_id, permission.PermissionObjects[prm], permission.PermissionActions[prm])
				if err != nil {
					err := appErr.NewErrorResponse(appErr.ErrPermissionPermissionNotFound)
					constant.ResponseJson(c, err, appErr.StatusCodes[appErr.ErrPermissionPermissionNotFound])
					c.AbortWithStatus(appErr.StatusCodes[appErr.ErrPermissionPermissionNotFound])
					return
				}
				if !isAuthorized {
					err := appErr.NewErrorResponse(appErr.ErrUnauthorizedClient)
					constant.ResponseJson(c, err, appErr.StatusCodes[appErr.ErrUnauthorizedClient])
					c.AbortWithStatus(appErr.StatusCodes[appErr.ErrUnauthorizedClient])
					return
				}
			}
		}

		isAuthorized, err = n.utils.Enforcer.Enforce(claims.Subject, c_id, permission.DraftPermissions[prm], permission.PermissionActions[prm])
		if err != nil {
			err := appErr.NewErrorResponse(appErr.ErrPermissionPermissionNotFound)
			constant.ResponseJson(c, err, appErr.StatusCodes[appErr.ErrPermissionPermissionNotFound])
			c.AbortWithStatus(appErr.StatusCodes[appErr.ErrPermissionPermissionNotFound])
			return
		}
		if !isAuthorized {
			err := appErr.NewErrorResponse(appErr.ErrUnauthorizedClient)
			constant.ResponseJson(c, err, appErr.StatusCodes[appErr.ErrUnauthorizedClient])
			c.AbortWithStatus(appErr.StatusCodes[appErr.ErrUnauthorizedClient])
			return
		}
		c.Next()

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
