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

var videoService interfaces.IVideoService

func init() {
	videoService = &db.VideoService{}
}

var GetAllVideos = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	filter, page, pageSize := getQueryParams(r.URL.Query())
	videos, err := videoService.GetAll(filter, page, pageSize)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if videos == nil {
		respondWithJson(w, http.StatusNotFound, []models.Video{})
		return
	}
	respondWithJson(w, http.StatusOK, videos)
})

var GetVideoByID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	video, err := videoService.GetByID(id)
	if err != nil {
		respondWithJson(w, http.StatusNotFound, nil)
		return
	}
	respondWithJson(w, http.StatusOK, video)
})

var CreateVideo = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	createdVideo, err := videoService.Create(video)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, createdVideo)
})

var UpdateVideoByID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	updatedVideo, err := videoService.Update(id, video)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, updatedVideo)
})

var DeleteVideoByID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	if err := videoService.Delete(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusNoContent, nil)
})
