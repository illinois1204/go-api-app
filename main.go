package main

import (
	"fmt"
	"go-api-app/app/auth"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	DEV, err := strconv.ParseBool(os.Getenv("DEV_MODE"))
	if err == nil {
		if !DEV {
			gin.SetMode(gin.ReleaseMode)
		}
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
