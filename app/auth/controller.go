package auth

import (
	"fmt"
	"go-api-app/app/auth/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func register(c *gin.Context) {

	var user dto.UserCreate
	// if err := c.BindJSON(&user); err != nil {
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		fmt.Println(11111)

		oke, ok := err.(*validator.InvalidValidationError)
		fmt.Println(oke)
		fmt.Println(ok)
		// ok {
		// 	fmt.Println(err)
		// 	fmt.Println(22222)
		// 	return
		// }

		// for _, err := range err.(validator.ValidationErrors) {
		// 	fmt.Println(err.Namespace())
		// 	fmt.Println(err.Field())
		// 	fmt.Println(err.StructNamespace())
		// 	fmt.Println(err.StructField())
		// 	fmt.Println(err.Tag())
		// 	fmt.Println(err.ActualTag())
		// 	fmt.Println(err.Kind())
		// 	fmt.Println(err.Type())
		// 	fmt.Println(err.Value())
		// 	fmt.Println(err.Param())
		// }
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("user: %v\n", user)
	c.JSON(http.StatusOK, gin.H{
		"data": 123,
		"info": gin.H{
			"a": 0,
			"b": true}})
}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{"data": 101})
}
