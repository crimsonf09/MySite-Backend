package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TimeStamp time.Time          `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Sender    string             `bson:"sender,omitempty" json:"sender,omitempty"`
	IPAddess  string             `bson:"ip,omitempty" json:"ip,omitempty"`
	UID       string             `bson:"uiId,omitempty" json:"uiId,omitempty"`
	Message   string             `bson:"message" json:"message"`
}
