package router

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/cristovaoolegario/aluraflix-api/dto"
	"github.com/cristovaoolegario/aluraflix-api/mocked_data"
	"github.com/cristovaoolegario/aluraflix-api/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllVideos_ShouldReturnVideosArrayAndOKStatusResponse_WhenTheresItemsToShow(t *testing.T) {

	videoArray := []models.Video{*mocked_data.GetValidVideo()}
	videoArrayJson, _ := json.Marshal(videoArray)
	videoService = &VideoServiceMock{}

	videoServiceMockGetAll = func() ([]models.Video, error) {
		return videoArray, nil
	}

	r, _ := http.NewRequest("GET", "/api/v1/videos", nil)
	w := httptest.NewRecorder()

	GetAllVideos(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, videoArrayJson, w.Body.Bytes())
}

func TestGetAllVideos_ShouldReturnEmptyVideosArrayAndNotFoundStatusResponse_WhenTheresNoItemsToShow(t *testing.T) {

	videoService = &VideoServiceMock{}

	videoServiceMockGetAll = func() ([]models.Video, error) {
		return nil, nil
	}

	r, _ := http.NewRequest("GET", "/api/v1/videos", nil)
	w := httptest.NewRecorder()

	GetAllVideos(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, []byte("[]"), w.Body.Bytes())
}

func TestGetAllVideos_ShouldReturnErrorAndInternalServerErrorStatusResponse_WhenTheresAnError(t *testing.T) {

	videoService = &VideoServiceMock{}

	videoServiceMockGetAll = func() ([]models.Video, error) {
		return nil, errors.New("Error test")
	}

	r, _ := http.NewRequest("GET", "/api/v1/videos", nil)
	w := httptest.NewRecorder()

	GetAllVideos(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Error test\"}"), w.Body.Bytes())
}

func TestGetVideoByID_ShouldReturnEmptyBodyAndNotFoundStatusResponse_WhenTheresNoItemsToShow(t *testing.T) {
	videoService = &VideoServiceMock{}

	videoServiceMockGetById = func(id primitive.ObjectID) (*models.Video, error) {
		return nil, errors.New("not found error")
	}
	id := primitive.NewObjectID().Hex()

	r, _ := http.NewRequest("GET", "/api/v1/videos/"+id, nil)
	w := httptest.NewRecorder()

	GetVideoByID(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Nil(t, w.Body.Bytes())
}

func TestGetVideoByID_ShouldReturnVideoInBodyAndOKStatusResponse_WhenTheresItemsToShow(t *testing.T) {
	videoService = &VideoServiceMock{}

	id := primitive.NewObjectID()
	video := mocked_data.GetValidVideoWithId(id)
	videoJson, _ := json.Marshal(video)

	videoServiceMockGetById = func(id primitive.ObjectID) (*models.Video, error) {
		return video, nil
	}

	r, _ := http.NewRequest("GET", "/api/v1/videos/"+id.Hex(), nil)
	w := httptest.NewRecorder()

	GetVideoByID(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, videoJson, w.Body.Bytes())
}

func TestCreateVideo_ShouldReturnInvalidPayloadErrorAndBadRequestStatusResponse_WhenPayloadIsInvalid(t *testing.T) {
	r, _ := http.NewRequest("POST", "/api/v1/videos", bytes.NewReader([]byte("")))
	w := httptest.NewRecorder()

	CreateVideo(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Invalid request payload\"}"), w.Body.Bytes())
}

func TestCreateVideo_ShouldReturnAnErrorAndBadRequestStatusResponse_WhenIsAnInvalidVideo(t *testing.T) {
	videoDto := mocked_data.GetInvalidInsertVideoDto()
	videoDtoJson, _ := json.Marshal(videoDto)

	r, _ := http.NewRequest("POST", "/api/v1/videos", bytes.NewReader(videoDtoJson))
	w := httptest.NewRecorder()

	CreateVideo(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Titulo is required.\"}"), w.Body.Bytes())
}

func TestCreateVideo_ShouldReturnAnErrorAndInternalServerErrorStatusResponse_WhenTheresAProblemOnVideoService(t *testing.T) {
	videoService = &VideoServiceMock{}
	videoDto := mocked_data.GetValidInsertVideoDto()
	videoDtoJson, _ := json.Marshal(videoDto)

	r, _ := http.NewRequest("POST", "/api/v1/videos", bytes.NewReader(videoDtoJson))
	w := httptest.NewRecorder()

	videoServiceMockCreate = func(dto dto.InsertVideo) (*models.Video, error) {
		return nil, errors.New("There's an error")
	}

	CreateVideo(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, []byte("{\"error\":\"There's an error\"}"), w.Body.Bytes())
}

