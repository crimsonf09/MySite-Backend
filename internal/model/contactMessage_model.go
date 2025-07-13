package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ContactMessage struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name        string             `bson:"name" json:"name"`
	Email       string             `bson:"email" json:"email"`
	CompanyName string             `bson:"companyName" json:"companyName"`
	Subject     string             `bson:"subject" json:"subject"`
	Message     string             `bson:"message" json:"message"`
}
