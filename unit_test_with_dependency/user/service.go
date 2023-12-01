package user

import (
	"fmt"
)

type Repository interface {
	GetUserByEmail(email string) (*User, error)
}

type Service struct {
	//this is a dependency for this service
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

type User struct {
	ID       uint
	Email    string
	Password string
	Name     string
}

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	//jwt token
	Token string
}

func (s *Service) Login(request *LoginRequest) (*LoginResponse, error) {
	//get the user from database by Email
	user, err := s.repo.GetUserByEmail(request.Email)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	//check the user's password
	if user.Password != request.Password {
		return nil, fmt.Errorf("password is not correct")
	}
	//generate jwt token
	token := "234qakjsaqeprtiuqp0870aadfkeop"

	//return the response
	return &LoginResponse{Token: token}, nil
}
