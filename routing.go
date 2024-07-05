package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	ports "youtube_service_arq_hex/youtube-service/ports/drivers"

)

func CreateRouters(youtubeAdapter ports.ForYouTube) *gin.Engine {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set Gin mode based on environment variable
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "debug"
	}
	//gin.SetMode(ginMode)

	// Create a new router
	router := gin.Default()

	// Define a route for searching YouTube
	router.GET("/youtube/search", func(gc *gin.Context) {
			youtubeAdapter.GetSearch(gc)
		})

	return router
}