func TestCreateVideo_ShouldReturnCreatedVideoAndCreatedStatusResponse_WhenPayloadIsOK(t *testing.T) {
	videoService = &VideoServiceMock{}
	videoModel := mocked_data.GetValidVideo()
	videoDto := mocked_data.GetValidInsertVideoDto()
	videoDtoJson, _ := json.Marshal(videoDto)
	videoModelJson, _ := json.Marshal(videoModel)

	videoServiceMockCreate = func(dto dto.InsertVideo) (*models.Video, error) {
		return videoModel, nil
	}

	r, _ := http.NewRequest("POST", "/api/v1/videos", bytes.NewReader(videoDtoJson))
	w := httptest.NewRecorder()

	CreateVideo(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, videoModelJson, w.Body.Bytes())
}

func TestUpdateVideo_ShouldReturnInvalidPayloadErrorAndBadRequestStatusResponse_WhenPayloadIsInvalid(t *testing.T) {
	r, _ := http.NewRequest("PUT", "/api/v1/videos/1", bytes.NewReader([]byte("")))
	w := httptest.NewRecorder()

	UpdateVideoByID(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Invalid request payload\"}"), w.Body.Bytes())
}

func TestUpdateVideo_ShouldReturnAnErrorAndBadRequestStatusResponse_WhenIsAnInvalidVideo(t *testing.T) {
	videoDto := mocked_data.GetInvalidInsertVideoDto()
	videoDtoJson, _ := json.Marshal(videoDto)

	r, _ := http.NewRequest("PUT", "/api/v1/videos/1", bytes.NewReader(videoDtoJson))
	w := httptest.NewRecorder()

	UpdateVideoByID(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Titulo is required.\"}"), w.Body.Bytes())
}

func TestUpdateVideo_ShouldReturnAnErrorAndInternalServerErrorStatusResponse_WhenTheresAProblemOnVideoService(t *testing.T) {
	videoService = &VideoServiceMock{}
	videoDto := mocked_data.GetValidInsertVideoDto()
	videoDtoJson, _ := json.Marshal(videoDto)

	r, _ := http.NewRequest("PUT", "/api/v1/videos/1", bytes.NewReader(videoDtoJson))
	w := httptest.NewRecorder()

	videoServiceMockUpdate = func(id primitive.ObjectID, dto dto.InsertVideo) (*models.Video, error) {
		return nil, errors.New("There's an error")
	}

	UpdateVideoByID(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, []byte("{\"error\":\"There's an error\"}"), w.Body.Bytes())
}

func TestUpdateVideo_ShouldReturnOKStatusResponse_WhenPayloadIsOK(t *testing.T) {
	videoService = &VideoServiceMock{}
	videoModel := mocked_data.GetValidVideo()
	videoDto := mocked_data.GetValidInsertVideoDto()
	videoDtoJson, _ := json.Marshal(videoDto)
	videoModelJson, _ := json.Marshal(videoModel)

	videoServiceMockUpdate = func(id primitive.ObjectID, dto dto.InsertVideo) (*models.Video, error) {
		return videoModel, nil
	}

	r, _ := http.NewRequest("PUT", "/api/v1/videos/"+videoModel.ID.Hex(), bytes.NewReader(videoDtoJson))
	w := httptest.NewRecorder()

	UpdateVideoByID(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, videoModelJson, w.Body.Bytes())
}

func TestDeleteVideo_ShouldReturnAnErrorAndInternalServerErrorStatusResponse_WhenTheresAProblemOnVideoService(t *testing.T) {
	videoService = &VideoServiceMock{}

	r, _ := http.NewRequest("DELETE", "/api/v1/videos/" + primitive.NewObjectID().Hex(), nil)
	w := httptest.NewRecorder()

	videoServiceMockDelete = func(id primitive.ObjectID) error {
		return errors.New("There's an error")
	}

	DeleteVideoByID(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, []byte("{\"error\":\"There's an error\"}"), w.Body.Bytes())
}

func TestDeleteVideo_ShouldReturnNoContentResponse_WhenTheItemCouldBeDeleted(t *testing.T) {
	videoService = &VideoServiceMock{}

	r, _ := http.NewRequest("DELETE", "/api/v1/videos/" + primitive.NewObjectID().Hex(), nil)
	w := httptest.NewRecorder()

	videoServiceMockDelete = func(id primitive.ObjectID) error {
		return nil
	}

	DeleteVideoByID(w, r)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Nil(t, w.Body.Bytes())
}