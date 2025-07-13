package service

import (
	"context"
	"log"
	"time"

	"github.com/crimsonf09/MySite-Backend/internal/db"
	"github.com/crimsonf09/MySite-Backend/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContactMessageInput struct {
	Message_ID  primitive.ObjectID `bson:"message_ID,omitempty" json:"message_ID"`
	Name        string             `bson:"name" json:"name"`
	Email       string             `bson:"email" json:"email"`
	CompanyName string             `bson:"companyName" json:"companyName"`
	Subject     string             `bson:"subject" json:"subject"`
	Message     string             `bson:"message" json:"message"`
}

func CreateContactMessage(input ContactMessageInput) (*model.ContactMessage, error) {
	newContactMessage := model.ContactMessage{
		ID:          primitive.NewObjectID(),
		Name:        input.Name,
		Email:       input.Email,
		CompanyName: input.CompanyName,
		Subject:     input.Subject,
		Message:     input.Message,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ContactMessageCollection.InsertOne(ctx, newContactMessage)
	if err != nil {
		return nil, err
	}
	log.Printf("Got new message %v", newContactMessage.Message)
	return &newContactMessage, nil
}
