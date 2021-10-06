package routing

import (
	"template/internal/adapter/http/rest/server/role"

	"github.com/gin-gonic/gin"
)

// RoleRoutes UserRoutes registers users routes
func RoleRoutes(grp *gin.RouterGroup, roleHandler role.RolesHandler) {
	roleGrp := grp.Group("/roles")
	roleGrp.POST("", roleHandler.RoleMiddleWare, roleHandler.AddRole)
	roleGrp.GET("/:name", roleHandler.GetRoleByName)
	roleGrp.DELETE("/:name", roleHandler.DeleteRole)
	roleGrp.GET("", roleHandler.GetRoles)
}
