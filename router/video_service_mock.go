package router

import (
	"github.com/cristovaoolegario/aluraflix-api/db"
	"github.com/cristovaoolegario/aluraflix-api/dto"
	"github.com/cristovaoolegario/aluraflix-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ db.IVideoService = (*VideoServiceMock)(nil)

var videoServiceMockGetAll func() ([]models.Video, error)
var videoServiceMockGetById func(id primitive.ObjectID) (*models.Video, error)
var videoServiceMockCreate func(video dto.InsertVideo) (*models.Video, error)

type VideoServiceMock struct{}

func (vs *VideoServiceMock) GetAll() ([]models.Video, error){
	return videoServiceMockGetAll()
}

func (vs *VideoServiceMock) GetByID(id primitive.ObjectID) (*models.Video, error) {
	return videoServiceMockGetById(id)
}

func (vs *VideoServiceMock) Create(video dto.InsertVideo) (*models.Video, error) {
	return videoServiceMockCreate(video)
}