package router

import (
	"encoding/json"
	"github.com/auth0/go-jwt-middleware"
	"github.com/cristovaoolegario/aluraflix-api/auth"
	_ "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"strconv"
)

func Router() *mux.Router {
	r := mux.Router{}
	addVideosResources(&r, auth.JwtMiddleware)
	addCategoriesResources(&r, auth.JwtMiddleware)
	return &r
}

func addVideosResources(r *mux.Router, middleware *jwtmiddleware.JWTMiddleware) {
	r.Handle("/api/v1/videos", middleware.Handler(GetAllVideos)).Methods("GET")
	r.Handle("/api/v1/videos/{id}", middleware.Handler(GetVideoByID)).Methods("GET")
	r.Handle("/api/v1/videos", middleware.Handler(CreateVideo)).Methods("POST")
	r.Handle("/api/v1/videos/{id}", middleware.Handler(UpdateVideoByID)).Methods("PUT")
	r.Handle("/api/v1/videos/{id}", middleware.Handler(DeleteVideoByID)).Methods("DELETE")
}

func addCategoriesResources(r *mux.Router, middleware *jwtmiddleware.JWTMiddleware) {
	r.Handle("/api/v1/categories", middleware.Handler(GetAllCategories)).Methods("GET")
	r.Handle("/api/v1/categories/{id}", middleware.Handler(GetCategoryByID)).Methods("GET")
	r.Handle("/api/v1/categories/{id}/videos", middleware.Handler(GetAllVideosByCategoryID)).Methods("GET")
	r.Handle("/api/v1/categories", middleware.Handler(CreateCategory)).Methods("POST")
	r.Handle("/api/v1/categories/{id}", middleware.Handler(UpdateCategoryByID)).Methods("PUT")
	r.Handle("/api/v1/categories/{id}", middleware.Handler(DeleteCategoryByID)).Methods("DELETE")
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
