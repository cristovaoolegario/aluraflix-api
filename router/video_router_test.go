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

func TestGetByID_ShouldReturnEmptyBodyAndNotFoundStatusResponse_WhenTheresNoItemsToShow(t *testing.T) {
	videoService = &VideoServiceMock{}

	videoServiceMockGetById = func(id primitive.ObjectID) (*models.Video, error) {
		return nil, errors.New("not found error")
	}
	id := primitive.NewObjectID().Hex()

	r, _ := http.NewRequest("GET", "/api/v1/videos/"+id, nil)
	w := httptest.NewRecorder()

	GetByID(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Nil(t, w.Body.Bytes())
}

func TestGetByID_ShouldReturnVideoInBodyAndOKStatusResponse_WhenTheresItemsToShow(t *testing.T) {
	videoService = &VideoServiceMock{}

	id := primitive.NewObjectID()
	video := mocked_data.GetValidVideoWithId(id)
	videoJson, _ := json.Marshal(video)

	videoServiceMockGetById = func(id primitive.ObjectID) (*models.Video, error) {
		return video, nil
	}

	r, _ := http.NewRequest("GET", "/api/v1/videos/"+id.Hex(), nil)
	w := httptest.NewRecorder()

	GetByID(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, videoJson, w.Body.Bytes())
}