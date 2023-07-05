package auth

import (
	"github.com/gin-gonic/gin"
	"go-api-app/app/auth/dto"
	"net/http"
)

func register(ctx *gin.Context) {
	body, _ := ctx.MustGet("body").(authDto.UserCreate)
	ctx.JSON(http.StatusOK, body)
}

func login(ctx *gin.Context) {
	query, _ := ctx.MustGet("query").(authDto.Filter)
	ctx.JSON(http.StatusOK, query)
	// ctx.JSON(http.StatusOK, map[string]interface{}{"data": 101})
}
