package rest

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/http/auth/jwt"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/http/rest/resources"
	"github.com/gorilla/mux"
	"net/http"
)

func ProvideRouter(videoRouter resources.VideoRouter, categoryRouter resources.CategoryRouter) mux.Router {
	r := mux.Router{}
	addVideosResources(videoRouter, &r, jwt.JwtMiddleware)
	addCategoriesResources(categoryRouter, &r, jwt.JwtMiddleware)
	return r
}

func addVideosResources(videoRouter resources.VideoRouter, r *mux.Router, middleware *jwtmiddleware.JWTMiddleware) {
	r.HandleFunc("/api/v1/videos/free", videoRouter.GetAllFreeVideos).Methods("GET")
	r.Handle("/api/v1/videos", middleware.Handler(http.HandlerFunc(videoRouter.GetAllVideos))).Methods("GET")
	r.Handle("/api/v1/videos/{id}", middleware.Handler(http.HandlerFunc(videoRouter.GetVideoByID))).Methods("GET")
	r.Handle("/api/v1/videos", middleware.Handler(http.HandlerFunc(videoRouter.CreateVideo))).Methods("POST")
	r.Handle("/api/v1/videos/{id}", middleware.Handler(http.HandlerFunc(videoRouter.UpdateVideoByID))).Methods("PUT")
	r.Handle("/api/v1/videos/{id}", middleware.Handler(http.HandlerFunc(videoRouter.DeleteVideoByID))).Methods("DELETE")
}

func addCategoriesResources(categoryRouter resources.CategoryRouter, r *mux.Router, middleware *jwtmiddleware.JWTMiddleware) {
	r.Handle("/api/v1/categories", middleware.Handler(http.HandlerFunc(categoryRouter.GetAllCategories))).Methods("GET")
	r.Handle("/api/v1/categories/{id}", middleware.Handler(http.HandlerFunc(categoryRouter.GetCategoryByID))).Methods("GET")
	r.Handle("/api/v1/categories/{id}/videos", middleware.Handler(http.HandlerFunc(categoryRouter.GetAllVideosByCategoryID))).Methods("GET")
	r.Handle("/api/v1/categories", middleware.Handler(http.HandlerFunc(categoryRouter.CreateCategory))).Methods("POST")
	r.Handle("/api/v1/categories/{id}", middleware.Handler(http.HandlerFunc(categoryRouter.UpdateCategoryByID))).Methods("PUT")
	r.Handle("/api/v1/categories/{id}", middleware.Handler(http.HandlerFunc(categoryRouter.DeleteCategoryByID))).Methods("DELETE")
}
