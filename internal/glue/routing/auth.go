package routing

import (
	"github.com/gin-gonic/gin"
	"template/internal/adapter/http/rest/server"
)

func AuthRoutes(grp *gin.RouterGroup, loginHandler server.AuthHandler) {
	grp.POST("/login", loginHandler.Login)

}
