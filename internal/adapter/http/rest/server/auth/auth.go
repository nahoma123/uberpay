package auth

import (
	"net/http"
	"ride_plus/internal/adapter/http/rest/server"
	appErr "ride_plus/internal/constant/errors"
	model "ride_plus/internal/constant/model/dbmodel"
	utils "ride_plus/internal/constant/model/init"
	"ride_plus/internal/constant/rest"
	"ride_plus/internal/module"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authUseCase module.LoginUseCase
	utils       utils.Utils
}

func NewAuthHandler(authUseCase module.LoginUseCase, utils utils.Utils) server.AuthHandler {
	return &authHandler{
		authUseCase: authUseCase,
		utils:       utils,
	}
}

func (n authHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	authData := &model.User{}
	err := c.Bind(authData)
	if err != nil {
		rest.ErrorResponseJson(c, appErr.ServiceError(appErr.ErrorUnableToBindJsonToStruct), http.StatusBadRequest)
		return
	}
	loginResponse, err := n.authUseCase.Login(ctx, authData.Phone, authData.Password)
	if err != nil {
		rest.ErrorResponseJson(c, appErr.ServiceError(err), http.StatusUnauthorized)
		return
	}
	rest.ErrorResponseJson(c, loginResponse, http.StatusOK)
}
