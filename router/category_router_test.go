package router

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/cristovaoolegario/aluraflix-api/dto"
	"github.com/cristovaoolegario/aluraflix-api/mocked_data"
	"github.com/cristovaoolegario/aluraflix-api/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllCategories_ShouldReturnEmptyCategoryArrayAndOKStatusResponse_WhenTheresNoItemsToShow(t *testing.T) {

	categoryService = &CategoryServiceMock{}

	categoryServiceMockGetAll = func() ([]models.Category, error) {
		return []models.Category{}, nil
	}

	r, _ := http.NewRequest("GET", "/api/v1/category", nil)
	w := httptest.NewRecorder()

	GetAllCategories(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, []byte("[]"), w.Body.Bytes())
}

func TestGetAllCategories_ShouldReturnErrorAndInternalServerErrorStatusResponse_WhenTheresAnError(t *testing.T) {

	categoryService = &CategoryServiceMock{}

	categoryServiceMockGetAll = func() ([]models.Category, error) {
		return []models.Category{}, errors.New("Error test")
	}

	r, _ := http.NewRequest("GET", "/api/v1/category", nil)
	w := httptest.NewRecorder()

	GetAllCategories(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Error test\"}"), w.Body.Bytes())
}

func TestGetCategoryByID_ShouldReturnEmptyCategoryAnd404StatusResponse_WhenTheresIsNoCategoryWithThatId(t *testing.T) {
	categoryService = &CategoryServiceMock{}
	id := primitive.NewObjectID().Hex()

	categoryServiceMockGetByID = func(id primitive.ObjectID) (*models.Category, error) {
		return nil, errors.New("Error test")
	}

	r, _ := http.NewRequest("GET", "/api/v1/categories/"+id, nil)
	w := httptest.NewRecorder()

	GetCategoryByID(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Nil(t, w.Body.Bytes())
}

func TestGetCategoryByID_ShouldReturnCategoryInBodyAndOKStatusResponse_WhenTheresItemsToShow(t *testing.T) {
	categoryService = &CategoryServiceMock{}
	id := primitive.NewObjectID()
	category := mocked_data.GetValidCategoryWithId(id)
	categoryJson, _ := json.Marshal(category)

	categoryServiceMockGetByID = func(id primitive.ObjectID) (*models.Category, error) {
		return category, nil
	}

	r, _ := http.NewRequest("GET", "/api/v1/categories/"+id.Hex(), nil)
	w := httptest.NewRecorder()

	GetCategoryByID(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, categoryJson, w.Body.Bytes())
}

func TestCreateCategory_ShouldReturnInvalidRequestPayloadAndBadRequestStatusResponse_WhenPayloadIsInvalid(t *testing.T) {
	r, _ := http.NewRequest("POST", "/api/v1/categories", bytes.NewReader([]byte("")))
	w := httptest.NewRecorder()

	CreateCategory(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Invalid request payload\"}"), w.Body.Bytes())
}

func TestCreateCategory_ShouldReturnInvalidDataAndBadRequestStatusResponse_WhenInsertCategoryIsNotValid(t *testing.T) {
	categoryDto := mocked_data.GetInvalidInsertCategoryDto()
	categoryDtoJson, _ := json.Marshal(categoryDto)

	r, _ := http.NewRequest("POST", "/api/v1/categories", bytes.NewReader(categoryDtoJson))
	w := httptest.NewRecorder()

	CreateCategory(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Titulo is required.\"}"), w.Body.Bytes())
}

func TestCreateCategory_ShouldReturnErrorAndInternalServerErrorStatusResponse_WhenTheresAProblemWithTheCreation(t *testing.T) {
	categoryService = &CategoryServiceMock{}
	categoryDto := mocked_data.GetValidInsertCategoryDto()
	categoryDtoJson, _ := json.Marshal(categoryDto)

	r, _ := http.NewRequest("POST", "/api/v1/categories", bytes.NewReader(categoryDtoJson))
	w := httptest.NewRecorder()

	categoryServiceMockCreate = func(insertCategory dto.InsertCategory) (*models.Category, error) {
		return nil, errors.New("There's an error")
	}

	CreateCategory(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, []byte("{\"error\":\"There's an error\"}"), w.Body.Bytes())
}

