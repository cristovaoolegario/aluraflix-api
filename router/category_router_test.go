package router

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/cristovaoolegario/aluraflix-api/dto"
	"github.com/cristovaoolegario/aluraflix-api/mocked_data"
	"github.com/cristovaoolegario/aluraflix-api/mocked_services"
	"github.com/cristovaoolegario/aluraflix-api/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllCategories_ShouldReturnCategoryArrayAndOKStatusResponse_WhenTheresNoItemsToShow(t *testing.T) {
	var router = CategoryRouter{}

	router.service = &mocked_services.CategoryServiceMock{}
	categoryArray := []models.Category{*mocked_data.GetValidCategory()}
	categoryArrayJson, _ := json.Marshal(categoryArray)

	mocked_services.CategoryServiceMockGetAll = func(filter string, page int64, pageSize int64) ([]models.Category, error) {
		return categoryArray, nil
	}

	r, _ := http.NewRequest("GET", "/api/v1/category", nil)
	w := httptest.NewRecorder()

	router.GetAllCategories(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, categoryArrayJson, w.Body.Bytes())
}

func TestGetAllCategories_ShouldReturnErrorAndInternalServerErrorStatusResponse_WhenTheresAnError(t *testing.T) {
	var router = CategoryRouter{}
	router.service = &mocked_services.CategoryServiceMock{}

	mocked_services.CategoryServiceMockGetAll = func(filter string, page int64, pageSize int64) ([]models.Category, error) {
		return nil, errors.New("Error test")
	}

	r, _ := http.NewRequest("GET", "/api/v1/category", nil)
	w := httptest.NewRecorder()

	router.GetAllCategories(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Error test\"}"), w.Body.Bytes())
}

func TestGetAllCategories_ShouldReturnEmptyArrayAndNotFoundStatusResponse_WhenTheresAnError(t *testing.T) {
	var router = CategoryRouter{}
	router.service = &mocked_services.CategoryServiceMock{}

	mocked_services.CategoryServiceMockGetAll = func(filter string, page int64, pageSize int64) ([]models.Category, error) {
		return nil, nil
	}

	r, _ := http.NewRequest("GET", "/api/v1/category", nil)
	w := httptest.NewRecorder()

	router.GetAllCategories(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, []byte("[]"), w.Body.Bytes())
}

func TestGetCategoryByID_ShouldReturnEmptyCategoryAnd404StatusResponse_WhenTheresIsNoCategoryWithThatId(t *testing.T) {
	var router = CategoryRouter{}
	router.service = &mocked_services.CategoryServiceMock{}
	id := primitive.NewObjectID().Hex()

	mocked_services.CategoryServiceMockGetByID = func(id primitive.ObjectID) (*models.Category, error) {
		return nil, errors.New("Error test")
	}

	r, _ := http.NewRequest("GET", "/api/v1/categories/"+id, nil)
	w := httptest.NewRecorder()

	router.GetCategoryByID(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Nil(t, w.Body.Bytes())
}

func TestGetCategoryByID_ShouldReturnCategoryInBodyAndOKStatusResponse_WhenTheresItemsToShow(t *testing.T) {
	var router = CategoryRouter{}
	router.service = &mocked_services.CategoryServiceMock{}
	id := primitive.NewObjectID()
	category := mocked_data.GetValidCategoryWithId(id)
	categoryJson, _ := json.Marshal(category)

	mocked_services.CategoryServiceMockGetByID = func(id primitive.ObjectID) (*models.Category, error) {
		return category, nil
	}

	r, _ := http.NewRequest("GET", "/api/v1/categories/"+id.Hex(), nil)
	w := httptest.NewRecorder()

	router.GetCategoryByID(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, categoryJson, w.Body.Bytes())
}

func TestCreateCategory_ShouldReturnInvalidRequestPayloadAndBadRequestStatusResponse_WhenPayloadIsInvalid(t *testing.T) {
	r, _ := http.NewRequest("POST", "/api/v1/categories", bytes.NewReader([]byte("")))
	w := httptest.NewRecorder()

	var router = CategoryRouter{}
	router.CreateCategory(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Invalid request payload\"}"), w.Body.Bytes())
}

