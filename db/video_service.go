package db

import (
	"context"
	"github.com/cristovaoolegario/aluraflix-api/dto"
	"github.com/cristovaoolegario/aluraflix-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IVideoService interface {
    GetAll() ([]models.Video, error)
	GetByID(id primitive.ObjectID) (*models.Video, error)
	Create(video dto.InsertVideo) (*models.Video, error)
    Update(id primitive.ObjectID, newData dto.InsertVideo) (*models.Video, error)
}

var _ IVideoService = (*VideoService)(nil)

type VideoService struct {}

func (vs *VideoService) GetAll() ([]models.Video, error){
	var Videos []models.Video
	cursor, err := videosCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil,err
	}
	_ = cursor.All(context.TODO(), &Videos)
	return Videos, err
}

func (vs *VideoService) GetByID(id primitive.ObjectID) (*models.Video, error) {
	Video := models.Video{}
	if err := videosCollection.FindOne(context.TODO(), bson.M{"_id":id}).Decode(&Video); err != nil{
		return nil, err
	}
	return &Video, nil
}

func (vs *VideoService) Create(model dto.InsertVideo) (*models.Video, error) {
	convertedVideo := model.ConvertToVideo()
	_, err := videosCollection.InsertOne(context.TODO(), &convertedVideo)
	if err != nil {
		return nil, err
	}
	return &convertedVideo, err
}

func (vs *VideoService) Update(id primitive.ObjectID, newData dto.InsertVideo) (*models.Video, error) {
	var video *models.Video
	if err := videosCollection.FindOneAndUpdate(
		context.Background(),
		bson.D{
			{"_id", id},
		},
		bson.D{{"$set", newData}},
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&video); err != nil {
		return nil, err
	}
	return video, nil
}
