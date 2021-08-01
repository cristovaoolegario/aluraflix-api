package dto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertInsertVideoToVideo(t *testing.T){
	videoToInsert := InsertVideo{
		Titulo: "Input video test title",
		Descricao: "Input video test description",
		Url: "www.url.com",
	}

	convertedVideo := videoToInsert.ConvertToVideo()

	assert.NotNil(t, convertedVideo.ID)
	assert.Equal(t, videoToInsert.Titulo, convertedVideo.Titulo, "Title must be the same.")
	assert.Equal(t, videoToInsert.Descricao, convertedVideo.Descricao, "Description must be the same.")
	assert.Equal(t, videoToInsert.Url, convertedVideo.Url, "Url must be the same.")
}