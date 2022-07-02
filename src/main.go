package main

import (
	"log"
	"matech-backend/src/controllers"
	"matech-backend/src/services"
	"os"

	"github.com/gin-gonic/gin"
)

const PORT = "8080"
const ENV = "development"

func main() {
	// setup
	if os.Getenv("DATA_URL") == "" {
		log.Panic("DATA_URL is not set")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = PORT
	}
	prefix := os.Getenv("PREFIX")
	envMode := os.Getenv("ENV")
	if envMode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	matcherService := services.NewMatcherService()
	matcherController := controllers.NewMatcherController(matcherService)

	r.GET(prefix+"/get-match-sections", matcherController.GetMatchSections)
	log.Fatal(r.Run(":" + port))
}
