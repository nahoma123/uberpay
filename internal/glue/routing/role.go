package routing

import (
	"ride_plus/internal/adapter/http/rest/server/auth"

	"github.com/gin-gonic/gin"
)

// RoleRoutes UserRoutes registers users routes
func RoleRoutes(grp *gin.RouterGroup, roleHandler auth.RolesHandler) {
	roleGrp := grp.Group("/roles")
	roleGrp.POST("", roleHandler.RoleMiddleWare, roleHandler.AddRole)
	roleGrp.GET("/:name", roleHandler.GetRoleByName)
	roleGrp.DELETE("/:name", roleHandler.DeleteRole)
	roleGrp.GET("", roleHandler.GetRoles)
}
