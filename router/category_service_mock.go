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