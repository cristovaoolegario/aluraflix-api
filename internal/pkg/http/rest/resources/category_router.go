package resources

import (
	"encoding/json"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/http/dto"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/interfaces"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/storage/bson/db/models"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/storage/bson/db/services"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)


type CategoryRouter struct {
	service interfaces.ICategoryService
}

func ProvideCategoryRouter(s services.CategoryService) CategoryRouter {
	return CategoryRouter{&s}
}

func (cs *CategoryRouter) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	filter, page, pageSize := GetQueryParams(r.URL.Query())
	categories, err := cs.service.GetAll(filter, page, pageSize)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if categories == nil {
		RespondWithJson(w, http.StatusNotFound, []models.Video{})
		return
	}
	RespondWithJson(w, http.StatusOK, categories)
}

func (cs *CategoryRouter) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	category, err := cs.service.GetById(id)
	if err != nil {
		RespondWithJson(w, http.StatusNotFound, nil)
		return
	}
	RespondWithJson(w, http.StatusOK, category)
}

func (cs *CategoryRouter) CreateCategory(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var category dto.InsertCategory
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err = category.Validate(); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	insertedVideo, err := cs.service.Create(category)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusCreated, insertedVideo)
}

func (cs *CategoryRouter) UpdateCategoryByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var category dto.InsertCategory
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := category.Validate(); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	id, _ := primitive.ObjectIDFromHex(params["id"])
	updatedCategory, err := cs.service.Update(id, category)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, updatedCategory)
}

func (cs *CategoryRouter) DeleteCategoryByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	if err := cs.service.Delete(id); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusNoContent, nil)
}

func (cs *CategoryRouter) GetAllVideosByCategoryID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	videos, err := cs.service.GetVideosByCategoryId(id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if videos == nil {
		RespondWithJson(w, http.StatusNotFound, []models.Video{})
		return
	}
	RespondWithJson(w, http.StatusOK, videos)
}
