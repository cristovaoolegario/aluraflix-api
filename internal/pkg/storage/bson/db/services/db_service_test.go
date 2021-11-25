package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBService_mountServerConnection(t *testing.T) {
	t.Run("Should return local db connection if environment id empty or dev", func(t *testing.T) {
		environments := []string{"", "dev"}
		for _, env := range environments {

			connectionString := mountServerConnection(env, "", "", "", "")

			assert.Equal(t, "mongodb://mongo:27017/dev_env", connectionString)
		}
	})

	t.Run("Should return mount connection string when environment is different than dev or empty", func(t *testing.T) {
		environments := []string{"staging", "prod"}
		for _, env := range environments {

			connectionString := mountServerConnection(env, "test_user", "test_password", "test_hostname", "test_db")

			assert.Equal(t, "mongodb+srv://test_user:test_password@test_hostname/test_db?retryWrites=true&w=majority", connectionString)
		}
	})
}
