package auth

import (
	"github.com/gin-gonic/gin"
	"go-api-app/app/auth/dto"
	Validator "go-api-app/app/common/middleware"
)

func Routes(router *gin.Engine) {
	routeMap := router.Group("/auth")
	routeMap.POST("/register", Validator.CheckBody[dto.UserCreate], register)
	routeMap.GET("/login", login)
}
