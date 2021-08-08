package db

import (
	"github.com/cristovaoolegario/aluraflix-api/mocked_data"
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



}