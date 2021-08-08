package dto

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestInsertCategory_ConvertToCategory(t *testing.T) {
	insertCategory := InsertCategory{
		Titulo: "Unit test title",
		Cor:    "Blue",
	}
	category := insertCategory.ConvertToCategory()

	assert.Equal(t, insertCategory.Cor, category.Cor)
	assert.Equal(t, insertCategory.Titulo, category.Titulo)
	assert.Equal(t, true, category.Active)
	assert.IsType(t, primitive.ObjectID{}, category.ID)
}

func TestInsertVideo_Validate_ShouldReturnError_WhenTituloIsEmpty(t *testing.T) {
	ic := InsertCategory{
		Titulo: "",
		Cor: "Red",
	}
	err := ic.Validate()

	assert.Equal(t, "Titulo is required.", err.Error())
}

func TestInsertVideo_Validate_ShouldReturnError_WhenCorIsEmpty(t *testing.T) {
	ic := InsertCategory{
		Titulo: "Unit Test Title",
		Cor: "",
	}
	err := ic.Validate()

	assert.Equal(t, "Cor is required.", err.Error())
}

func TestInsertVideo_Validate_ShouldNotReturnError_WhenInsertCategoryIsValid(t *testing.T) {
	ic := InsertCategory{
		Titulo: "Unit Test Title",
		Cor: "Red",
	}
	err := ic.Validate()

	assert.Nil(t, err)
}


