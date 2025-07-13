package main

import (
	"github.com/crimsonf09/MySite-Backend/internal/db"
	"github.com/crimsonf09/MySite-Backend/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "log"
	// "net/http"
)

func main() {
	godotenv.Load()
	db.InitMongoDB()

	router := gin.Default()

	routes.RegisterRoutes(router)

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
