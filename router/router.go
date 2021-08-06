package router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router{
	r := mux.Router{}
	r.HandleFunc("/api/v1/videos", GetAllVideos).Methods("GET")
	r.HandleFunc("/api/v1/videos/{id}", GetVideoByID).Methods("GET")
	r.HandleFunc("/api/v1/videos", CreateVideo).Methods("POST")
	r.HandleFunc("/api/v1/videos/{id}", UpdateVideoByID).Methods("PUT")
	r.HandleFunc("/api/v1/videos/{id}", DeleteVideoByID).Methods("DELETE")
	r.HandleFunc("/api/v1/category", GetAllCategories).Methods("GET")
	return &r
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