package mocked_services

import (
	"github.com/cristovaoolegario/aluraflix-api/dto"
	"github.com/cristovaoolegario/aluraflix-api/interfaces"
	"github.com/cristovaoolegario/aluraflix-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ interfaces.ICategoryService = (*CategoryServiceMock)(nil)

var CategoryServiceMockGetAll func(filter string, page int64, pageSize int64) ([]models.Category, error)
var CategoryServiceMockGetByID func(id primitive.ObjectID) (*models.Category, error)
var CategoryServiceMockCreate func(insertCategory dto.InsertCategory) (*models.Category, error)
var CategoryServiceMockUpdate func(id primitive.ObjectID, insertCategory dto.InsertCategory) (*models.Category, error)
var CategoryServiceMockDelete func(id primitive.ObjectID) error
var CategoryServiceMockGetVideosByCategoryId func(id primitive.ObjectID) ([]models.Video, error)
var CategoryServiceMockGetFreeCategory func() *models.Category

type CategoryServiceMock struct {}

func (cs *CategoryServiceMock) GetById(id primitive.ObjectID) (*models.Category, error) {
	return CategoryServiceMockGetByID(id)
}

func (cs *CategoryServiceMock) GetAll(filter string, page int64, pageSize int64) ([]models.Category, error){
	return CategoryServiceMockGetAll(filter, page, pageSize)
}

func (cs *CategoryServiceMock) Create(insertCategory dto.InsertCategory) (*models.Category, error) {
	return CategoryServiceMockCreate(insertCategory)
}

func (cs *CategoryServiceMock) Update(id primitive.ObjectID, newData dto.InsertCategory) (*models.Category, error) {
	return CategoryServiceMockUpdate(id, newData)
}

func (cs *CategoryServiceMock) Delete(id primitive.ObjectID) error {
	return CategoryServiceMockDelete(id)
}

func (cs *CategoryServiceMock) GetVideosByCategoryId(id primitive.ObjectID) ([]models.Video, error) {
	return CategoryServiceMockGetVideosByCategoryId(id)
}

func (cs *CategoryServiceMock) GetFreeCategory() *models.Category {
	return CategoryServiceMockGetFreeCategory()
}