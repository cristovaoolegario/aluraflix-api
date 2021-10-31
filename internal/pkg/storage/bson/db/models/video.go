package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Video represents a model of videos
type Video struct {
	ID         primitive.ObjectID `bson:"_id" json:"id" example:"000000000000000000000000"`
	CategoryID primitive.ObjectID `bson:"category_id" json:"categoriaID" example:"000000000000000000000000"`
	Titulo     string             `bson:"titulo" json:"titulo" example:"Example video"`
	Descricao  string             `bson:"descricao" json:"descricao" example:"Example description"`
	Url        string             `bson:"url" json:"url" example:"https://www.example-url.com"`
	Active     bool               `bson:"active" json:"active" example:"true"`
}

var _ interface{} = (*Video)(nil)
