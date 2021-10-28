package routing

import (
	"ride_plus/internal/adapter/http/rest/middleware"
	"ride_plus/internal/adapter/http/rest/server/auth"
	"ride_plus/internal/constant/permission"

	"github.com/gin-gonic/gin"
)

// RoleRoutes UserRoutes registers users routes
func RoleRoutes(grp *gin.RouterGroup, authMiddleware middleware.AuthMiddleware, roleHandler auth.RolesHandler) {
	roleGrp := grp.Group("/roles")
	roleGrp.POST("", roleHandler.RoleMiddleWare, authMiddleware.Authorizer(permission.CreateSystemRole), roleHandler.AddRole)
	roleGrp.GET("/:name", roleHandler.GetRoleByName)
	roleGrp.DELETE("/:name", roleHandler.DeleteRole)
	roleGrp.GET("", roleHandler.GetRoles)
}
