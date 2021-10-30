package mocked_data

import (
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/http/dto"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/storage/bson/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetValidVideo() *models.Video {
	return GetValidVideoWithId(primitive.NewObjectID())
}

func GetValidVideoWithId(id primitive.ObjectID) *models.Video {
	return &models.Video{
		ID:        id,
		Titulo:    "unit test title",
		Descricao: "unit test description",
		Url:       "www.unit-test.com",
		Active:    true,
	}
}

func GetBsonFromVideo(model *models.Video) bson.D {
	return bson.D{
		{"_id", model.ID},
		{"titulo", model.Titulo},
		{"descricao", model.Descricao},
		{"url", model.Url},
		{"active", model.Active},
	}
}

func GetValidInsertVideoDto() dto.InsertVideo {
	return dto.InsertVideo{
		Titulo:    "unit test title",
		Descricao: "unit test description",
		Url:       "https://www.unit-test.com",
	}
}

func GetInvalidInsertVideoDto() dto.InsertVideo {
	return dto.InsertVideo{
		Titulo:    "",
		Descricao: "unit test description",
		Url:       "www.unit-test.com",
	}
}