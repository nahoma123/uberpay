package routing

import (
	"template/internal/adapter/http/rest/server/permission"

	"github.com/gin-gonic/gin"
)

// UserRoutes registers users routes
func PermissionRoutes(grp *gin.RouterGroup, prmHandler permission.PermissionHandler) {
	permGrp:=grp.Group("/permissions")
	permGrp.GET("", prmHandler.Persmisions)
	// permGrp.GET("/:id", prmHandler.Persmision)
	permGrp.POST("",prmHandler.MiddleWareValidatePermission, prmHandler.StorePersmision)
	permGrp.DELETE("", prmHandler.MiddleWareValidatePermission,prmHandler.DeletePersmision)
}
