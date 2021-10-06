package routing

import (
	"github.com/gin-gonic/gin"
	"template/internal/adapter/http/rest/server"
)

// PolicyRoutes UserRoutes registers users routes
func PolicyRoutes(permGrp *gin.RouterGroup, prmHandler server.PolicyHandler) {
	permGrp.GET("/policies", prmHandler.Policies)
	permGrp.POST("/policies", prmHandler.PolicyMiddleWare, prmHandler.StorePolicy)
	permGrp.PUT("/policies", prmHandler.UpdatePolicy)
	permGrp.DELETE("/policies", prmHandler.PolicyMiddleWare, prmHandler.RemovePolicy)
}
