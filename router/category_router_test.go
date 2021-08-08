package router

import (
	"encoding/json"
	"errors"
	"github.com/cristovaoolegario/aluraflix-api/mocked_data"
	"github.com/cristovaoolegario/aluraflix-api/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllCategories_ShouldReturnEmptyCategoryArrayAndOKStatusResponse_WhenTheresNoItemsToShow(t *testing.T) {

	categoryService = &CategoryServiceMock{}

	categoryServiceMockGetAll = func() ([]models.Category, error) {
		return []models.Category{}, nil
	}

	r, _ := http.NewRequest("GET", "/api/v1/category", nil)
	w := httptest.NewRecorder()

	GetAllCategories(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, []byte("[]"), w.Body.Bytes())
}

func TestGetAllCategories_ShouldReturnErrorAndInternalServerErrorStatusResponse_WhenTheresAnError(t *testing.T) {

	categoryService = &CategoryServiceMock{}

	categoryServiceMockGetAll = func() ([]models.Category, error) {
		return []models.Category{}, errors.New("Error test")
	}

	r, _ := http.NewRequest("GET", "/api/v1/category", nil)
	w := httptest.NewRecorder()

	GetAllCategories(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Error test\"}"), w.Body.Bytes())
}

func TestGetCategoryByID_ShouldReturnEmptyCategoryAnd404StatusResponse_WhenTheresIsNoCategoryWithThatId(t *testing.T) {
	categoryService = &CategoryServiceMock{}
	id := primitive.NewObjectID().Hex()

	categoryServiceMockGetByID = func(id primitive.ObjectID) (*models.Category, error) {
		return nil, errors.New("Error test")
	}

	r, _ := http.NewRequest("GET", "/api/v1/categories/"+id, nil)
	w := httptest.NewRecorder()

	GetCategoryByID(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Nil(t, w.Body.Bytes())
}

func TestGetCategoryByID_ShouldReturnCategoryInBodyAndOKStatusResponse_WhenTheresItemsToShow(t *testing.T) {
	categoryService = &CategoryServiceMock{}
	id := primitive.NewObjectID()
	category := mocked_data.GetValidCategoryWithId(id)
	categoryJson, _ := json.Marshal(category)

	categoryServiceMockGetByID = func(id primitive.ObjectID) (*models.Category, error) {
		return category, nil
	}

	r, _ := http.NewRequest("GET", "/api/v1/categories/"+id.Hex(), nil)
	w := httptest.NewRecorder()

	GetCategoryByID(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, categoryJson, w.Body.Bytes())
}
