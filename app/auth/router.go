package auth

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routeMap := router.Group("/auth")
	routeMap.GET("/register", register)
	routeMap.GET("/login", login)
}
