package dto

import (
	"github.com/cristovaoolegario/aluraflix-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InsertVideo struct{
	Titulo    string        `bson:"titulo" json:"titulo"`
	Descricao string        `bson:"descricao" json:"descricao"`
	Url       string        `bson:"url" json:"url"`
	Active    bool          `bson:"active" json:"active"`
}

func (video *InsertVideo) ConvertToVideo() models.Video {
	return models.Video{
		ID: primitive.NewObjectID(),
		Titulo: video.Titulo,
		Descricao: video.Descricao,
		Url: video.Url,
	}
}