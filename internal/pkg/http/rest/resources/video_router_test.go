package resources

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/http/dto"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/storage/bson/db/models"
	"github.com/cristovaoolegario/aluraflix-api/internal/tests/mocked_data"
	"github.com/cristovaoolegario/aluraflix-api/internal/tests/mocked_services"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllFreeVideos(t *testing.T) {
	t.Run("Should return free videos array and ok (200) status response when theres items to show", func(t *testing.T) {
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
	})

	t.Run("Should return empty free videos array and not found (404) status response when theres no items to show", func(t *testing.T) {
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

	})

	t.Run("Should return error and internal server error (500) status response when theres an error", func(t *testing.T) {
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
	})
}

func TestGetAllVideos(t *testing.T) {
	t.Run("Should return videos array and ok (200) status response when theres items to show", func(t *testing.T) {
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
	})

	t.Run("Should return empty videos array and not found (404) status response when theres no items to show", func(t *testing.T) {
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
	})

	t.Run("Should return error and internal server error (500) status response when theres an error", func(t *testing.T) {
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
	})
}

func TestGetVideoByID(t *testing.T) {
	t.Run("Should return empty body and not found (404) status response when theres no items to show", func(t *testing.T) {
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
	})

	t.Run("Should return video in body and ok (200) status response when theres items to show", func(t *testing.T) {
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
	})
}

func TestCreateVideo(t *testing.T) {
	t.Run("Should return invalid payload error and bad request (400) status response when payload is invalid", func(t *testing.T) {
		var router = VideoRouter{}
		r, _ := http.NewRequest("POST", "/api/v1/videos", bytes.NewReader([]byte("")))
		w := httptest.NewRecorder()

		router.CreateVideo(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, []byte("{\"error\":\"Invalid request payload\"}"), w.Body.Bytes())
	})

	t.Run("Should return an error and bad request (400) status response when is an invalid video", func(t *testing.T) {
		var router = VideoRouter{}
		videoDto := mocked_data.GetInvalidInsertVideoDto()
		videoDtoJson, _ := json.Marshal(videoDto)

		r, _ := http.NewRequest("POST", "/api/v1/videos", bytes.NewReader(videoDtoJson))
		w := httptest.NewRecorder()

		router.CreateVideo(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, []byte("{\"error\":\"Titulo is required.\"}"), w.Body.Bytes())
	})

	t.Run("Should return an error and internal server error (500) status response when theres a problem on VideoService", func(t *testing.T) {
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
	})

	t.Run("Should return created video and created (201) status response when payload is ok", func(t *testing.T) {
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
	})
}

func TestUpdateVideo(t *testing.T) {
	t.Run("Should return invalid payload error and bad request (400) status response when payload is invalid", func(t *testing.T) {
		var router = VideoRouter{}
		r, _ := http.NewRequest("PUT", "/api/v1/videos/1", bytes.NewReader([]byte("")))
		w := httptest.NewRecorder()

		router.UpdateVideoByID(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, []byte("{\"error\":\"Invalid request payload\"}"), w.Body.Bytes())
	})

	t.Run("Should return an error and bad request (400) status response when is an invalid video", func(t *testing.T) {
		var router = VideoRouter{}
		videoDto := mocked_data.GetInvalidInsertVideoDto()
		videoDtoJson, _ := json.Marshal(videoDto)

		r, _ := http.NewRequest("PUT", "/api/v1/videos/1", bytes.NewReader(videoDtoJson))
		w := httptest.NewRecorder()

		router.UpdateVideoByID(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, []byte("{\"error\":\"Titulo is required.\"}"), w.Body.Bytes())
	})

	t.Run("Should return an error and internal server error (500) status response when theres a problem on VideoService", func(t *testing.T) {
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
	})

	t.Run("Should return ok (200) status response when payload is ok", func(t *testing.T) {
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
	})
}

func TestDeleteVideo(t *testing.T) {
	t.Run("Should return an error and internal server error (500) status response when theres a problem on videoservice", func(t *testing.T) {
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
	})

	t.Run("Should return no content (204) response when the item could be deleted", func(t *testing.T) {
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
	})
}
