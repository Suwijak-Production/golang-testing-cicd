package app

import (
	"log"
	"myapp/internal/config"
	"myapp/internal/repository"
	"myapp/internal/router"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	// Set Trusted Proxies
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	config.Load()         // Load configuration
	repository.Connect()  // Connect to the database
	router.SetupRoutes(r) // Setup routes

	err = r.Run(config.Config.ServerAddress)
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
