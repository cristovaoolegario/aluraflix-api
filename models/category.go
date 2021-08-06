package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID     primitive.ObjectID `bson:"_id" json:"id"`
	Titulo string             `bson:"titulo" json:"titulo"`
	Cor    string             `bson:"cor" json:"cor"`
	Active    bool          `bson:"active" json:"active"`
}
