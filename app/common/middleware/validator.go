package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate = validator.New()

type _IError struct {
	Field string  `json:"field"`
	Tag   string  `json:"tag"`
	Param *string `json:"param"`
}

func errorParser(dto any) []_IError {
	var errors []_IError
	if err := validate.Struct(dto); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var errorItem _IError
			errorItem.Field = err.Field()
			errorItem.Tag = err.ActualTag()
			if paramMSG := err.Param(); paramMSG != "" {
				var msgP *string = &paramMSG
				errorItem.Param = msgP
			}
			errors = append(errors, errorItem)
		}
	}
	return errors
}

func ValidateBody[T any](ctx *gin.Context) {
	var dto T
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusTeapot, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	var errors []_IError = errorParser(dto)
	if len(errors) != 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": errors})
		ctx.Abort()
		return
	}

	ctx.Set("body", dto)
	ctx.Next()
}

func ValidateQuery[T any](ctx *gin.Context) {
	var dto T
	if err := ctx.ShouldBindQuery(&dto); err != nil {
		ctx.JSON(http.StatusTeapot, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	var errors []_IError = errorParser(dto)
	if len(errors) != 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": errors})
		ctx.Abort()
		return
	}

	ctx.Set("query", dto)
	ctx.Next()
}
