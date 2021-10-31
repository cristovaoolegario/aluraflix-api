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

// GetAllCategories godoc
// @Summary Get details of all categories
// @Description Get details of all categories
// @Tags categories
// @Accept  json
// @Produce  json
// @Param search query string false "Search by name"
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Security ApiKeyAuth
// @Success 200 {array} models.Category
// @Failure 400 {object} ErrorMessage
// @Failure 401 {string} string
// @Failure 404
// @Failure 500 {object} ErrorMessage
// @Router /categories [get]
func (cs *CategoryRouter) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	filter, page, pageSize := GetQueryParams(r.URL.Query())
	categories, err := cs.service.GetAll(filter, page, pageSize)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if categories == nil {
		RespondWithJson(w, http.StatusNotFound, []models.Category{})
		return
	}
	RespondWithJson(w, http.StatusOK, categories)
}

// GetCategoryByID godoc
// @Summary Get details of a category by ID
// @Description Get details of a category by ID
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Security ApiKeyAuth
// @Success 200 {object} models.Category
// @Failure 400 {object} ErrorMessage
// @Failure 401 {string} string 
// @Failure 404
// @Failure 500 {object} ErrorMessage
// @Router /categories/{id} [get]
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

// CreateCategory godoc
// @Summary Create a new Category
// @Description Create a new Category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param category body dto.InsertCategory true "New category"
// @Security ApiKeyAuth
// @Success 201 {object} models.Category
// @Failure 400 {object} ErrorMessage
// @Failure 401 {string} string 
// @Failure 500 {object} ErrorMessage
// @Router /categories [post]
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

// UpdateCategoryByID godoc
// @Summary Update a category by ID
// @Description Update a category by ID
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Param category body dto.InsertCategory true "New category values"
// @Security ApiKeyAuth
// @Success 200 {object} models.Category
// @Failure 400 {object} ErrorMessage
// @Failure 401 {string} string
// @Failure 404
// @Failure 500 {object} ErrorMessage
// @Router /categories [put]
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

// DeleteCategoryByID godoc
// @Summary Delete a category by ID
// @Description Delete a category by ID
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Security ApiKeyAuth
// @Success 200
// @Failure 400 {object} ErrorMessage
// @Failure 401 {string} string
// @Failure 404
// @Failure 500 {object} ErrorMessage
// @Router /categories [delete]
func (cs *CategoryRouter) DeleteCategoryByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	if err := cs.service.Delete(id); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusNoContent, nil)
}

// GetAllVideosByCategoryID godoc
// @Summary Get all videos by category ID
// @Description Get all videos by category ID
// @Tags videos
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Security ApiKeyAuth
// @Success 200 {array} models.Video
// @Failure 400 {object} ErrorMessage
// @Failure 401 {string} string
// @Failure 404
// @Failure 500 {object} ErrorMessage
// @Router /categories/{id}/videos [get]
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
