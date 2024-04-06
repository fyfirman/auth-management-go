package service

import (
	"time"

	"github.com/fyfirman/auth-management-go/internal/dto"
)

type UserService struct {
	// Here, you would typically have a repository for database operations.
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) RegisterUser(req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	// Here, you would hash the password and save the user to the database.
	// This example simply returns a mock response.

	response := &dto.RegisterResponse{
		ID:        1, // Mock ID
		Username:  req.Username,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}

	return response, nil
}
