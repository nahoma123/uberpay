package routing

import (
	"template/internal/adapter/http/rest/server/company"

	"github.com/gin-gonic/gin"
)

// CompanyRoutes registers users companies
func CompanyRoutes(grp *gin.RouterGroup, compHandler company.CompanyHandler) {
	grp.POST("/companies", compHandler.CreateCompany)
	grp.GET("/companies/:id", compHandler.GetCompanyById)
	grp.GET("/companies", compHandler.Companies)
	grp.DELETE("/companies/:id", compHandler.DeleteCompany)
}
