package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"oms/config"
	"oms/repository"
	"oms/routes"
)

func main() {
	if err := config.ConnectDatabase(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if config.DB == nil {
		log.Fatal("Database client is nil after connection")
	}
	repository.IntializeCollections()

	router := gin.Default()
	//routes.SetRoutes(router)

	routes.SetupRoutes(router)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
