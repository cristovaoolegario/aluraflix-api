package mocked_services

import (
	"github.com/cristovaoolegario/aluraflix-api/dto"
	"github.com/cristovaoolegario/aluraflix-api/interfaces"
	"github.com/cristovaoolegario/aluraflix-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ interfaces.IVideoService = (*VideoServiceMock)(nil)

var VideoServiceMockGetAllFreeVideos func() ([]models.Video, error)
var VideoServiceMockGetAll func(filter string, page int64, pageSize int64) ([]models.Video, error)
var VideoServiceMockGetById func(id primitive.ObjectID) (*models.Video, error)
var VideoServiceMockCreate func(video dto.InsertVideo) (*models.Video, error)
var VideoServiceMockUpdate func(id primitive.ObjectID, newData dto.InsertVideo) (*models.Video, error)
var VideoServiceMockDelete func(id primitive.ObjectID) error

type VideoServiceMock struct{}

func (vs *VideoServiceMock) GetAllFreeVideos() ([]models.Video, error){
	return VideoServiceMockGetAllFreeVideos()
}
func (vs *VideoServiceMock) GetAll(filter string, page int64, pageSize int64) ([]models.Video, error){
	return VideoServiceMockGetAll(filter, page, pageSize)
}

func (vs *VideoServiceMock) GetByID(id primitive.ObjectID) (*models.Video, error) {
	return VideoServiceMockGetById(id)
}

func (vs *VideoServiceMock) Create(video dto.InsertVideo) (*models.Video, error) {
	return VideoServiceMockCreate(video)
}

func (vs *VideoServiceMock) Update(id primitive.ObjectID, newData dto.InsertVideo) (*models.Video, error) {
	return VideoServiceMockUpdate(id, newData)
}

func (vs *VideoServiceMock) Delete(id primitive.ObjectID) error {
	return VideoServiceMockDelete(id)
}