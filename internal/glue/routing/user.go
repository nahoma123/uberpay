package routing

import (
	"ride_plus/internal/adapter/http/rest/middleware"
	"ride_plus/internal/adapter/http/rest/server"

	"github.com/gin-gonic/gin"
)

// UserRoutes registers users routes
func UserRoutes(grp *gin.RouterGroup, authMiddleware middleware.AuthMiddleware, usrHandler server.UserHandler) {
	grp.POST("/users", usrHandler.UserMiddleWare, usrHandler.RegisterUser)
	grp.PUT("/users/:user-id", usrHandler.UserMiddleWare, usrHandler.UpdateUser)
	grp.GET("/users/:user-id", usrHandler.UserByID)
	grp.GET("/users", usrHandler.Users)
	grp.DELETE("/users/:user-id", usrHandler.DeleteUser)
	grp.GET("/companies/:company-id/users", usrHandler.GetCompanyUsers)
	grp.GET("/companies/:company-id/users/:user-id", usrHandler.GetCompanyUserByID)
	grp.DELETE("/companies/:company-id/users/:user-id", usrHandler.RemoveUserFromCompany)
	grp.POST("/companies/:company-id/add-users", usrHandler.AddUserToCompany)
}
