package db

import (
	"context"
	"errors"
	"github.com/cristovaoolegario/aluraflix-api/dto"
	"github.com/cristovaoolegario/aluraflix-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CategoryService struct {
	categoryCollection *mongo.Collection
	videosCollection   *mongo.Collection
}

func ProvideCategoryService(db DatabaseService) CategoryService {
	return CategoryService{db.Collection(CategoriesCollection), db.Collection(VideoCollection)}
}

func (cs *CategoryService) GetAll(filter string, page int64, pageSize int64) ([]models.Category, error) {
	collectionFilter, findOptions := makeFindOptions(filter, page, pageSize)
	var Categories []models.Category
	cursor, err := cs.categoryCollection.Find(context.TODO(), collectionFilter, findOptions)

	if err != nil {
		return nil, err
	}
	_ = cursor.All(context.TODO(), &Categories)
	return Categories, err
}

func (cs *CategoryService) GetById(id primitive.ObjectID) (*models.Category, error) {
	category := models.Category{}
	if err := cs.categoryCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&category); err != nil {
		return nil, err
	}
	return &category, nil
}

func (cs *CategoryService) Create(insertCategory dto.InsertCategory) (*models.Category, error) {
	convertedCategory := insertCategory.ConvertToCategory()
	_, err := cs.categoryCollection.InsertOne(context.TODO(), &convertedCategory)
	if err != nil {
		return nil, err
	}
	return &convertedCategory, nil
}

func (cs *CategoryService) Update(id primitive.ObjectID, newData dto.InsertCategory) (*models.Category, error) {
	var category *models.Category
	if err := cs.categoryCollection.FindOneAndUpdate(
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
	result, err := cs.categoryCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if result.DeletedCount == 0 {
		return errors.New("no document deleted")
	}
	return err
}

func (cs *CategoryService) GetVideosByCategoryId(id primitive.ObjectID) ([]models.Video, error) {
	var videos []models.Video
	cursor, err := cs.videosCollection.Find(context.TODO(), bson.M{"category_id": id})
	if err != nil {
		return nil, err
	}
	_ = cursor.All(context.TODO(), &videos)

	return videos, err
}

func (cs *CategoryService) GetFreeCategory() *models.Category {
	category := models.Category{}
	if err := cs.categoryCollection.FindOne(context.TODO(), bson.M{"titulo": "FREE"}).Decode(&category); err != nil {
		category = *models.GetFreeCategory()
		_, err := cs.categoryCollection.InsertOne(context.TODO(), &category)
		if err != nil {
			return nil
		}
		return &category
	}
	return &category
}
