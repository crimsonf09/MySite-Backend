package api

import (
	"context"
	"log"

	"github.com/crimsonf09/MySite-Backend/internal/db"
	"github.com/crimsonf09/MySite-Backend/internal/middleware"
	"github.com/crimsonf09/MySite-Backend/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func StartAPIServer() error {
	godotenv.Load()

	client, err := db.InitMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("MongoDB disconnect error: %v", err)
		}
	}()

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	routes.ContactMessageRoutes(router)
	routes.ProjectRoutes(router)

	log.Println("API server listening on :8080")
	return router.Run(":8080")
}
