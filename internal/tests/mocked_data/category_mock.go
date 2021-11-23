package mocked_data

import (
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/http/dto"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/storage/bson/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetValidCategory() *models.Category {
	return GetValidCategoryWithId(primitive.NewObjectID())
}

func GetValidCategoryWithId(id primitive.ObjectID) *models.Category {
	return &models.Category{
		ID:     id,
		Titulo: "unit test title",
		Cor:    "blue",
		Active: true,
	}
}

func GetBsonFromCategory(model *models.Category) bson.D {
	return bson.D{
		primitive.E{Key: "_id", Value: model.ID},
		primitive.E{Key: "titulo", Value: model.Titulo},
		primitive.E{Key: "cor", Value: model.Cor},
		primitive.E{Key: "active", Value: model.Active},
	}
}

func GetInvalidInsertCategoryDto() dto.InsertCategory {
	return dto.InsertCategory{
		Titulo: "",
		Cor:    "blur",
	}
}
func GetValidInsertCategoryDto() dto.InsertCategory {
	return dto.InsertCategory{
		Titulo: "unit test title",
		Cor:    "blur",
	}
}
