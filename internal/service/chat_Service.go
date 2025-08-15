package service

import (
	"context"
	"log"
	"time"

	"github.com/crimsonf09/MySite-Backend/internal/db"
	"github.com/crimsonf09/MySite-Backend/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageInput struct {
	TimeStamp time.Time `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	IPAddress string    `bson:"ip,omitempty" json:"ip,omitempty"`
	UID       string    `bson:"uiId,omitempty" json:"uiId,omitempty"`
	Sender    string    `bson:"sender,omitempty" json:"sender,omitempty"`
	Message   string    `bson:"message" json:"message"`
}

// Save a message to MongoDB
func SaveMessage(message model.Chat) (*model.Chat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ChatCollection.InsertOne(ctx, message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

// Process new user message and generate bot response
func GotNewMessage(input MessageInput) (*model.Chat, *model.Chat, error) {
	// Save user message
	userMessage := model.Chat{
		ID:        primitive.NewObjectID(),
		IPAddress: input.IPAddress,
		UID:       input.UID,
		TimeStamp: time.Now(),
		Sender:    input.Sender,
		Message:   input.Message,
	}

	savedUserMsg, err := SaveMessage(userMessage)
	if err != nil {
		return nil, nil, err
	}

	// Generate bot response
	botRes := ResponseMessage(input)
	botMessage := model.Chat{
		ID:        primitive.NewObjectID(),
		IPAddress: "1.1.1.1",
		UID:       input.UID,
		TimeStamp: time.Now(),
		Sender:    "Armin",
		Message:   botRes,
	}

	savedBotMsg, err := SaveMessage(botMessage)
	if err != nil {
		return savedUserMsg, nil, err
	}

	return savedUserMsg, savedBotMsg, nil
}

// Generate bot response text
func ResponseMessage(input MessageInput) string {
	log.Printf("User: %s", input.Message)
	const res = "Good morning, Armin here! How can I assist you today?"
	return res
}
