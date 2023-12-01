package user_test

import (
	"backend/model"
	"backend/unit_test_with_dependency/user"
	"fmt"
	"testing"
)

type mockRepo struct {
}

func (m *mockRepo) GetUserByEmail(email string) (*model.User, error) {
	if email == "" {
		return nil, fmt.Errorf("")
	}
	return &model.User{
		ID:       1,
		Email:    email,
		Password: "1234",
		Name:     "jack",
	}, nil
}
func TestLogin(t *testing.T) {
	t.Run("no user found", func(t *testing.T) {
		m := new(mockRepo)
		service := user.NewService(m)
		_, err := service.Login(&user.LoginRequest{
			Email:    "",
			Password: "1234",
		})

		if err == nil {
			t.Fatal("error is nil ")
		}
	})

	t.Run("password is not correct", func(t *testing.T) {
		m := new(mockRepo)
		service := user.NewService(m)
		_, err := service.Login(&user.LoginRequest{
			Email:    "hello@example.com",
			Password: "1111",
		})

		if err == nil {
			t.Fatal("error is nil ")
		}
	})

	t.Run("successful login", func(t *testing.T) {
		m := new(mockRepo)
		service := user.NewService(m)
		resp, err := service.Login(&user.LoginRequest{
			Email:    "hello@example.com",
			Password: "1234",
		})

		if err != nil {
			t.Fatal("error is not nil ")
		}

		if resp.Token == "" {
			t.Fatal("token is empty")
		}

	})
}
