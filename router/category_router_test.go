package router

import (
	"errors"
	"github.com/cristovaoolegario/aluraflix-api/models"
	"github.com/stretchr/testify/assert"
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