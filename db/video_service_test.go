package db

import (
	"github.com/cristovaoolegario/aluraflix-api/mocked_data"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestVideoService(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("GetAll method Should return object when has objects", func(mt *mtest.T) {
		videosCollection = mt.Coll
		firstId := primitive.NewObjectID()
		secondId := primitive.NewObjectID()

		firstVideo := mtest.CreateCursorResponse(1,"foo.bar", mtest.FirstBatch, mocked_data.GetBsonFromVideo(mocked_data.GetValidVideoWithId(firstId)))

		secondVideo := mtest.CreateCursorResponse(1,"foo.bar", mtest.NextBatch, mocked_data.GetBsonFromVideo(mocked_data.GetValidVideoWithId(secondId)))

		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(firstVideo, secondVideo, killCursors)

		var videoService = VideoService{}

		videoResponse, err := videoService.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(videoResponse))
		mt.ClearMockResponses()
	})

	mt.Run("GetAll method Should return error when dont has objects", func(mt *mtest.T) {
		videosCollection = mt.Coll

		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(bson.D{}, killCursors)

		var videoService = VideoService{}

		videoResponse, err := videoService.GetAll()
		assert.NotNil(t, err)
		assert.Equal(t, 0, len(videoResponse))
		mt.ClearMockResponses()
	})

	mt.Run("GetByID method Should return object when object with id exists", func(mt *mtest.T) {
		videosCollection = mt.Coll
		id := primitive.NewObjectID()
		expectedVideo := mocked_data.GetValidVideoWithId(id)

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, mocked_data.GetBsonFromVideo(expectedVideo)))

		var videoService = VideoService{}

		videoResponse, err := videoService.GetByID(expectedVideo.ID)
		assert.Nil(t, err)
		assert.Equal(t, expectedVideo, videoResponse)
		mt.ClearMockResponses()
	})

	mt.Run("GetByID method Should return error when object dont exists", func(mt *mtest.T) {
		videosCollection = mt.Coll
		id := primitive.NewObjectID()

		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{}))

		var videoService = VideoService{}

		videoResponse, err := videoService.GetByID(id)
		assert.NotNil(t, err)
		assert.Nil(t, videoResponse)
		mt.ClearMockResponses()
	})
}