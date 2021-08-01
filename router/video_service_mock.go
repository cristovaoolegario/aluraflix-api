package router

import (
	"github.com/cristovaoolegario/aluraflix-api/db"
	"github.com/cristovaoolegario/aluraflix-api/models"
)

var _ db.IVideoService = (*VideoServiceMock)(nil)

var videoServiceMockGetAll func() ([]models.Video, error)

type VideoServiceMock struct{}

func (vs *VideoServiceMock) GetAll() ([]models.Video, error){
	return videoServiceMockGetAll()
}