package initiator

import (
	authHandler "ride_plus/internal/adapter/http/rest/server/auth"
	permHandler "ride_plus/internal/adapter/http/rest/server/policy"
	rlHandler "ride_plus/internal/adapter/http/rest/server/role"
	"ride_plus/internal/adapter/storage/persistence/role"
	"ride_plus/internal/adapter/storage/persistence/user"
	utils "ride_plus/internal/constant/model/init"
	routing2 "ride_plus/internal/glue/routing"
	authUsecase "ride_plus/internal/module/auth"
	roleUsecase "ride_plus/internal/module/role"
	authorization "ride_plus/platform/authorization"

	"github.com/gin-gonic/gin"
)

const (
	authModel = "config/rbac_model.conf"
)

func AuthInit(utils utils.Utils, router *gin.RouterGroup) {
	rolePersistent := role.RoleInit(utils)
	roleUsecase := roleUsecase.RoleInitialize(rolePersistent, utils)
	roleHandler := rlHandler.NewRoleHandler(roleUsecase, utils)

	casbinAuth := authorization.NewEnforcer(utils.Conn)
	permHandler := permHandler.PolicyInit(casbinAuth)

	usrPersistence := user.UserInit(utils)

	jwtManager := authUsecase.NewJWTManager("secret")
	authUsecases := authUsecase.Initialize(usrPersistence, *jwtManager, utils)
	authHandlers := authHandler.NewAuthHandler(authUsecases, casbinAuth)

	router.Use(authHandlers.Authorizer(utils.Enforcer))
	routing2.RoleRoutes(router, roleHandler)
	routing2.PolicyRoutes(router, permHandler)
	routing2.AuthRoutes(router, authHandlers)

}
