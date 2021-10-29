package routing

import (
	"ride_plus/internal/adapter/http/rest/middleware"
	"ride_plus/internal/adapter/http/rest/server"
	"ride_plus/internal/constant/permission"

	"github.com/gin-gonic/gin"
)

// CompanyRoutes registers users companies
func CompanyRoutes(grp *gin.RouterGroup, authMiddleware middleware.AuthMiddleware, compHandler server.CompanyHandler) {
	grp.POST("/companies", authMiddleware.Authorizer(permission.CreateCompany), compHandler.StoreCompany)
	grp.GET("/companies", compHandler.Companies)
	grp.GET("/companies/:company-id", compHandler.CompanyByID)
	grp.GET("/companies/images", compHandler.CompanyImages)
	grp.PUT("/companies/images", compHandler.UpdateCompanyImage)
	grp.POST("/companies/images", compHandler.StoreCompanyImage)
	grp.PUT("/companies/:company-id", compHandler.UpdateCompany)
	grp.DELETE("/companies/:company-id", compHandler.DeleteCompany)

}
