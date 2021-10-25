package initiator

import (
	"log"
	authHandler "template/internal/adapter/http/rest/server/auth"
	permHandler "template/internal/adapter/http/rest/server/policy"
	rlHandler "template/internal/adapter/http/rest/server/role"
	"template/internal/adapter/storage/persistence/role"
	"template/internal/adapter/storage/persistence/user"
	"template/internal/constant/model"
	utils "template/internal/constant/model/init"
	routing2 "template/internal/glue/routing"
	authUsecase "template/internal/module/auth"
	roleUsecase "template/internal/module/role"
	casAuth "template/platform/casbin"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	"github.com/gin-gonic/gin"
)

func AuthInit(utils utils.Utils, router *gin.RouterGroup) {

	a, err := gormadapter.NewAdapterByDBWithCustomTable(utils.Conn, &model.CasbinRule{})
	if err != nil {
		log.Fatal("error ", err)
	}

	e, err := casbin.NewEnforcer("../../rbac_model.conf", a)
	if err != nil {
		log.Fatal("error ", err)
	}
	rolePersistent := role.RoleInit(utils)

	roleUsecase := roleUsecase.RoleInitialize(rolePersistent, utils)
	roleHandler := rlHandler.NewRoleHandler(roleUsecase, utils)

	casbinAuth := casAuth.NewEnforcer(e, utils)
	permHandler := permHandler.PolicyInit(casbinAuth)

	usrPersistence := user.UserInit(utils)

	jwtManager := authUsecase.NewJWTManager("secret")
	authUsecases := authUsecase.Initialize(usrPersistence, *jwtManager, utils)
	authHandlers := authHandler.NewAuthHandler(authUsecases, casbinAuth)

	router.Use(authHandlers.Authorizer(e))
	routing2.RoleRoutes(router, roleHandler)
	routing2.PolicyRoutes(router, permHandler)
	routing2.AuthRoutes(router, authHandlers)

}