func TestCreateCategory_ShouldReturnInvalidDataAndBadRequestStatusResponse_WhenInsertCategoryIsNotValid(t *testing.T) {
	categoryDto := mocked_data.GetInvalidInsertCategoryDto()
	categoryDtoJson, _ := json.Marshal(categoryDto)

	r, _ := http.NewRequest("POST", "/api/v1/categories", bytes.NewReader(categoryDtoJson))
	w := httptest.NewRecorder()
	var router = CategoryRouter{}
	router.CreateCategory(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Titulo is required.\"}"), w.Body.Bytes())
}

func TestCreateCategory_ShouldReturnErrorAndInternalServerErrorStatusResponse_WhenTheresAProblemWithTheCreation(t *testing.T) {
	var router = CategoryRouter{}
	router.service = &mocked_services.CategoryServiceMock{}
	categoryDto := mocked_data.GetValidInsertCategoryDto()
	categoryDtoJson, _ := json.Marshal(categoryDto)

	r, _ := http.NewRequest("POST", "/api/v1/categories", bytes.NewReader(categoryDtoJson))
	w := httptest.NewRecorder()

	mocked_services.CategoryServiceMockCreate = func(insertCategory dto.InsertCategory) (*models.Category, error) {
		return nil, errors.New("There's an error")
	}

	router.CreateCategory(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, []byte("{\"error\":\"There's an error\"}"), w.Body.Bytes())
}

func TestCreateCategory_ShouldReturnCreatedCategoryAndCreatedStatusResponse_WhenPayloadIsOk(t *testing.T) {
	var router = CategoryRouter{}
	router.service = &mocked_services.CategoryServiceMock{}
	categoryModel := mocked_data.GetValidCategory()
	categoryDto := mocked_data.GetValidInsertCategoryDto()
	categoryModelJson, _ := json.Marshal(categoryModel)
	categoryDtoJson, _ := json.Marshal(categoryDto)

	r, _ := http.NewRequest("POST", "/api/v1/categories", bytes.NewReader(categoryDtoJson))
	w := httptest.NewRecorder()

	mocked_services.CategoryServiceMockCreate = func(insertCategory dto.InsertCategory) (*models.Category, error) {
		return categoryModel, nil
	}

	router.CreateCategory(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, categoryModelJson, w.Body.Bytes())
}

func TestUpdateCategoryByID_ShouldReturnInvalidPayloadErrorAndBadRequestStatusResponse_WhenIsAnInvalidPayload(t *testing.T) {
	r, _ := http.NewRequest("PUT", "/api/v1/categories"+primitive.NewObjectID().Hex(), bytes.NewReader([]byte("")))
	w := httptest.NewRecorder()

	var router = CategoryRouter{}
	router.UpdateCategoryByID(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Invalid request payload\"}"), w.Body.Bytes())
}

func TestUpdateCategoryByID_ShouldReturnErrorAndBadRequestStatusResponse_WhenIsAnInvalidCategory(t *testing.T) {
	categoryDto := mocked_data.GetInvalidInsertCategoryDto()
	categoryDtoJson, _ := json.Marshal(categoryDto)

	r, _ := http.NewRequest("PUT", "/api/v1/categories"+primitive.NewObjectID().Hex(), bytes.NewReader(categoryDtoJson))
	w := httptest.NewRecorder()

	var router = CategoryRouter{}
	router.UpdateCategoryByID(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Titulo is required.\"}"), w.Body.Bytes())
}

func TestUpdateCategoryByID_ShouldReturnErrorAndInternalServerErrorStatusResponse_WhenTheresAProblemWithTheService(t *testing.T) {
	var router = CategoryRouter{}
	router.service = &mocked_services.CategoryServiceMock{}
	categoryDto := mocked_data.GetValidCategory()
	categoryDtoJson, _ := json.Marshal(categoryDto)

	r, _ := http.NewRequest("PUT", "/api/v1/categories"+primitive.NewObjectID().Hex(), bytes.NewReader(categoryDtoJson))
	w := httptest.NewRecorder()

	mocked_services.CategoryServiceMockUpdate = func(id primitive.ObjectID, insertCategory dto.InsertCategory) (*models.Category, error) {
		return nil, errors.New("There's an error")
	}

	router.UpdateCategoryByID(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, []byte("{\"error\":\"There's an error\"}"), w.Body.Bytes())
}

