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

func InitMongoDB() {
	host := os.Getenv("MONGODB_HOST")           // e.g. "localhost"
	port := os.Getenv("MONGODB_PORT")           // e.g. "27017"
	dbName := os.Getenv("MONGODB_DB")           // e.g. "myDB"
	collName := os.Getenv("MONGODB_COLLECTION") // e.g. "projects"

	uri := fmt.Sprintf("mongodb://%s:%s", host, port)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	ProjectCollection = client.Database(dbName).Collection(collName)
	log.Printf("MongoDB connected at %s, using %s.%s", uri, dbName, collName)
}
