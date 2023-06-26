package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-api-app/app/auth/dto"
	"net/http"
)

var validate = validator.New()

func register(ctx *gin.Context) {
	var user dto.UserCreate

	// if err := ctx.ShouldBindJSON(&user); err != nil {
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(user); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
		}
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("user: %v\n", user)

	ctx.JSON(http.StatusOK, gin.H{
		"data": 123,
		"info": gin.H{
			"a": 0,
			"b": true}})
}

func login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{"data": 101})
}
