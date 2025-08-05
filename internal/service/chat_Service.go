package service

import (
	"context"
	"time"

	"github.com/crimsonf09/MySite-Backend/internal/db"
	"github.com/crimsonf09/MySite-Backend/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageInput struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TimeStamp time.Time          `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	IPAddess  string             `bson:"ip,omitempty" json:"ip,omitempty"`
	UID       string             `bson:"uiId,omitempty" json:"uiId,omitempty"`
	Sender    string             `bson:"sender,omitempty" json:"sender,omitempty"`
	Message   string             `bson:"message" json:"message"`
}

func GotNewMessage(input MessageInput) (*model.Chat, error) {
	newMessage := model.Chat{
		ID:        primitive.NewObjectID(),
		IPAddess:  input.IPAddess,
		UID:       input.UID,
		TimeStamp: time.Now(),
		Sender:    input.Sender,
		Message:   input.Message,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ChatCollection.InsertOne(ctx, newMessage)
	if err != nil {
		return nil, err
	}

	return &newMessage, nil
}
