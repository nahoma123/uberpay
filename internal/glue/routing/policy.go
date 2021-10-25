package routing

import (
	"ride_plus/internal/adapter/http/rest/server"

	"github.com/gin-gonic/gin"
)

// PolicyRoutes UserRoutes registers users routes
func PolicyRoutes(permGrp *gin.RouterGroup, prmHandler server.PolicyHandler) {
	permGrp.GET("/policies", prmHandler.Policies)
	permGrp.GET("/companies/policies", prmHandler.GetAllCompaniesPolicy)
	permGrp.GET("/companies/:company-id/policies", prmHandler.GetCompanyPolicyByID)
	permGrp.POST("/policies", prmHandler.PolicyMiddleWare, prmHandler.StorePolicy)
	permGrp.PUT("/policies", prmHandler.UpdatePolicy)
	permGrp.DELETE("/policies", prmHandler.PolicyMiddleWare, prmHandler.RemovePolicy)
}
