package routing

import (
	"github.com/gin-gonic/gin"
	"template/internal/adapter/http/rest/server"
)

// UserRoutes registers users routes
func UserRoutes(grp *gin.RouterGroup, usrHandler server.UserHandler) {
	grp.POST("/users", usrHandler.StoreUser)
	grp.PUT("/users/:user-id", usrHandler.UpdateUser)
	grp.GET("/users/:user-id", usrHandler.UserByID)
	grp.GET("/users", usrHandler.Users)
	grp.DELETE("/users/:user-id", usrHandler.DeleteUser)
	grp.GET("/companies/:company-id/users", usrHandler.GetCompanyUsers)
	grp.GET("/companies/:company-id/user", usrHandler.GetCompanyUserByID)
	grp.DELETE("/companies/:company-id/users/:user-id", usrHandler.RemoveUserFromCompany)
	grp.POST("/companies/:company-id/add-users", usrHandler.AddUserToCompany)
}
