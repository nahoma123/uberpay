package routing

import (
	"github.com/gin-gonic/gin"
	"template/internal/adapter/http/rest/server"
)

// CompanyRoutes registers users companies
func CompanyRoutes(grp *gin.RouterGroup, compHandler server.CompanyHandler) {
	grp.POST("/companies", compHandler.Companies)
	grp.GET("/companies", compHandler.Companies)
	grp.GET("/companies/:company-id", compHandler.CompanyByID)
	grp.PUT("/companies/:company-id", compHandler.UpdateCompany)
	grp.DELETE("/companies/:company-id", compHandler.DeleteCompany)
}

