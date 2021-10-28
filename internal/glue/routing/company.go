package routing

import (
	"ride_plus/internal/adapter/http/rest/server"

	"github.com/gin-gonic/gin"
)

// CompanyRoutes registers users companies
func CompanyRoutes(grp *gin.RouterGroup, compHandler server.CompanyHandler) {
	grp.POST("/companies", compHandler.StoreCompany)
	grp.GET("/companies", compHandler.Companies)
	grp.GET("/companies/:company-id", compHandler.CompanyByID)
	grp.GET("/companies/images", compHandler.CompanyImages)
	grp.PUT("/companies/images", compHandler.UpdateCompanyImage)
	grp.POST("/companies/images", compHandler.StoreCompanyImage)
	grp.PUT("/companies/:company-id", compHandler.UpdateCompany)
	grp.DELETE("/companies/:company-id", compHandler.DeleteCompany)

}
