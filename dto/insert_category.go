package dto

import (
	"github.com/cristovaoolegario/aluraflix-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InsertCategory struct {
	Titulo string `bson:"titulo" json:"titulo"`
	Cor    string `bson:"cor" json:"cor"`
}

func (category *InsertCategory) ConvertToCategory() models.Category {
	return models.Category{
		ID:     primitive.NewObjectID(),
		Titulo: category.Titulo,
		Cor:    category.Cor,
		Active: true,
	}
}

func (category *InsertCategory) Validate() error {
	if len(category.Titulo) == 0 {
		return MissingFieldError("Titulo")
	}
	if len(category.Cor) == 0 {
		return MissingFieldError("Cor")
	}
	return nil
}
