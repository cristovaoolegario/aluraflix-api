package db

import (
	"github.com/cristovaoolegario/aluraflix-api/dto"
	"github.com/cristovaoolegario/aluraflix-api/mocked_data"
	"github.com/cristovaoolegario/aluraflix-api/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestCategoryService(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("GetAllCategories method Should return object when has objects", func(mt *mtest.T) {
		categoriesCollection = mt.Coll
		firstId := primitive.NewObjectID()
		secondId := primitive.NewObjectID()

		firstCategory := mtest.CreateCursorResponse(1,"foo.bar", mtest.FirstBatch, mocked_data.GetBsonFromCategory(mocked_data.GetValidCategoryWithId(firstId)))

		secondCategory := mtest.CreateCursorResponse(1,"foo.bar", mtest.NextBatch, mocked_data.GetBsonFromCategory(mocked_data.GetValidCategoryWithId(secondId)))

		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(firstCategory, secondCategory, killCursors)

		var categoryService = CategoryService{}

		response, err := categoryService.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(response))
		mt.ClearMockResponses()
	})

	mt.Run("GetAllCategories method Should return error when dont has objects", func(mt *mtest.T) {
		categoriesCollection = mt.Coll

		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(bson.D{}, killCursors)

		var categoryService = CategoryService{}

		response, err := categoryService.GetAll()
		assert.NotNil(t, err)
		assert.Equal(t, 0, len(response))
		mt.ClearMockResponses()
	})

	mt.Run("GetCategoryById method Should return error when object dont exists", func(mt *mtest.T) {
		categoriesCollection = mt.Coll
		id := primitive.NewObjectID()

		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(bson.D{}, killCursors)

		var categoryService = CategoryService{}

		response, err := categoryService.GetById(id)
		assert.NotNil(t, err)
		assert.Nil(t, response)
		mt.ClearMockResponses()
	})

	mt.Run("GetCategoryById method Should return object when object exists", func(mt *mtest.T) {
		categoriesCollection = mt.Coll
		id := primitive.NewObjectID()
		expectedCategory := mocked_data.GetValidCategoryWithId(id)

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, mocked_data.GetBsonFromCategory(expectedCategory)))

		var categoryService = CategoryService{}

		response, err := categoryService.GetById(id)
		assert.Nil(t, err)
		assert.Equal(t, expectedCategory, response)
		mt.ClearMockResponses()
	})

	mt.Run("CreateCategory method Should return error when object could not be created", func(mt *mtest.T) {
		categoriesCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "Con't insert data",
		}))

		var categoryService = CategoryService{}

		response, err := categoryService.Create(dto.InsertCategory{})
		assert.Nil(t, response)
		assert.NotNil(t, err)
		mt.ClearMockResponses()
	})

	mt.Run("CreateCategory method Should return error when object could not be created", func(mt *mtest.T) {
		categoriesCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		var categoryService = CategoryService{}

		response, err := categoryService.Create(mocked_data.GetValidInsertCategoryDto())

		assert.NotNil(t, response)
		assert.Nil(t, err)
		mt.ClearMockResponses()
	})

	mt.Run("UpdateCategory method Should return error When could not update object", func(mt *mtest.T) {
		categoriesCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "Con't update data",
		}))
		id := primitive.NewObjectID()

		var categoryService = CategoryService{}

		response, err := categoryService.Update(id, dto.InsertCategory{})

		assert.Nil(t, response)
		assert.NotNil(t, err)
		mt.ClearMockResponses()
	})

	mt.Run("UpdateCategory method Should update fields When object exists", func(mt *mtest.T) {
		categoriesCollection = mt.Coll
		id := primitive.NewObjectID()
		categoryData := mocked_data.GetValidInsertCategoryDto()
		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"value", mocked_data.GetBsonFromCategory(mocked_data.GetValidCategoryWithId(id))},
		})

		var categoryService = CategoryService{}

		response, err := categoryService.Update(id, categoryData)

		assert.NotNil(t, response)
		assert.Nil(t, err)
		mt.ClearMockResponses()
	})

	mt.Run("DeleteCategory method Should delete an item When the item can be deleted", func(mt *mtest.T) {
		categoriesCollection = mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 1}})
		var categoryService = CategoryService{}
		err := categoryService.Delete(primitive.NewObjectID())
		assert.Nil(t, err)
		mt.ClearMockResponses()
	})

	mt.Run("DeleteCategory method Should return no document deleted error When document dont exists", func(mt *mtest.T) {
		categoriesCollection = mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 0}})
		var categoryService = CategoryService{}
		err := categoryService.Delete(primitive.NewObjectID())
		assert.NotNil(t, err)
		mt.ClearMockResponses()
	})

	mt.Run("GetVideosByCategoryId method Should return object when has objects", func(mt *mtest.T) {
		videosCollection = mt.Coll
		firstId := primitive.NewObjectID()
		secondId := primitive.NewObjectID()

		firstCategory := mtest.CreateCursorResponse(1,"foo.bar", mtest.FirstBatch, mocked_data.GetBsonFromVideo(mocked_data.GetValidVideoWithId(firstId)))

		secondCategory := mtest.CreateCursorResponse(1,"foo.bar", mtest.NextBatch, mocked_data.GetBsonFromVideo(mocked_data.GetValidVideoWithId(secondId)))

		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(firstCategory, secondCategory, killCursors)

		var categoryService = CategoryService{}

		response, err := categoryService.GetVideosByCategoryId(primitive.ObjectID{})
		assert.Nil(t, err)
		assert.Equal(t, 2, len(response))
		mt.ClearMockResponses()
	})

	mt.Run("GetVideosByCategoryId method Should return error when dont has objects", func(mt *mtest.T) {
		videosCollection = mt.Coll

		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(bson.D{}, killCursors)

		var categoryService = CategoryService{}

		response, err := categoryService.GetVideosByCategoryId(primitive.ObjectID{})
		assert.NotNil(t, err)
		assert.Equal(t, 0, len(response))
		mt.ClearMockResponses()
	})

	mt.Run("GetFreeCategory method Should return free category object when object already exists", func(mt *mtest.T) {
		categoriesCollection = mt.Coll

		expectedCategory := mocked_data.GetValidCategory()

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, mocked_data.GetBsonFromCategory(expectedCategory)))

		var categoryService = CategoryService{}

		response := categoryService.GetFreeCategory()
		assert.NotNil(t, response)
		mt.ClearMockResponses()
	})

	mt.Run("GetFreeCategory method Should create free category and return object when object dont exists", func(mt *mtest.T) {
		categoriesCollection = mt.Coll
		expectedFreeCategory := models.GetFreeCategory()

		firstResponse := mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "Con't find data",
		})
		secondResponse := mtest.CreateSuccessResponse()
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(firstResponse, secondResponse, killCursors)

		var categoryService = CategoryService{}
		response := categoryService.GetFreeCategory()

		assert.NotNil(t, response)
		assert.Equal(t, expectedFreeCategory, response)
		mt.ClearMockResponses()
	})

	mt.Run("GetFreeCategory method Should return nil when free category dont exists and cant be created", func(mt *mtest.T) {
		categoriesCollection = mt.Coll

		firstResponse := mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "Con't find data",
		})
		secondResponse := mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "Con't insert data",
		})
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(firstResponse, secondResponse, killCursors)

		var categoryService = CategoryService{}

		response := categoryService.GetFreeCategory()
		assert.Nil(t, response)
		mt.ClearMockResponses()
	})
}