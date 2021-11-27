package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertVideo_ConvertToVideo(t *testing.T) {
	videoToInsert := InsertVideo{
		Titulo:    "Input video test title",
		Descricao: "Input video test description",
		Url:       "https://www.url.com",
	}

	convertedVideo := videoToInsert.ConvertToVideo()

	assert.NotNil(t, convertedVideo.ID)
	assert.Equal(t, videoToInsert.Titulo, convertedVideo.Titulo, "Title must be the same.")
	assert.Equal(t, videoToInsert.Descricao, convertedVideo.Descricao, "Description must be the same.")
	assert.Equal(t, videoToInsert.Url, convertedVideo.Url, "Url must be the same.")
}

func TestInsertVideo_Validate(t *testing.T) {
	t.Run("Should return error when theres an empty title", func(t *testing.T) {
		videoToInsert := InsertVideo{
			Titulo:    "",
			Descricao: "Input video test description",
			Url:       "https://www.url.com",
		}

		err := videoToInsert.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "Titulo is required.", err.Error())
	})

	t.Run("Should return error when theres an empty description", func(t *testing.T) {
		videoToInsert := InsertVideo{
			Titulo:    "Input Title test",
			Descricao: "",
			Url:       "https://www.url.com",
		}

		err := videoToInsert.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "Descricao is required.", err.Error())
	})

	t.Run("Should return error when theres an empty url", func(t *testing.T) {
		videoToInsert := InsertVideo{
			Titulo:    "Input Title test",
			Descricao: "Input video test description",
			Url:       "",
		}

		err := videoToInsert.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "Url is required.", err.Error())
	})

	t.Run("Should return error when theres an invalid url", func(t *testing.T) {
		videoToInsert := InsertVideo{
			Titulo:    "Input Title test",
			Descricao: "Input video test description",
			Url:       "https//www.url.com",
		}

		err := videoToInsert.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "Url inválida.", err.Error())
	})

	t.Run("Should return nil when insert video object is valid", func(t *testing.T) {
		videoToInsert := InsertVideo{
			Titulo:    "Input Title test",
			Descricao: "Input video test description",
			Url:       "https://www.url.com",
		}

		err := videoToInsert.Validate()

		assert.Nil(t, err)
	})
}