func TestCreateCategory_ShouldReturnCreatedCategoryAndCreatedStatusResponse_WhenPayloadIsOk(t *testing.T) {
	categoryService = &CategoryServiceMock{}
	categoryModel := mocked_data.GetValidCategory()
	categoryDto := mocked_data.GetValidInsertCategoryDto()
	categoryModelJson, _ := json.Marshal(categoryModel)
	categoryDtoJson, _ := json.Marshal(categoryDto)

	r, _ := http.NewRequest("POST", "/api/v1/categories", bytes.NewReader(categoryDtoJson))
	w := httptest.NewRecorder()

	categoryServiceMockCreate = func(insertCategory dto.InsertCategory) (*models.Category, error) {
		return categoryModel, nil
	}

	CreateCategory(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, categoryModelJson, w.Body.Bytes())
}

func TestUpdateCategoryByID_ShouldReturnInvalidPayloadErrorAndBadRequestStatusResponse_WhenIsAnInvalidPayload(t *testing.T) {
	r, _ := http.NewRequest("PUT", "/api/v1/categories"+ primitive.NewObjectID().Hex(), bytes.NewReader([]byte("")))
	w := httptest.NewRecorder()

	UpdateCategoryByID(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Invalid request payload\"}"), w.Body.Bytes())
}

func TestUpdateCategoryByID_ShouldReturnErrorAndBadRequestStatusResponse_WhenIsAnInvalidCategory(t *testing.T) {
	categoryDto := mocked_data.GetInvalidInsertCategoryDto()
	categoryDtoJson, _ := json.Marshal(categoryDto)

	r, _ := http.NewRequest("PUT", "/api/v1/categories"+ primitive.NewObjectID().Hex(), bytes.NewReader(categoryDtoJson))
	w := httptest.NewRecorder()

	UpdateCategoryByID(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, []byte("{\"error\":\"Titulo is required.\"}"), w.Body.Bytes())
}

func TestUpdateCategoryByID_ShouldReturnErrorAndInternalServerErrorStatusResponse_WhenTheresAProblemWithTheService(t *testing.T) {
	categoryService = &CategoryServiceMock{}
	categoryDto := mocked_data.GetValidCategory()
	categoryDtoJson, _ := json.Marshal(categoryDto)

	r, _ := http.NewRequest("PUT", "/api/v1/categories"+ primitive.NewObjectID().Hex(), bytes.NewReader(categoryDtoJson))
	w := httptest.NewRecorder()

	categoryServiceMockUpdate = func(id primitive.ObjectID, insertCategory dto.InsertCategory) (*models.Category, error){
		return nil, errors.New("There's an error")
	}

	UpdateCategoryByID(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, []byte("{\"error\":\"There's an error\"}"), w.Body.Bytes())
}

func TestUpdateCategoryByID_ShouldReturnOKStatusResponse_WhenPayloadIsOK(t *testing.T) {
	categoryService = &CategoryServiceMock{}
	categoryModel := mocked_data.GetValidCategory()
	categoryModelJson, _ := json.Marshal(categoryModel)
	categoryDto := mocked_data.GetValidCategory()
	categoryDtoJson, _ := json.Marshal(categoryDto)

	r, _ := http.NewRequest("PUT", "/api/v1/categories"+ primitive.NewObjectID().Hex(), bytes.NewReader(categoryDtoJson))
	w := httptest.NewRecorder()

	categoryServiceMockUpdate = func(id primitive.ObjectID, insertCategory dto.InsertCategory) (*models.Category, error){
		return categoryModel, nil
	}

	UpdateCategoryByID(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, categoryModelJson, w.Body.Bytes())
}

func TestDeleteCategoryByID_ShouldReturnErrorAndInternalServerErrorStatusResponse_WhenTheresAProblemDeletingTheObject(t *testing.T) {
	categoryService = &CategoryServiceMock{}

	r, _ := http.NewRequest("DELETE", "/api/v1/categories/" + primitive.NewObjectID().Hex(), nil)
	w := httptest.NewRecorder()

	categoryServiceMockDelete = func(id primitive.ObjectID) error {
		return errors.New("There's an error")
	}

	DeleteCategoryByID(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, []byte("{\"error\":\"There's an error\"}"), w.Body.Bytes())
}

func TestDeleteCategoryByID_ShouldReturnNoContentResponse_WhenTheItemCouldBeDeleted(t *testing.T) {
	categoryService = &CategoryServiceMock{}
	r, _ := http.NewRequest("DELETE", "/api/v1/categories/" + primitive.NewObjectID().Hex(), nil)
	w := httptest.NewRecorder()

	categoryServiceMockDelete = func(id primitive.ObjectID) error {
		return nil
	}

	DeleteCategoryByID(w, r)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Nil(t, w.Body.Bytes())

}