package dto

import (
	"errors"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/storage/bson/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/url"
)

// InsertVideo represents the DTO of a new or an updating video
type InsertVideo struct {
	Titulo     string             `json:"titulo" example:"Example video"`
	Descricao  string             `json:"descricao" example:"Example description"`
	Url        string             `json:"url" example:"https://www.example-url.com"`
	CategoryID primitive.ObjectID `json:"categoriaID" example:"000000000000000000000000"`
}

func (video *InsertVideo) ConvertToVideo() models.Video {
	return models.Video{
		ID:         primitive.NewObjectID(),
		Titulo:     video.Titulo,
		Descricao:  video.Descricao,
		Url:        video.Url,
		CategoryID: video.CategoryID,
		Active:     true,
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
