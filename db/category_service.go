package db

import (
	"context"
	"errors"
	"github.com/cristovaoolegario/aluraflix-api/dto"
	"github.com/cristovaoolegario/aluraflix-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ICategoryService interface {
	GetAll() ([]models.Category, error)
	GetById(id primitive.ObjectID) (*models.Category, error)
	Create(insertCategory dto.InsertCategory) (*models.Category, error)
	Update(id primitive.ObjectID, newData dto.InsertCategory) (*models.Category, error)
	Delete(id primitive.ObjectID) error
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

func (cs *CategoryService) GetById(id primitive.ObjectID) (*models.Category, error) {
	category := models.Category{}
	if err := categoriesCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&category); err != nil {
		return nil, err
	}
	return &category, nil
}

func (cs *CategoryService) Create(insertCategory dto.InsertCategory) (*models.Category, error) {
	convertedCategory := insertCategory.ConvertToCategory()
	_, err := categoriesCollection.InsertOne(context.TODO(), &convertedCategory)
	if err != nil {
		return nil, err
	}
	return &convertedCategory, nil
}

func (cs *CategoryService) Update(id primitive.ObjectID, newData dto.InsertCategory) (*models.Category, error) {
	var category *models.Category
	if err := categoriesCollection.FindOneAndUpdate(
		context.Background(),
		bson.D{
			{"_id", id},
		},
		bson.D{{"$set", newData}},
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&category); err != nil {
		return nil, err
	}
	return category, nil
}

func (cs *CategoryService) Delete(id primitive.ObjectID) error {
	result, err := categoriesCollection.DeleteOne(context.TODO(), bson.M{"_id":id})
	if result.DeletedCount == 0 {
		return errors.New("no document deleted")
	}
	return err
}