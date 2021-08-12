package router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"strconv"
)

func Router() *mux.Router {
	r := mux.Router{}
	addVideosResources(&r)
	addCategoriesResources(&r)
	return &r
}

func addVideosResources(r *mux.Router) {
	r.HandleFunc("/api/v1/videos", GetAllVideos).Methods("GET")
	r.HandleFunc("/api/v1/videos", GetAllVideos).Queries("search", "{search}").Methods("GET")
	r.HandleFunc("/api/v1/videos/{id}", GetVideoByID).Methods("GET")
	r.HandleFunc("/api/v1/videos", CreateVideo).Methods("POST")
	r.HandleFunc("/api/v1/videos/{id}", UpdateVideoByID).Methods("PUT")
	r.HandleFunc("/api/v1/videos/{id}", DeleteVideoByID).Methods("DELETE")
}

func addCategoriesResources(r *mux.Router) {
	r.HandleFunc("/api/v1/categories", GetAllCategories).Methods("GET")
	r.HandleFunc("/api/v1/categories/{id}", GetCategoryByID).Methods("GET")
	r.HandleFunc("/api/v1/categories/{id}/videos", GetAllVideosByCategoryID).Methods("GET")
	r.HandleFunc("/api/v1/categories", CreateCategory).Methods("POST")
	r.HandleFunc("/api/v1/categories/{id}", UpdateCategoryByID).Methods("PUT")
	r.HandleFunc("/api/v1/categories/{id}", DeleteCategoryByID).Methods("DELETE")
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

func getQueryParams(queryParams url.Values) (filter string, page int64, pageSize int64){
	filter = queryParams.Get("search")
	page = 1
	if n, err := strconv.Atoi(queryParams.Get("page")); err == nil{
		page = int64(n)
	}
	pageSize = 5
	if n, err := strconv.Atoi(queryParams.Get("pageSize")); err == nil{
		pageSize = int64(n)
	}
	return filter, page, pageSize
}
