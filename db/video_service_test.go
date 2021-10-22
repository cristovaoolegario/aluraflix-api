package db

import (
	"github.com/cristovaoolegario/aluraflix-api/dto"
	"github.com/cristovaoolegario/aluraflix-api/mocked_data"
	"github.com/cristovaoolegario/aluraflix-api/mocked_services"
	"github.com/cristovaoolegario/aluraflix-api/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestVideoService(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("GetAllFreeVideos method Should return object when has objects", func(mt *mtest.T) {

		var videoService = VideoService{}
		videoService.videosCollection = mt.Coll
		firstId := primitive.NewObjectID()
		secondId := primitive.NewObjectID()

		firstVideo := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, mocked_data.GetBsonFromVideo(mocked_data.GetValidVideoWithId(firstId)))

		secondVideo := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, mocked_data.GetBsonFromVideo(mocked_data.GetValidVideoWithId(secondId)))

		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(firstVideo, secondVideo, killCursors)

		videoService.categoryService = &mocked_services.CategoryServiceMock{}
		mocked_services.CategoryServiceMockGetFreeCategory = func() *models.Category {
			return models.GetFreeCategory()
		}

		videoResponse, err := videoService.GetAllFreeVideos()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(videoResponse))
		mt.ClearMockResponses()
	})

	mt.Run("GetAllFreeVideos method Should return error when dont has objects", func(mt *mtest.T) {

		var videoService = VideoService{}
		videoService.videosCollection = mt.Coll

		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(bson.D{}, killCursors)

		videoService.categoryService = &mocked_services.CategoryServiceMock{}
		mocked_services.CategoryServiceMockGetFreeCategory = func() *models.Category {
			return models.GetFreeCategory()
		}

		videoResponse, err := videoService.GetAllFreeVideos()
		assert.NotNil(t, err)
		assert.Nil(t, videoResponse)
		mt.ClearMockResponses()
	})

	mt.Run("GetAllVideos method Should return object when has objects", func(mt *mtest.T) {
		var videoService = VideoService{}
		videoService.videosCollection = mt.Coll

		firstId := primitive.NewObjectID()
		secondId := primitive.NewObjectID()

		firstVideo := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, mocked_data.GetBsonFromVideo(mocked_data.GetValidVideoWithId(firstId)))

		secondVideo := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, mocked_data.GetBsonFromVideo(mocked_data.GetValidVideoWithId(secondId)))

		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(firstVideo, secondVideo, killCursors)

		videoResponse, err := videoService.GetAll("", 1, 5)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(videoResponse))
		mt.ClearMockResponses()
	})

	mt.Run("GetAllVideos method with filter Should return object when has objects", func(mt *mtest.T) {
		var videoService = VideoService{}
		videoService.videosCollection = mt.Coll

		firstId := primitive.NewObjectID()
		secondId := primitive.NewObjectID()

		firstVideo := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, mocked_data.GetBsonFromVideo(mocked_data.GetValidVideoWithId(firstId)))

		secondVideo := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, mocked_data.GetBsonFromVideo(mocked_data.GetValidVideoWithId(secondId)))

		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(firstVideo, secondVideo, killCursors)

		videoResponse, err := videoService.GetAll("test", 1, 5)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(videoResponse))
		mt.ClearMockResponses()
	})

	mt.Run("GetAllVideos method Should return error when dont has objects", func(mt *mtest.T) {
		var videoService = VideoService{}
		videoService.videosCollection = mt.Coll

		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(bson.D{}, killCursors)

		videoResponse, err := videoService.GetAll("", 1, 5)
		assert.NotNil(t, err)
		assert.Equal(t, 0, len(videoResponse))
		mt.ClearMockResponses()
	})

	mt.Run("GetVideoByID method Should return object when object with id exists", func(mt *mtest.T) {
		var videoService = VideoService{}
		videoService.videosCollection = mt.Coll

		id := primitive.NewObjectID()
		expectedVideo := mocked_data.GetValidVideoWithId(id)

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, mocked_data.GetBsonFromVideo(expectedVideo)))

		videoResponse, err := videoService.GetByID(expectedVideo.ID)
		assert.Nil(t, err)
		assert.Equal(t, expectedVideo, videoResponse)
		mt.ClearMockResponses()
	})

	mt.Run("GetVideoByID method Should return error when object dont exists", func(mt *mtest.T) {
		var videoService = VideoService{}
		videoService.videosCollection = mt.Coll
		id := primitive.NewObjectID()

		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{}))

		videoResponse, err := videoService.GetByID(id)
		assert.NotNil(t, err)
		assert.Nil(t, videoResponse)
		mt.ClearMockResponses()
	})

	mt.Run("CreateVideo method Should return object when inserted", func(mt *mtest.T) {
		var videoService = VideoService{}
		videoService.videosCollection = mt.Coll
		id := primitive.NewObjectID()
		expectedCategory := mocked_data.GetValidCategoryWithId(id)

		firstResponse := mtest.CreateSuccessResponse()
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(firstResponse, killCursors)

		videoService.categoryService = &mocked_services.CategoryServiceMock{}
		mocked_services.CategoryServiceMockGetFreeCategory = func() *models.Category {
			return nil
		}
		mocked_services.CategoryServiceMockGetByID = func(id primitive.ObjectID) (*models.Category, error) {
			return expectedCategory, nil
		}

		insertedVideo, err := videoService.Create(mocked_data.GetValidInsertVideoDto())

		assert.NotNil(t, insertedVideo)
		assert.Nil(t, err)
		mt.ClearMockResponses()
	})

	mt.Run("CreateVideo method Should return error when could not insert", func(mt *mtest.T) {
		var videoService = VideoService{}
		videoService.videosCollection = mt.Coll

		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "Con't insert data",
		}))

		videoService.categoryService = &mocked_services.CategoryServiceMock{}
		mocked_services.CategoryServiceMockGetFreeCategory = func() *models.Category {
			return nil
		}
		mocked_services.CategoryServiceMockGetByID = func(id primitive.ObjectID) (*models.Category, error) {
			return mocked_data.GetValidCategory(), nil
		}

		insertedVideo, err := videoService.Create(dto.InsertVideo{})
		assert.Nil(t, insertedVideo)
		assert.NotNil(t, err)
		mt.ClearMockResponses()
	})

	mt.Run("CreateVideo method Should return error when category dont exist", func(mt *mtest.T) {
		var videoService = VideoService{}
		videoService.videosCollection = mt.Coll
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(bson.D{}, killCursors)

		videoService.categoryService = &mocked_services.CategoryServiceMock{}
		mocked_services.CategoryServiceMockGetByID = func(id primitive.ObjectID) (*models.Category, error) {
			return nil, mongo.ErrNoDocuments
		}

		insertedVideo, err := videoService.Create(dto.InsertVideo{})
		assert.Nil(t, insertedVideo)
		assert.NotNil(t, err)
		mt.ClearMockResponses()
	})

	mt.Run("UpdateVideo method Should update fields When object exists", func(mt *mtest.T) {
		var videoService = VideoService{}
		videoService.videosCollection = mt.Coll
		id := primitive.NewObjectID()
		videoData := mocked_data.GetValidInsertVideoDto()
		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"value", mocked_data.GetBsonFromVideo(mocked_data.GetValidVideoWithId(id))},
		})

		_, err := videoService.Update(id, videoData)

		assert.Nil(t, err)
		mt.ClearMockResponses()
	})

	mt.Run("UpdateVideo method Should return error When could not update object", func(mt *mtest.T) {
		var videoService = VideoService{}
		videoService.videosCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "Con't update data",
		}))
		id := primitive.NewObjectID()


		updateVideo, err := videoService.Update(id, dto.InsertVideo{})
		assert.Nil(t, updateVideo)
		assert.NotNil(t, err)
		mt.ClearMockResponses()
	})

	mt.Run("DeleteVideo method Should delete an item When the item can be deleted", func(mt *mtest.T) {
		var videoService = VideoService{}
		videoService.videosCollection = mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 1}})
		err := videoService.Delete(primitive.NewObjectID())
		assert.Nil(t, err)
		mt.ClearMockResponses()
	})

	mt.Run("DeleteVideo method Should return no document deleted error When document dont exists", func(mt *mtest.T) {
		var videoService = VideoService{}
		videoService.videosCollection = mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 0}})
		err := videoService.Delete(primitive.NewObjectID())
		assert.NotNil(t, err)
		mt.ClearMockResponses()
	})
}
