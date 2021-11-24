package resources

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/http/dto"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/storage/bson/db/models"
	"github.com/cristovaoolegario/aluraflix-api/internal/tests/mocked_data"
	"github.com/cristovaoolegario/aluraflix-api/internal/tests/mocked_services"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllCategories(t *testing.T) {
	t.Run("Should return category array and OK (200) status response When theres no items to show", func(t *testing.T) {
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
	})

	t.Run("Should return error and internal server error (500) status response When theres an error", func(t *testing.T) {
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
	})

	t.Run("Should return empty array and not found (404) status response When theres and error", func(t *testing.T) {
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
	})
}

func TestGetCategoryByID(t *testing.T) {
	t.Run("Should return empty category and not found (404) status response when theres is no category with that id", func(t *testing.T) {
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
	})

	t.Run("Should return category in body and ok (200) status response when theres items to show", func(t *testing.T) {
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
	})
}

func TestCreateCategory(t *testing.T) {
	t.Run("Should return invalid request payload and bad request (400) status response when payload is invalid", func(t *testing.T) {
		r, _ := http.NewRequest("POST", "/api/v1/categories", bytes.NewReader([]byte("")))
		w := httptest.NewRecorder()

		var router = CategoryRouter{}
		router.CreateCategory(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, []byte("{\"error\":\"Invalid request payload\"}"), w.Body.Bytes())
	})

	t.Run("Should return invalid data and bad request (400) status response when insert category is not valid", func(t *testing.T) {
		categoryDto := mocked_data.GetInvalidInsertCategoryDto()
		categoryDtoJson, _ := json.Marshal(categoryDto)

		r, _ := http.NewRequest("POST", "/api/v1/categories", bytes.NewReader(categoryDtoJson))
		w := httptest.NewRecorder()
		var router = CategoryRouter{}
		router.CreateCategory(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, []byte("{\"error\":\"Titulo is required.\"}"), w.Body.Bytes())
	})

	t.Run("should return error and internal server (500) error status response when theres a problem with the creation", func(t *testing.T) {
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
	})

	t.Run("Should return created category and created (201) status response when payload is ok", func(t *testing.T) {
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
	})
}

func TestUpdateCategoryByID(t *testing.T) {
	t.Run("Should return invalid payload error and bad request (400) status response when is an invalid payload", func(t *testing.T) {
		r, _ := http.NewRequest("PUT", "/api/v1/categories"+primitive.NewObjectID().Hex(), bytes.NewReader([]byte("")))
		w := httptest.NewRecorder()

		var router = CategoryRouter{}
		router.UpdateCategoryByID(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, []byte("{\"error\":\"Invalid request payload\"}"), w.Body.Bytes())
	})

	t.Run("Should return error and bad request (400) status response when is an invalid category", func(t *testing.T) {
		categoryDto := mocked_data.GetInvalidInsertCategoryDto()
		categoryDtoJson, _ := json.Marshal(categoryDto)

		r, _ := http.NewRequest("PUT", "/api/v1/categories"+primitive.NewObjectID().Hex(), bytes.NewReader(categoryDtoJson))
		w := httptest.NewRecorder()

		var router = CategoryRouter{}
		router.UpdateCategoryByID(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, []byte("{\"error\":\"Titulo is required.\"}"), w.Body.Bytes())
	})

	t.Run("Should return error and internal server error (500) status response when theres a problem with the service", func(t *testing.T) {
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
	})

	t.Run("Should return ok (200) status response when payload is ok", func(t *testing.T) {
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
	})
}

func TestDeleteCategoryByID(t *testing.T) {

	t.Run("Should return error and internal server error (500) status response when theres a problem deleting the object", func(t *testing.T) {
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

	})

	t.Run("Should return no content (204) status response when the item could be deleted", func(t *testing.T) {
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
	})
}

func TestGetAllVideosByCategoryID(t *testing.T) {
	t.Run("Should return empty video array and ok (200) status response when theres no items to show", func(t *testing.T) {
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
	})

	t.Run("Should return empty array and not found (404) status response when theres no items to show", func(t *testing.T) {
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
	})

	t.Run("Should return error and internal server error (500) status response when theres an error", func(t *testing.T) {
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

	})
}
