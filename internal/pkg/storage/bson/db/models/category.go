package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Category represents a model of categories
type Category struct {
	ID     primitive.ObjectID `bson:"_id" json:"id" example:"000000000000000000000000"`
	Titulo string             `bson:"titulo" json:"titulo" example:"Example category"`
	Cor    string             `bson:"cor" json:"cor" example:"Red"`
	Active bool               `bson:"active" json:"active" example:"true"`
}

func GetFreeCategory() *Category {
	return &Category{
		ID:     primitive.ObjectID{},
		Titulo: "FREE",
		Cor:    "FREE",
		Active: true,
	}
}
