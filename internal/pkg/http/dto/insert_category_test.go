package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func TestInsertCategory_Validate(t *testing.T) {
	t.Run("Should return error when titulo is empty", func(t *testing.T) {
		ic := InsertCategory{
			Titulo: "",
			Cor:    "Red",
		}
		err := ic.Validate()

		assert.Equal(t, "Titulo is required.", err.Error())
	})

	t.Run("Should return error when cor is empty", func(t *testing.T) {
		ic := InsertCategory{
			Titulo: "Unit Test Title",
			Cor:    "",
		}
		err := ic.Validate()

		assert.Equal(t, "Cor is required.", err.Error())
	})

	t.Run("Should not return error when insert category is valid", func(t *testing.T) {
		ic := InsertCategory{
			Titulo: "Unit Test Title",
			Cor:    "Red",
		}
		err := ic.Validate()

		assert.Nil(t, err)
	})
}
