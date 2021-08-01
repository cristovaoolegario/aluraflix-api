package mocked_data

import (
	"github.com/cristovaoolegario/aluraflix-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetValidVideo() *models.Video{
	return GetValidVideoWithId(primitive.NewObjectID())
}

func GetValidVideoWithId(id primitive.ObjectID) *models.Video{
	return &models.Video{
		ID: id,
		Titulo:  "unit test title",
		Descricao: "unit test description",
		Url: "www.unit-test.com",
		Active: true,
	}
}

func GetBsonFromVideo(model *models.Video) bson.D{
	return bson.D{
		{"_id", model.ID},
		{"titulo", model.Titulo},
		{"descricao", model.Descricao},
		{"url", model.Url},
		{"active", model.Active},
	}
}