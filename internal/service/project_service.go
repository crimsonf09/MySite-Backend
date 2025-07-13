package service

import (
	"context"
	"log"
	"time"

	"github.com/crimsonf09/MySite-Backend/internal/db"
	"github.com/crimsonf09/MySite-Backend/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectInput struct {
	Title            string   `json:"title" binding:"required"`
	ShortDescription string   `json:"shortDescription" binding:"required"`
	Description      string   `json:"description" binding:"required"`
	TechStack        []string `json:"techStack" binding:"required"`
	URL              string   `json:"url" binding:"required"`
	Type             []string `json:"type" binding:"required"`
	Images           []string `json:"images,omitempty"`
	ShortMedia       string   `json:"shortMedia,omitempty"`
}

// CreateProject inserts a new project into MongoDB
func CreateProject(input ProjectInput) (*model.Project, error) {
	newProject := model.Project{
		ID:               primitive.NewObjectID(),
		Title:            input.Title,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		TechStack:        input.TechStack,
		URL:              input.URL,
		Type:             input.Type,
		Images:           input.Images,
		ShortMedia:       input.ShortMedia,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ProjectCollection.InsertOne(ctx, newProject)
	if err != nil {
		return nil, err
	}

	return &newProject, nil
}

func GetAllProjects() ([]model.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := db.ProjectCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("GetAllProjects error: %v", err) // Log error server-side
		return nil, err
	}
	defer cursor.Close(ctx)

	var projects []model.Project
	for cursor.Next(ctx) {
		var project model.Project
		if err := cursor.Decode(&project); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}

func GetProjectByID(id string) (*model.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var project model.Project
	err = db.ProjectCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}