func TestUpdateCategoryByID_ShouldReturnOKStatusResponse_WhenPayloadIsOK(t *testing.T) {
	var router = CategoryRouter{}
	router.service = &mocked_services.CategoryServiceMock{}
	categoryModel := mocked_data.GetValidCategory()
	categoryModelJson, _ := json.Marshal(categoryModel)
	categoryDto := mocked_data.GetValidCategory()
	categoryDtoJson, _ := json.Marshal(categoryDto)

	r, _ := http.NewRequest("PUT", "/api/v1/categories"+primitive.NewObjectID().Hex(), bytes.NewReader(categoryDtoJson))
	w := httptest.NewRecorder()

	mocked_services.CategoryServiceMockUpdate = func(id primitive.ObjectID, insertCategory dto.InsertCategory) (*models.Category, error) {
		return categoryModel, nil
	}

	router.UpdateCategoryByID(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, categoryModelJson, w.Body.Bytes())
}

func TestDeleteCategoryByID_ShouldReturnErrorAndInternalServerErrorStatusResponse_WhenTheresAProblemDeletingTheObject(t *testing.T) {
	var router = CategoryRouter{}
	router.service = &mocked_services.CategoryServiceMock{}

	r, _ := http.NewRequest("DELETE", "/api/v1/categories/"+primitive.NewObjectID().Hex(), nil)
	w := httptest.NewRecorder()

	mocked_services.CategoryServiceMockDelete = func(id primitive.ObjectID) error {
		return errors.New("There's an error")
	}

	router.DeleteCategoryByID(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, []byte("{\"error\":\"There's an error\"}"), w.Body.Bytes())
}

func TestDeleteCategoryByID_ShouldReturnNoContentResponse_WhenTheItemCouldBeDeleted(t *testing.T) {
	var router = CategoryRouter{}
	router.service = &mocked_services.CategoryServiceMock{}
	r, _ := http.NewRequest("DELETE", "/api/v1/categories/"+primitive.NewObjectID().Hex(), nil)
	w := httptest.NewRecorder()

	mocked_services.CategoryServiceMockDelete = func(id primitive.ObjectID) error {
		return nil
	}

	router.DeleteCategoryByID(w, r)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Nil(t, w.Body.Bytes())
}

func TestGetAllVideosByCategoryID_ShouldReturnEmptyVideoArrayAndOKStatusResponse_WhenTheresNoItemsToShow(t *testing.T) {
	var router = CategoryRouter{}
	router.service = &mocked_services.CategoryServiceMock{}
	videosArray := []models.Video{*mocked_data.GetValidVideo()}
	videosArrayJson, _ := json.Marshal(videosArray)

	mocked_services.CategoryServiceMockGetVideosByCategoryId = func(id primitive.ObjectID) ([]models.Video, error) {
		return videosArray, nil
	}

	r, _ := http.NewRequest("GET", "/api/v1/category/"+primitive.NewObjectID().Hex()+"/videos", nil)
	w := httptest.NewRecorder()

	router.GetAllVideosByCategoryID(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, videosArrayJson, w.Body.Bytes())
}

func TestGetAllVideosByCategoryID_ShouldReturnEmptyArrayAndNotFoundStatusResponse_WhenTheresNoItemsToShow(t *testing.T) {
	var router = CategoryRouter{}
	router.service = &mocked_services.CategoryServiceMock{}

	mocked_services.CategoryServiceMockGetVideosByCategoryId = func(id primitive.ObjectID) ([]models.Video, error) {
		return nil, nil
	}

	r, _ := http.NewRequest("GET", "/api/v1/category/"+primitive.NewObjectID().Hex()+"/videos", nil)
	w := httptest.NewRecorder()

	router.GetAllVideosByCategoryID(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, []byte("[]"), w.Body.Bytes())
}

func TestGetAllVideosByCategoryID_ShouldReturnErrorAndInternalServerErrorStatusResponse_WhenTheresAnError(t *testing.T) {
	var router = CategoryRouter{}
	router.service = &mocked_services.CategoryServiceMock{}

	mocked_services.CategoryServiceMockGetVideosByCategoryId = func(id primitive.ObjectID) ([]models.Video, error) {
		return nil, errors.New("Error test")
	}

	r, _ := http.NewRequest("GET", "/api/v1/category/"+primitive.NewObjectID().Hex()+"/videos", nil)
	w := httptest.NewRecorder()

	router.GetAllVideosByCategoryID(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Error test\"}"), w.Body.Bytes())
}
