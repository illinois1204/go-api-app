package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-api-app/app/auth"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	gin.DefaultWriter = io.Discard
	router := gin.Default()
	auth.Routes(router)

	PORT := os.Getenv("PORT")
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", PORT),
		Handler:      router,
		WriteTimeout: 10 * time.Minute,
	}
	fmt.Printf("Server starting on port: %s\n", PORT)
	server.ListenAndServe()
}
