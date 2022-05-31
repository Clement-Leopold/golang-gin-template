package routers

import (
	"github.com/gin-gonic/gin"
)

// init route group
func InitUserRouters(g *gin.RouterGroup) {
	UserEntityRoute(g)
}
