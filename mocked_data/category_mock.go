package mocked_data

import (
	"github.com/cristovaoolegario/aluraflix-api/dto"
	"github.com/cristovaoolegario/aluraflix-api/models"
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
		{"_id", model.ID},
		{"titulo", model.Titulo},
		{"cor", model.Cor},
		{"active", model.Active},
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
