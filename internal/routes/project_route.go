package routes

import (
	"github.com/crimsonf09/MySite-Backend/internal/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/projects", controller.GetAllProjectsHandler)
		api.GET("/projects/:id", controller.GetProjectByIDHandler)
		api.POST("/projects", controller.CreateProjectHandler)
	}
}
