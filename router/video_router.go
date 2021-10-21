package router

import (
	"encoding/json"
	"github.com/cristovaoolegario/aluraflix-api/db"
	"github.com/cristovaoolegario/aluraflix-api/dto"
	"github.com/cristovaoolegario/aluraflix-api/interfaces"
	"github.com/cristovaoolegario/aluraflix-api/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type VideoRouter struct {
	service interfaces.IVideoService
}

func ProvideVideoRouter(s db.VideoService) VideoRouter {
	return VideoRouter{&s}
}

func (vr *VideoRouter) GetAllFreeVideos(w http.ResponseWriter, r *http.Request) {
	videos, err := vr.service.GetAllFreeVideos()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if videos == nil {
		respondWithJson(w, http.StatusNotFound, []models.Video{})
		return
	}
	respondWithJson(w, http.StatusOK, videos)
}

func (vr *VideoRouter) GetAllVideos(w http.ResponseWriter, r *http.Request) {
	filter, page, pageSize := getQueryParams(r.URL.Query())
	videos, err := vr.service.GetAll(filter, page, pageSize)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if videos == nil {
		respondWithJson(w, http.StatusNotFound, []models.Video{})
		return
	}
	respondWithJson(w, http.StatusOK, videos)
}

func (vr *VideoRouter) GetVideoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	video, err := vr.service.GetByID(id)
	if err != nil {
		respondWithJson(w, http.StatusNotFound, nil)
		return
	}
	respondWithJson(w, http.StatusOK, video)
}

func (vr *VideoRouter) CreateVideo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var video dto.InsertVideo
	err := json.NewDecoder(r.Body).Decode(&video)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err = video.Validate(); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	createdVideo, err := vr.service.Create(video)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, createdVideo)
}

func (vr *VideoRouter) UpdateVideoByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var video dto.InsertVideo
	if err := json.NewDecoder(r.Body).Decode(&video); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := video.Validate(); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	id, _ := primitive.ObjectIDFromHex(params["id"])
	updatedVideo, err := vr.service.Update(id, video)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, updatedVideo)
}

func (vr *VideoRouter) DeleteVideoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	if err := vr.service.Delete(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusNoContent, nil)
}
