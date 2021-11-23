package dto

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestInsertVideo_Validate_ShouldReturnError_WhenTheresAnEmptyTitle(t *testing.T) {
	videoToInsert := InsertVideo{
		Titulo:    "",
		Descricao: "Input video test description",
		Url:       "https://www.url.com",
	}

	err := videoToInsert.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, "Titulo is required.", err.Error())
}

func TestInsertVideo_Validate_ShouldReturnError_WhenTheresAnEmptyDescription(t *testing.T) {
	videoToInsert := InsertVideo{
		Titulo:    "Input Title test",
		Descricao: "",
		Url:       "https://www.url.com",
	}

	err := videoToInsert.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, "Descricao is required.", err.Error())
}

func TestInsertVideo_Validate_ShouldReturnError_WhenTheresAnEmptyUrl(t *testing.T) {
	videoToInsert := InsertVideo{
		Titulo:    "Input Title test",
		Descricao: "Input video test description",
		Url:       "",
	}

	err := videoToInsert.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, "Url is required.", err.Error())
}

func TestInsertVideo_Validate_ShouldReturnError_WhenTheresAnInvalidUrl(t *testing.T) {
	videoToInsert := InsertVideo{
		Titulo:    "Input Title test",
		Descricao: "Input video test description",
		Url:       "https//www.url.com",
	}

	err := videoToInsert.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, "Url inválida.", err.Error())
}

func TestInsertVideo_Validate_ShouldReturnNil_WhenInsertVideoObjectIsValid(t *testing.T) {
	videoToInsert := InsertVideo{
		Titulo:    "Input Title test",
		Descricao: "Input video test description",
		Url:       "https://www.url.com",
	}

	err := videoToInsert.Validate()

	assert.Nil(t, err)
}
