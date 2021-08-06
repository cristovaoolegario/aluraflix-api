package router

import (
	"github.com/cristovaoolegario/aluraflix-api/db"
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
