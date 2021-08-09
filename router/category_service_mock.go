package router

import (
	"github.com/cristovaoolegario/aluraflix-api/db"
	"github.com/cristovaoolegario/aluraflix-api/dto"
	"github.com/cristovaoolegario/aluraflix-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ db.ICategoryService = (*CategoryServiceMock)(nil)

var categoryServiceMockGetAll func() ([]models.Category, error)
var categoryServiceMockGetByID func(id primitive.ObjectID) (*models.Category, error)
var categoryServiceMockCreate func(insertCategory dto.InsertCategory) (*models.Category, error)
var categoryServiceMockUpdate func(id primitive.ObjectID, insertCategory dto.InsertCategory) (*models.Category, error)
var categoryServiceMockDelete func(id primitive.ObjectID) error

type CategoryServiceMock struct {}

func (cs *CategoryServiceMock) GetById(id primitive.ObjectID) (*models.Category, error) {
	return categoryServiceMockGetByID(id)
}

func (cs *CategoryServiceMock) GetAll() ([]models.Category, error){
	return categoryServiceMockGetAll()
}

func (cs *CategoryServiceMock) Create(insertCategory dto.InsertCategory) (*models.Category, error) {
	return categoryServiceMockCreate(insertCategory)
}

func (cs *CategoryServiceMock) Update(id primitive.ObjectID, newData dto.InsertCategory) (*models.Category, error) {
	return categoryServiceMockUpdate(id, newData)
}

func (cs *CategoryServiceMock) Delete(id primitive.ObjectID) error {
	return categoryServiceMockDelete(id)
}