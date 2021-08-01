package router

import (
	"encoding/json"
	"github.com/cristovaoolegario/aluraflix-api/db"
	"net/http"
)

var videoService db.IVideoService

func init(){
	videoService = &db.VideoService{}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if payload != nil {
		response, _ := json.Marshal(payload)
		w.Write(response)
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	videos, err := videoService.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, videos)
}