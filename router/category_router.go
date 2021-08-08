package router

import (
	"github.com/cristovaoolegario/aluraflix-api/db"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

var categoryService db.ICategoryService

func init() {
	categoryService = &db.CategoryService{}
}

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := categoryService.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, categories)
}

func GetCategoryByID(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	category, err := categoryService.GetById(id)
	if err != nil {
		respondWithJson(w, http.StatusNotFound, nil)
		return
	}
	respondWithJson(w, http.StatusOK, category)
}