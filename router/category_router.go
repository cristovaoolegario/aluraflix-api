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

var categoryService interfaces.ICategoryService

func init() {
	categoryService = &db.CategoryService{}
}

var GetAllCategories = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	filter, page, pageSize := getQueryParams(r.URL.Query())
	categories, err := categoryService.GetAll(filter, page, pageSize)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if categories == nil {
		respondWithJson(w, http.StatusNotFound, []models.Video{})
		return
	}
	respondWithJson(w, http.StatusOK, categories)
})

var GetCategoryByID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	category, err := categoryService.GetById(id)
	if err != nil {
		respondWithJson(w, http.StatusNotFound, nil)
		return
	}
	respondWithJson(w, http.StatusOK, category)
})

var CreateCategory = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var category dto.InsertCategory
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err = category.Validate(); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	insertedVideo, err := categoryService.Create(category)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, insertedVideo)
})

var UpdateCategoryByID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var category dto.InsertCategory
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := category.Validate(); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	id, _ := primitive.ObjectIDFromHex(params["id"])
	updatedCategory, err := categoryService.Update(id, category)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, updatedCategory)
})

var DeleteCategoryByID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	if err := categoryService.Delete(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusNoContent, nil)
})

var GetAllVideosByCategoryID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	videos, err := categoryService.GetVideosByCategoryId(id)
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
