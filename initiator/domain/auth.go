package domain

import (
	"ride_plus/internal/adapter/http/rest/middleware"
	authHandler "ride_plus/internal/adapter/http/rest/server/auth"
	authPersistence "ride_plus/internal/adapter/storage/persistence/auth"
	"ride_plus/internal/adapter/storage/persistence/user"
	utils "ride_plus/internal/constant/model/init"
	routing2 "ride_plus/internal/glue/routing"
	authUsecase "ride_plus/internal/module/auth"
	roleUsecase "ride_plus/internal/module/auth"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(utils utils.Utils) middleware.AuthMiddleware {
	usrPersistence := user.UserInit(utils)
	jwtManager := authUsecase.NewJWTManager("secret")
	loginUseCase := authUsecase.Initialize(usrPersistence, *jwtManager, utils)

	authMiddleWare := middleware.NewAuthMiddleware(loginUseCase, utils)
	return authMiddleWare
}

func AuthInit(utils utils.Utils, router *gin.RouterGroup) {
	rolePersistent := authPersistence.RoleInit(utils)
	roleUsecase := roleUsecase.RoleInitialize(rolePersistent, utils)
	roleHandler := authHandler.NewRoleHandler(roleUsecase, utils)

	usrPersistence := user.UserInit(utils)

	jwtManager := authUsecase.NewJWTManager("secret")
	authUsecases := authUsecase.Initialize(usrPersistence, *jwtManager, utils)
	authHandlers := authHandler.NewAuthHandler(authUsecases, utils)

	permissionUsecases := authUsecase.InitializePermission(utils)
	permissionHandler := authHandler.NewPermissionHandler(permissionUsecases, utils)

	routing2.RoleRoutes(router, AuthMiddleware(utils), roleHandler)
	// routing2.PolicyRoutes(router, permHandler)
	routing2.AuthRoutes(router, AuthMiddleware(utils), authHandlers, permissionHandler)
}
