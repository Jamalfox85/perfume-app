package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jamalfox85/perfume-app/backend/api"
)

func main() {

	// Disable Extra Gin Logs
	gin.SetMode(gin.ReleaseMode)

	// Create app instance
	app := api.NewApplication()

	// Set server port and start
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not set
	}

	// Create server instance
	server := api.NewServer(port)
	server.Start(app)
}