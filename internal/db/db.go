package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ProjectCollection *mongo.Collection
var ContactMessageCollection *mongo.Collection
var ChatCollection *mongo.Collection

func InitMongoDB() (*mongo.Client, error) {
	mongoPassword := os.Getenv("MONGODB_PASSWORD")
	dbName := os.Getenv("MONGODB_DB")
	if mongoPassword == "" || dbName == "" {
		return nil, fmt.Errorf("missing required env vars: MONGODB_PASSWORD or MONGODB_DB")
	}

	uri := fmt.Sprintf("mongodb+srv://natouchf40_db_user:%s@base.rnpuo4t.mongodb.net/?retryWrites=true&w=majority&appName=Base", mongoPassword)

	// Setup client options
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Connect with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("mongo connect error: %w", err)
	}

	// Test connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("mongo ping error: %w", err)
	}

	// Select collections
	db := client.Database(dbName)
	ContactMessageCollection = db.Collection("messages")
	ProjectCollection = db.Collection("projects")
	ChatCollection = db.Collection("chats")

	return client, nil
}
