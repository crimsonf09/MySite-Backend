package controller

import (
	"net/http"

	"github.com/crimsonf09/MySite-Backend/internal/service"
	"github.com/gin-gonic/gin"
)

func GetAllProjectsHandler(c *gin.Context) {
	projects, err := service.GetAllProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}
	c.JSON(http.StatusOK, projects)
}

func GetProjectByIDHandler(c *gin.Context) {
	id := c.Param("id")

	project, err := service.GetProjectByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}
	c.JSON(http.StatusOK, project)
}
func CreateProjectHandler(c *gin.Context) {
	var newProject service.ProjectInput
	if err := c.ShouldBindJSON(&newProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	createdProject, err := service.CreateProject(newProject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	c.JSON(http.StatusCreated, createdProject)
}
