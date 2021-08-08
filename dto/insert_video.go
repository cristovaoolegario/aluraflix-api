package dto

import (
	"errors"
	"github.com/cristovaoolegario/aluraflix-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/url"
)

type InsertVideo struct {
	Titulo    string `bson:"titulo" json:"titulo"`
	Descricao string `bson:"descricao" json:"descricao"`
	Url       string `bson:"url" json:"url"`
}

func (video *InsertVideo) ConvertToVideo() models.Video {
	return models.Video{
		ID:        primitive.NewObjectID(),
		Titulo:    video.Titulo,
		Descricao: video.Descricao,
		Url:       video.Url,
		Active:    true,
	}
}

func (video *InsertVideo) Validate() error {
	if len(video.Titulo) == 0 {
		return MissingFieldError("Titulo")
	}
	if len(video.Descricao) == 0 {
		return MissingFieldError("Descricao")
	}
	if len(video.Url) == 0 {
		return MissingFieldError("Url")
	}
	if _, err := url.ParseRequestURI(video.Url); err != nil {
		return errors.New("Url inválida.")
	}
	return nil
}

