package resources

import (
	"encoding/json"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/http/dto"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/interfaces"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/storage/bson/db/models"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/storage/bson/db/services"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type VideoRouter struct {
	service interfaces.IVideoService
}

func ProvideVideoRouter(s services.VideoService) VideoRouter {
	return VideoRouter{&s}
}

// GetAllFreeVideos godoc
// @Summary Get all free videos
// @Description Get all free videos
// @Tags videos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Video
// @Failure 400 {object} ErrorMessage
// @Failure 401 {string} string
// @Failure 404
// @Failure 500 {object} ErrorMessage
// @Router /videos/free [get]
func (vr *VideoRouter) GetAllFreeVideos(w http.ResponseWriter, r *http.Request) {
	videos, err := vr.service.GetAllFreeVideos()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if videos == nil {
		RespondWithJson(w, http.StatusNotFound, []models.Video{})
		return
	}
	RespondWithJson(w, http.StatusOK, videos)
}

// GetAllVideos godoc
// @Summary Get details of all videos
// @Description Get details of all videos
// @Tags videos
// @Accept  json
// @Produce  json
// @Param search query string false "Search by name"
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Security ApiKeyAuth
// @Success 200 {array} models.Video
// @Failure 400 {object} ErrorMessage
// @Failure 401 {string} string
// @Failure 404
// @Failure 500 {object} ErrorMessage
// @Router /videos [get]
func (vr *VideoRouter) GetAllVideos(w http.ResponseWriter, r *http.Request) {
	filter, page, pageSize := GetQueryParams(r.URL.Query())
	videos, err := vr.service.GetAll(filter, page, pageSize)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if videos == nil {
		RespondWithJson(w, http.StatusNotFound, []models.Video{})
		return
	}
	RespondWithJson(w, http.StatusOK, videos)
}

// GetVideoByID godoc
// @Summary Get details of a video by ID
// @Description Get details of a video by ID
// @Tags videos
// @Accept  json
// @Produce  json
// @Param id path int true "Video ID"
// @Security ApiKeyAuth
// @Success 200 {object} models.Video
// @Failure 400 {object} ErrorMessage
// @Failure 401 {string} string
// @Failure 404
// @Failure 500 {object} ErrorMessage
// @Router /videos/{id} [get]
func (vr *VideoRouter) GetVideoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	video, err := vr.service.GetByID(id)
	if err != nil {
		RespondWithJson(w, http.StatusNotFound, nil)
		return
	}
	RespondWithJson(w, http.StatusOK, video)
}

// CreateVideo godoc
// @Summary Create a new Video
// @Description Create a new Video
// @Tags videos
// @Accept  json
// @Produce  json
// @Param video body dto.InsertVideo true "New video"
// @Security ApiKeyAuth
// @Success 201 {object} models.Video
// @Failure 400 {object} ErrorMessage
// @Failure 401 {string} string
// @Failure 500 {object} ErrorMessage
// @Router /videos [post]
func (vr *VideoRouter) CreateVideo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var video dto.InsertVideo
	err := json.NewDecoder(r.Body).Decode(&video)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err = video.Validate(); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	createdVideo, err := vr.service.Create(video)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusCreated, createdVideo)
}

// UpdateVideoByID godoc
// @Summary Update a video by ID
// @Description Update a video by ID
// @Tags videos
// @Accept  json
// @Produce  json
// @Param id path int true "Video ID"
// @Param category body dto.InsertVideo true "New video values"
// @Security ApiKeyAuth
// @Success 200 {object} models.Video
// @Failure 400 {object} ErrorMessage
// @Failure 401 {string} string
// @Failure 404
// @Failure 500 {object} ErrorMessage
// @Router /videos [put]
func (vr *VideoRouter) UpdateVideoByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var video dto.InsertVideo
	if err := json.NewDecoder(r.Body).Decode(&video); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := video.Validate(); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	id, _ := primitive.ObjectIDFromHex(params["id"])
	updatedVideo, err := vr.service.Update(id, video)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, updatedVideo)
}

// DeleteVideoByID godoc
// @Summary Delete a video by ID
// @Description Delete a video by ID
// @Tags videos
// @Accept  json
// @Produce  json
// @Param id path int true "Video ID"
// @Security ApiKeyAuth
// @Success 200
// @Failure 400 {object} ErrorMessage
// @Failure 401 {string} string
// @Failure 404
// @Failure 500 {object} ErrorMessage
// @Router /videos [delete]
func (vr *VideoRouter) DeleteVideoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	if err := vr.service.Delete(id); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusNoContent, nil)
}
