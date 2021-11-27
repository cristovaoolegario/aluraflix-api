package services

import (
	"context"
	"errors"

	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/http/dto"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/interfaces"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/storage/bson/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VideoService struct {
	categoryService  interfaces.ICategoryService
	videosCollection *mongo.Collection
}

func ProvideVideoService(cs CategoryService, service DatabaseService) VideoService {
	return VideoService{&cs, service.Collection(VideoCollection)}
}

func (vs *VideoService) GetAllFreeVideos() ([]models.Video, error) {
	var Videos []models.Video
	freeCategory := vs.categoryService.GetFreeCategory()
	cursor, err := vs.videosCollection.Find(context.TODO(), bson.M{"category_id": freeCategory.ID})

	if err != nil {
		return nil, err
	}
	_ = cursor.All(context.TODO(), &Videos)

	return Videos, nil
}

func (vs *VideoService) GetAll(filter string, page int64, pageSize int64) ([]models.Video, error) {
	collectionFilter, findOptions := makeFindOptions(filter, page, pageSize)
	var Videos []models.Video
	cursor, err := vs.videosCollection.Find(context.TODO(), collectionFilter, findOptions)

	if err != nil {
		return nil, err
	}
	_ = cursor.All(context.TODO(), &Videos)
	return Videos, err
}

func (vs *VideoService) GetByID(id primitive.ObjectID) (*models.Video, error) {
	Video := models.Video{}
	if err := vs.videosCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&Video); err != nil {
		return nil, err
	}
	return &Video, nil
}

func (vs *VideoService) Create(model dto.InsertVideo) (*models.Video, error) {
	convertedVideo := model.ConvertToVideo()
	if convertedVideo.CategoryID.IsZero() {
		_ = vs.categoryService.GetFreeCategory()
	}

	if _, err := vs.categoryService.GetById(convertedVideo.CategoryID); err == mongo.ErrNoDocuments {
		return nil, errors.New("Category with id " + convertedVideo.CategoryID.Hex() + " dont exists.")
	}
	_, err := vs.videosCollection.InsertOne(context.TODO(), &convertedVideo)
	if err != nil {
		return nil, err
	}
	return &convertedVideo, err
}

func (vs *VideoService) Update(id primitive.ObjectID, newData dto.InsertVideo) (*models.Video, error) {
	var video *models.Video
	if err := vs.videosCollection.FindOneAndUpdate(
		context.Background(),
		bson.D{
			primitive.E{Key: "_id", Value: id},
		},
		bson.D{primitive.E{Key: "$set", Value: newData}},
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&video); err != nil {
		return nil, err
	}
	return video, nil
}

func (vs *VideoService) Delete(id primitive.ObjectID) error {
	result, err := vs.videosCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if result.DeletedCount == 0 {
		return errors.New("no document deleted")
	}
	return err
}
