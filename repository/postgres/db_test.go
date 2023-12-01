package postgres_test

import (
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserByEmail(t *testing.T) {
	//setup db
	db, cleanup := setup(t)

	defer cleanup()

	//create user with random email
	user := createUser(t)

	t.Run("error on user not found", func(t *testing.T) {
		_, err := db.GetUserByEmail(faker.Email())
		assert.NotNil(t, err)
	})

	t.Run("successful", func(t *testing.T) {
		u, err := db.GetUserByEmail(user.Email)
		assert.Nil(t, err)
		assert.Equal(t, user.ID, u.ID)
	})
}
