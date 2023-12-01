package postgres_test

import (
	"backend/model"
	"backend/repository/postgres"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

var (
	dbConnection *gorm.DB
)

func setup(t *testing.T) (*postgres.Database, func()) {
	//connect to database postgres
	dsn := "host=localhost user=postgres password=1234 dbname=integration_test port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(pg.Open(dsn), &gorm.Config{})
	assert.Nil(t, err)

	//do all migrations
	err = db.AutoMigrate(&model.User{})
	assert.Nil(t, err)

	dbConnection = db
	//return gorm connection object
	return postgres.New(db), func() {
		//cleanup
		db.Exec("delete from users")
	}
}

func createUser(t *testing.T) *model.User {
	user := &model.User{
		ID:       0,
		Email:    faker.Email(),
		Password: faker.Password(),
		Name:     faker.Name(),
	}

	err := dbConnection.Create(user).Error
	assert.Nil(t, err)

	return user
}
