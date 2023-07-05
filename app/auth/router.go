package auth

import (
	"github.com/gin-gonic/gin"
	"go-api-app/app/auth/dto"
	"go-api-app/app/common/middleware"
)

func Routes(router *gin.Engine) {
	routeMap := router.Group("/auth")
	routeMap.POST("/register", middleware.ValidateBody[authDto.UserCreate], register)
	routeMap.GET("/login", middleware.ValidateQuery[authDto.Filter], login)
}
