package resources

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/http/dto"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/storage/bson/db/models"
	"github.com/cristovaoolegario/aluraflix-api/internal/tests/mocked_data"
	"github.com/cristovaoolegario/aluraflix-api/internal/tests/mocked_services"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllFreeVideos_ShouldReturnFreeVideosArrayAndOKStatusResponse_WhenTheresItemsToShow(t *testing.T) {
	var router = VideoRouter{}

	videoArray := []models.Video{*mocked_data.GetValidVideo()}
	videoArrayJson, _ := json.Marshal(videoArray)
	router.service = &mocked_services.VideoServiceMock{}

	mocked_services.VideoServiceMockGetAllFreeVideos = func() ([]models.Video, error) {
		return videoArray, nil
	}

	r, _ := http.NewRequest("GET", "/api/v1/videos/free", nil)
	w := httptest.NewRecorder()

	router.GetAllFreeVideos(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, videoArrayJson, w.Body.Bytes())
}

func TestGetAllFreeVideos_ShouldReturnEmptyVideosArrayAndNotFoundStatusResponse_WhenTheresNoItemsToShow(t *testing.T) {
	var router = VideoRouter{}
	router.service = &mocked_services.VideoServiceMock{}

	mocked_services.VideoServiceMockGetAllFreeVideos = func() ([]models.Video, error) {
		return nil, nil
	}

	r, _ := http.NewRequest("GET", "/api/v1/videos/free", nil)
	w := httptest.NewRecorder()

	router.GetAllFreeVideos(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, []byte("[]"), w.Body.Bytes())
}

func TestGetAllFreeVideos_ShouldReturnErrorAndInternalServerErrorStatusResponse_WhenTheresAnError(t *testing.T) {
	var router = VideoRouter{}
	router.service = &mocked_services.VideoServiceMock{}

	mocked_services.VideoServiceMockGetAllFreeVideos = func() ([]models.Video, error) {
		return nil, errors.New("Error test")
	}

	r, _ := http.NewRequest("GET", "/api/v1/videos/free", nil)
	w := httptest.NewRecorder()

	router.GetAllFreeVideos(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Error test\"}"), w.Body.Bytes())
}

func TestGetAllVideos_ShouldReturnVideosArrayAndOKStatusResponse_WhenTheresItemsToShow(t *testing.T) {
	var router = VideoRouter{}
	videoArray := []models.Video{*mocked_data.GetValidVideo()}
	videoArrayJson, _ := json.Marshal(videoArray)
	router.service = &mocked_services.VideoServiceMock{}

	mocked_services.VideoServiceMockGetAll = func(filter string, page int64, pageSize int64) ([]models.Video, error) {
		return videoArray, nil
	}

	r, _ := http.NewRequest("GET", "/api/v1/videos?page=1&pageSize=5", nil)
	w := httptest.NewRecorder()

	router.GetAllVideos(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, videoArrayJson, w.Body.Bytes())
}

func TestGetAllVideos_ShouldReturnEmptyVideosArrayAndNotFoundStatusResponse_WhenTheresNoItemsToShow(t *testing.T) {
	var router = VideoRouter{}
	router.service = &mocked_services.VideoServiceMock{}

	mocked_services.VideoServiceMockGetAll = func(filter string, page int64, pageSize int64) ([]models.Video, error) {
		return nil, nil
	}

	r, _ := http.NewRequest("GET", "/api/v1/videos", nil)
	w := httptest.NewRecorder()

	router.GetAllVideos(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, []byte("[]"), w.Body.Bytes())
}

func TestGetAllVideos_ShouldReturnErrorAndInternalServerErrorStatusResponse_WhenTheresAnError(t *testing.T) {
	var router = VideoRouter{}
	router.service = &mocked_services.VideoServiceMock{}

	mocked_services.VideoServiceMockGetAll = func(filter string, page int64, pageSize int64) ([]models.Video, error) {
		return nil, errors.New("Error test")
	}

	r, _ := http.NewRequest("GET", "/api/v1/videos", nil)
	w := httptest.NewRecorder()

	router.GetAllVideos(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Error test\"}"), w.Body.Bytes())
}

