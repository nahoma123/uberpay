package routing

import (
	"template/internal/adapter/http/rest/server/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(grp *gin.RouterGroup, loginHandler auth.AuthHandler) {
	grp.POST("/login", loginHandler.Login)

}
