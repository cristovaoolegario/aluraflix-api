package interfaces

import (
	"github.com/cristovaoolegario/aluraflix-api/dto"
	"github.com/cristovaoolegario/aluraflix-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IVideoService interface {
	GetAll(filter string) ([]models.Video, error)
	GetByID(id primitive.ObjectID) (*models.Video, error)
	Create(video dto.InsertVideo) (*models.Video, error)
	Update(id primitive.ObjectID, newData dto.InsertVideo) (*models.Video, error)
	Delete(id primitive.ObjectID) error
}