func TestGetVideoByID_ShouldReturnEmptyBodyAndNotFoundStatusResponse_WhenTheresNoItemsToShow(t *testing.T) {
	var router = VideoRouter{}
	router.service = &mocked_services.VideoServiceMock{}

	mocked_services.VideoServiceMockGetById = func(id primitive.ObjectID) (*models.Video, error) {
		return nil, errors.New("not found error")
	}
	id := primitive.NewObjectID().Hex()

	r, _ := http.NewRequest("GET", "/api/v1/videos/"+id, nil)
	w := httptest.NewRecorder()

	router.GetVideoByID(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Nil(t, w.Body.Bytes())
}

func TestGetVideoByID_ShouldReturnVideoInBodyAndOKStatusResponse_WhenTheresItemsToShow(t *testing.T) {
	var router = VideoRouter{}
	router.service = &mocked_services.VideoServiceMock{}

	id := primitive.NewObjectID()
	video := mocked_data.GetValidVideoWithId(id)
	videoJson, _ := json.Marshal(video)

	mocked_services.VideoServiceMockGetById = func(id primitive.ObjectID) (*models.Video, error) {
		return video, nil
	}

	r, _ := http.NewRequest("GET", "/api/v1/videos/"+id.Hex(), nil)
	w := httptest.NewRecorder()

	router.GetVideoByID(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, videoJson, w.Body.Bytes())
}

func TestCreateVideo_ShouldReturnInvalidPayloadErrorAndBadRequestStatusResponse_WhenPayloadIsInvalid(t *testing.T) {
	var router = VideoRouter{}
	r, _ := http.NewRequest("POST", "/api/v1/videos", bytes.NewReader([]byte("")))
	w := httptest.NewRecorder()

	router.CreateVideo(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Invalid request payload\"}"), w.Body.Bytes())
}

func TestCreateVideo_ShouldReturnAnErrorAndBadRequestStatusResponse_WhenIsAnInvalidVideo(t *testing.T) {
	var router = VideoRouter{}
	videoDto := mocked_data.GetInvalidInsertVideoDto()
	videoDtoJson, _ := json.Marshal(videoDto)

	r, _ := http.NewRequest("POST", "/api/v1/videos", bytes.NewReader(videoDtoJson))
	w := httptest.NewRecorder()

	router.CreateVideo(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Titulo is required.\"}"), w.Body.Bytes())
}

func TestCreateVideo_ShouldReturnAnErrorAndInternalServerErrorStatusResponse_WhenTheresAProblemOnVideoService(t *testing.T) {
	var router = VideoRouter{}
	router.service = &mocked_services.VideoServiceMock{}
	videoDto := mocked_data.GetValidInsertVideoDto()
	videoDtoJson, _ := json.Marshal(videoDto)

	r, _ := http.NewRequest("POST", "/api/v1/videos", bytes.NewReader(videoDtoJson))
	w := httptest.NewRecorder()

	mocked_services.VideoServiceMockCreate = func(dto dto.InsertVideo) (*models.Video, error) {
		return nil, errors.New("There's an error")
	}

	router.CreateVideo(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, []byte("{\"error\":\"There's an error\"}"), w.Body.Bytes())
}

func TestCreateVideo_ShouldReturnCreatedVideoAndCreatedStatusResponse_WhenPayloadIsOK(t *testing.T) {
	var router = VideoRouter{}
	router.service = &mocked_services.VideoServiceMock{}
	videoModel := mocked_data.GetValidVideo()
	videoDto := mocked_data.GetValidInsertVideoDto()
	videoDtoJson, _ := json.Marshal(videoDto)
	videoModelJson, _ := json.Marshal(videoModel)

	mocked_services.VideoServiceMockCreate = func(dto dto.InsertVideo) (*models.Video, error) {
		return videoModel, nil
	}

	r, _ := http.NewRequest("POST", "/api/v1/videos", bytes.NewReader(videoDtoJson))
	w := httptest.NewRecorder()

	router.CreateVideo(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, videoModelJson, w.Body.Bytes())
}

