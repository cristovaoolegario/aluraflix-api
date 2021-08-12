package interfaces

import (
	"github.com/cristovaoolegario/aluraflix-api/dto"
	"github.com/cristovaoolegario/aluraflix-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ICategoryService interface {
	GetAll(filter string, page int64, pageSize int64) ([]models.Category, error)
	GetById(id primitive.ObjectID) (*models.Category, error)
	Create(insertCategory dto.InsertCategory) (*models.Category, error)
	Update(id primitive.ObjectID, newData dto.InsertCategory) (*models.Category, error)
	Delete(id primitive.ObjectID) error
	GetVideosByCategoryId(id primitive.ObjectID) ([]models.Video, error)
	GetFreeCategory() *models.Category
}
