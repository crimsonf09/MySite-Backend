package controller

import (
	"net/http"

	"github.com/crimsonf09/MySite-Backend/internal/service"
	"github.com/gin-gonic/gin"
)

func CreateContactMessageHandler(c *gin.Context) {
	var newContactMessage service.ContactMessageInput
	if err := c.ShouldBindJSON(&newContactMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	createdContactMesage, err := service.CreateContactMessage(newContactMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send message"})
		return
	}
	c.JSON(http.StatusCreated, createdContactMesage)
}
