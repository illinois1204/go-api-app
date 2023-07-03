package auth

import (
	"github.com/gin-gonic/gin"
	"go-api-app/app/auth/dto"
	"net/http"
)

func register(ctx *gin.Context) {
	body, _ := ctx.MustGet("body").(dto.UserCreate)

	ctx.JSON(http.StatusOK, body)
	return

	ctx.JSON(http.StatusOK, gin.H{
		"data": 123,
		"info": gin.H{
			"a": 0,
			"b": true}})
}

func login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{"data": 101})
}
