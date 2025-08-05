package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ProjectCollection *mongo.Collection
var ContactMessageCollection *mongo.Collection
var ChatCollection *mongo.Collection

func InitMongoDB() (*mongo.Client, error) {
	host := os.Getenv("MONGODB_HOST")
	port := os.Getenv("MONGODB_PORT")
	dbName := os.Getenv("MONGODB_DB")

	if host == "" || port == "" || dbName == "" {
		return nil, fmt.Errorf("one or more required env vars are empty")
	}

	uri := fmt.Sprintf("mongodb://%s:%s", host, port)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("mongo connect error: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("mongo ping error: %w", err)
	}

	db := client.Database(dbName)
	ContactMessageCollection = db.Collection("messages")
	ProjectCollection = db.Collection("projects")

	log.Printf("MongoDB connected at %s, using database: %s", uri, dbName)
	return client, nil
}
