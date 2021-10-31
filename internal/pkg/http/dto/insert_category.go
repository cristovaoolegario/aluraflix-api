package dto

import (
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/storage/bson/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertCategory represents the DTO of a new or an updating category
type InsertCategory struct {
	Titulo string `json:"titulo" example:"Example video"`
	Cor    string `json:"cor" example:"blue"`
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
