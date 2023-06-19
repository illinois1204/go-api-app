package auth

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{"data": 123})
}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{"data": 101})
}
