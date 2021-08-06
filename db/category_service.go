package db

import (
	"context"
	"github.com/cristovaoolegario/aluraflix-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

type ICategoryService interface {
	GetAll() ([]models.Category, error)
}

var _ ICategoryService = (*CategoryService)(nil)

type CategoryService struct{}

func (cs *CategoryService) GetAll() ([]models.Category, error) {
	var Categories []models.Category
	cursor, err := categoriesCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}
	_ = cursor.All(context.TODO(), &Categories)
	return Categories, err
}
