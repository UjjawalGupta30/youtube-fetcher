package main

import (
	"log"
	"youtube-fetcher/config"
	"youtube-fetcher/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()
	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run(":8080")
}
