package routes

import (
	"github.com/crimsonf09/MySite-Backend/internal/controller"
	"github.com/gin-gonic/gin"
)

func ContactMessageRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/sendContactMessage", controller.CreateContactMessageHandler)
	}
}
