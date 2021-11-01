package routing

import (
	"ride_plus/internal/adapter/http/rest/middleware"
	"ride_plus/internal/adapter/http/rest/server"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(grp *gin.RouterGroup, authMiddleware middleware.AuthMiddleware, loginHandler server.AuthHandler, permissionHandler server.PermissionHandler) {
	grp.POST("/auth/login", loginHandler.Login)
	grp.GET("/auth/users/:user-id/permissions", permissionHandler.GetUserPermissions)
	grp.POST("/auth/users/:user-id/permissions", authMiddleware.BindPermissionRequest(), permissionHandler.AddPermission)

	// add role for system-user
	grp.POST("/auth/users/:user-id/roles", authMiddleware.BindRoleRequest(), permissionHandler.AddUserRole)
}
