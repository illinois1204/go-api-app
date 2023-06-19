package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-api-app/app/auth"
	"log"
	"net/http"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	router := gin.Default()
	auth.Routes(router)

	server := &http.Server{
		Addr:         ":5001",
		Handler:      router,
		WriteTimeout: 10 * time.Minute,
	}
	server.ListenAndServe()
}
