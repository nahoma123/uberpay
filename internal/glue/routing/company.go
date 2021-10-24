package routing

import (
	"github.com/gin-gonic/gin"
	"template/internal/adapter/http/rest/server"
)

// CompanyRoutes registers users companies
func CompanyRoutes(grp *gin.RouterGroup, compHandler server.CompanyHandler) {
	grp.POST("/companies", compHandler.CompanyMiddleWare, compHandler.StoreCompany)
	grp.GET("/companies", compHandler.Companies)
	grp.GET("/companies/:company-id", compHandler.CompanyByID)
	grp.GET("/companies/images", compHandler.CompanyImages)
	grp.PUT("/companies/images", compHandler.UpdateCompanyImage)
	grp.POST("/companies/images", compHandler.StoreCompanyImage)
	grp.PUT("/companies/:company-id", compHandler.CompanyMiddleWare, compHandler.UpdateCompany)
	grp.DELETE("/companies/:company-id", compHandler.DeleteCompany)
}
