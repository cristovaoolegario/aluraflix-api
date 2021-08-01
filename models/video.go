package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Video struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Titulo    string        `bson:"titulo" json:"titulo"`
	Descricao string        `bson:"descricao" json:"descricao"`
	Url       string        `bson:"url" json:"url"`
	Active    bool          `bson:"active" json:"active"`
}

var _ interface{} = (*Video)(nil)