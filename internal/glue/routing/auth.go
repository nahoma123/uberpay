package routing

import (
	"ride_plus/internal/adapter/http/rest/server"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(grp *gin.RouterGroup, loginHandler server.AuthHandler) {
	grp.POST("/login", loginHandler.Login)

}
