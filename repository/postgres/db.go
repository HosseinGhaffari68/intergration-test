package postgres

import (
	"backend/model"
	"gorm.io/gorm"
)

type Database struct {
	connection *gorm.DB
}

func New(connection *gorm.DB) *Database {
	return &Database{connection: connection}
}

func (d *Database) GetUserByEmail(email string) (*model.User, error) {
	user := new(model.User)
	if err := d.connection.Model(user).Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
