package main

import (
	"my-microservice/config"
	"my-microservice/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Load the configuration
	config.LoadConfig()

	// Setup routes
	routes.SetupRoutes(router)

	// Run the server on port 8000
	router.Run(":8000")
}
