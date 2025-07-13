package main

import (
	"context"
	"log"

	"github.com/crimsonf09/MySite-Backend/internal/db"
	"github.com/crimsonf09/MySite-Backend/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "net/http"
)

func main() {
	godotenv.Load()
	client, err := db.InitMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("Error disconnecting MongoDB client: %v", err)
		}
	}()
	router := gin.Default()

	routes.RegisterRoutes(router)
	routes.ContactMessageRoutes(router)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	router.GET("/api/skills", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"skills": []string{"Go", "Next.js", "React", "Tailwind", "PostgreSQL", "OpenAI API"},
		})
	})

	router.Run(":8080")
}
