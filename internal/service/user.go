package service

import (
	"context"
	"time"

	"github.com/fyfirman/auth-management-go/internal/datastruct"
	"github.com/fyfirman/auth-management-go/internal/dto"
	"github.com/fyfirman/auth-management-go/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) RegisterUser(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &datastruct.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hashedPassword,
	}

	err = s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	response := &dto.RegisterResponse{
		ID:        1, // Mock ID
		Username:  req.Username,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}

	return response, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
