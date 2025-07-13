package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title            string             `bson:"title" json:"title"`
	ShortDescription string             `bson:"shortDescription" json:"shortDescription"`
	Description      string             `bson:"description" json:"description"`
	TechStack        []string           `bson:"techStack" json:"techStack"`
	URL              string             `bson:"url" json:"url"`
	Type             []string           `bson:"type" json:"type"`
	Images           []string           `bson:"images,omitempty" json:"images,omitempty"`
	ShortMedia       string             `bson:"shortMedia,omitempty" json:"shortMedia,omitempty"`
}
