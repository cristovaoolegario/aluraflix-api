package interfaces

import (
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/http/dto"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/storage/bson/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IVideoService interface {
	GetAllFreeVideos() ([]models.Video, error)
	GetAll(filter string, page int64, pageSize int64) ([]models.Video, error)
	GetByID(id primitive.ObjectID) (*models.Video, error)
	Create(video dto.InsertVideo) (*models.Video, error)
	Update(id primitive.ObjectID, newData dto.InsertVideo) (*models.Video, error)
	Delete(id primitive.ObjectID) error
}