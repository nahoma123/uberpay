package routing

import (
	"template/internal/adapter/http/rest/server/user"

	"github.com/gin-gonic/gin"
)

// UserRoutes registers users routes
func UserRoutes(grp *gin.RouterGroup, usrHandler user.UserHandler) {
	grp.POST("/users", usrHandler.CreateSystemUser)
	grp.GET("/users", usrHandler.GetUsers)
	grp.GET("/users/:id", usrHandler.GetUserById)
	grp.DELETE("/users/:id", usrHandler.DeleteUser)
	grp.POST("/companies/:comp-id/users", usrHandler.CreateUser)
}
