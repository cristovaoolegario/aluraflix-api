package router

import (
	"errors"
	"github.com/cristovaoolegario/aluraflix-api/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAll_ShouldReturnEmptyVideosArrayAndOKStatusResponse_WhenTheresNoItemsToShow(t *testing.T) {

	videoService = &VideoServiceMock{}

	videoServiceMockGetAll = func() ([]models.Video, error) {
		return []models.Video{}, nil
	}

	r, _ := http.NewRequest("GET", "/api/v1/videos", nil)
	w := httptest.NewRecorder()

	GetAll(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, []byte("[]"), w.Body.Bytes())
}

func TestGetAll_ShouldReturnVideosArrayAndOKStatusResponse_WhenTheresItemsToShow(t *testing.T) {

	videoService = &VideoServiceMock{}

	videoServiceMockGetAll = func() ([]models.Video, error) {
		return []models.Video{}, errors.New("Error test")
	}

	r, _ := http.NewRequest("GET", "/api/v1/videos", nil)
	w := httptest.NewRecorder()

	GetAll(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Error test\"}"), w.Body.Bytes())
}