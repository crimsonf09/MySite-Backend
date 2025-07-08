package main

import (
	"github.com/gin-gonic/gin"
	// "log"
	// "net/http"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
		router.GET("/api/skills", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"skills": []string{"Go", "Next.js", "React", "Tailwind", "PostgreSQL", "OpenAI API"},
			})
		})
	})
	router.Run(":8080")
}
