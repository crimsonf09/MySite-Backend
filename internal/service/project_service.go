package service

import (
	"context"
	"time"

	"github.com/crimsonf09/MySite-Backend/internal/db"
	"github.com/crimsonf09/MySite-Backend/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

// func CreateProject(project model.Project) (*mongo.InsertOneResult, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	return db.ProjectCollection.InsertOne(ctx, project)
// }

func GetAllProjects() ([]model.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := db.ProjectCollection.Find(ctx, bson.M{})
	if err != nil {
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