func TestUpdateVideo_ShouldReturnInvalidPayloadErrorAndBadRequestStatusResponse_WhenPayloadIsInvalid(t *testing.T) {
	var router = VideoRouter{}
	r, _ := http.NewRequest("PUT", "/api/v1/videos/1", bytes.NewReader([]byte("")))
	w := httptest.NewRecorder()

	router.UpdateVideoByID(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Invalid request payload\"}"), w.Body.Bytes())
}

func TestUpdateVideo_ShouldReturnAnErrorAndBadRequestStatusResponse_WhenIsAnInvalidVideo(t *testing.T) {
	var router = VideoRouter{}
	videoDto := mocked_data.GetInvalidInsertVideoDto()
	videoDtoJson, _ := json.Marshal(videoDto)

	r, _ := http.NewRequest("PUT", "/api/v1/videos/1", bytes.NewReader(videoDtoJson))
	w := httptest.NewRecorder()

	router.UpdateVideoByID(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Titulo is required.\"}"), w.Body.Bytes())
}

func TestUpdateVideo_ShouldReturnAnErrorAndInternalServerErrorStatusResponse_WhenTheresAProblemOnVideoService(t *testing.T) {
	var router = VideoRouter{}
	router.service = &mocked_services.VideoServiceMock{}
	videoDto := mocked_data.GetValidInsertVideoDto()
	videoDtoJson, _ := json.Marshal(videoDto)

	r, _ := http.NewRequest("PUT", "/api/v1/videos/1", bytes.NewReader(videoDtoJson))
	w := httptest.NewRecorder()

	mocked_services.VideoServiceMockUpdate = func(id primitive.ObjectID, dto dto.InsertVideo) (*models.Video, error) {
		return nil, errors.New("There's an error")
	}

	router.UpdateVideoByID(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, []byte("{\"error\":\"There's an error\"}"), w.Body.Bytes())
}

func TestUpdateVideo_ShouldReturnOKStatusResponse_WhenPayloadIsOK(t *testing.T) {
	var router = VideoRouter{}
	router.service = &mocked_services.VideoServiceMock{}
	videoModel := mocked_data.GetValidVideo()
	videoDto := mocked_data.GetValidInsertVideoDto()
	videoDtoJson, _ := json.Marshal(videoDto)
	videoModelJson, _ := json.Marshal(videoModel)

	mocked_services.VideoServiceMockUpdate = func(id primitive.ObjectID, dto dto.InsertVideo) (*models.Video, error) {
		return videoModel, nil
	}

	r, _ := http.NewRequest("PUT", "/api/v1/videos/"+videoModel.ID.Hex(), bytes.NewReader(videoDtoJson))
	w := httptest.NewRecorder()

	router.UpdateVideoByID(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, videoModelJson, w.Body.Bytes())
}

func TestDeleteVideo_ShouldReturnAnErrorAndInternalServerErrorStatusResponse_WhenTheresAProblemOnVideoService(t *testing.T) {
	var router = VideoRouter{}
	router.service = &mocked_services.VideoServiceMock{}

	r, _ := http.NewRequest("DELETE", "/api/v1/videos/"+primitive.NewObjectID().Hex(), nil)
	w := httptest.NewRecorder()

	mocked_services.VideoServiceMockDelete = func(id primitive.ObjectID) error {
		return errors.New("There's an error")
	}

	router.DeleteVideoByID(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, []byte("{\"error\":\"There's an error\"}"), w.Body.Bytes())
}

func TestDeleteVideo_ShouldReturnNoContentResponse_WhenTheItemCouldBeDeleted(t *testing.T) {
	var router = VideoRouter{}
	router.service = &mocked_services.VideoServiceMock{}

	r, _ := http.NewRequest("DELETE", "/api/v1/videos/"+primitive.NewObjectID().Hex(), nil)
	w := httptest.NewRecorder()

	mocked_services.VideoServiceMockDelete = func(id primitive.ObjectID) error {
		return nil
	}

	router.DeleteVideoByID(w, r)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Nil(t, w.Body.Bytes())
}
