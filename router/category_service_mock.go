package router

import (
	"github.com/cristovaoolegario/aluraflix-api/db"
	"github.com/cristovaoolegario/aluraflix-api/models"
)

var _ db.ICategoryService = (*CategoryServiceMock)(nil)

var categoryServiceMockGetAll func() ([]models.Category, error)

type CategoryServiceMock struct {}

func (cs *CategoryServiceMock) GetAll() ([]models.Category, error){
	return categoryServiceMockGetAll()
}