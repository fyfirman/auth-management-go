package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/fyfirman/auth-management-go/internal/datastruct"
	"github.com/fyfirman/auth-management-go/internal/dto"
	"github.com/fyfirman/auth-management-go/internal/repository/mocks"
	"github.com/fyfirman/auth-management-go/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestUserService_RegisterUser(t *testing.T) {
	userRepository := new(mocks.UserRepositoryInterface)
	tokenRepository := new(mocks.TokenRepositoryInterface)
	userService := service.NewUserService(userRepository, tokenRepository)

	ctx := context.TODO()
	req := &dto.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password",
	}

	// Mock the CreateUser method in UserRepository
	userRepository.Mock.On("CreateUser", ctx, mock.AnythingOfType("*datastruct.User")).Return(nil)

	// Call the RegisterUser method
	res, err := userService.RegisterUser(ctx, req)

	// Assert that the response is as expected
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, req.Username, res.Username)
	assert.Equal(t, req.Email, res.Email)

	// Assert that the CreateUser method was called with the correct arguments
	userRepository.Mock.AssertCalled(t, "CreateUser", ctx, mock.AnythingOfType("*datastruct.User"))
}

func TestUserService_Login(t *testing.T) {
	t.Setenv("JWT_SECRET", "secret_jwt")
	t.Setenv("JWT_EXPIRY_TIME", "100000")
	userRepository := new(mocks.UserRepositoryInterface)
	tokenRepository := new(mocks.TokenRepositoryInterface)

	userService := service.NewUserService(userRepository, tokenRepository)

	ctx := context.TODO()
	email := "test@example.com"
	password := "password"

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// Mock the FindByEmail method in UserRepository
	user := &datastruct.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
	}
	userRepository.Mock.On("FindByEmail", ctx, email).Return(user, nil)

	// Call the Login method
	req := dto.LoginRequest{
		Email:    email,
		Password: password,
	}
	res, err := userService.Login(ctx, req)

	// Assert that the response is as expected
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.Token)

	// Assert that the FindByEmail method was called with the correct arguments
	userRepository.Mock.AssertCalled(t, "FindByEmail", ctx, email)

	// Assert that bcrypt.CompareHashAndPassword was called with the correct arguments
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	assert.NoError(t, err)
}

func TestUserService_Login_InvalidCredentials(t *testing.T) {
	userRepository := new(mocks.UserRepositoryInterface)
	mockTokenRepo := new(mocks.TokenRepositoryInterface)
	userService := service.NewUserService(userRepository, mockTokenRepo)

	ctx := context.TODO()
	email := "test@example.com"
	password := "wrongpassword"

	// Mock the FindByEmail method in UserRepository
	user := &datastruct.User{
		Email:        email,
		PasswordHash: "$2a$10$4vY5z7j8k9l0m1n2o3p4q5r6s7t8u9v0w1x2y3z4a5b6c7d8e9f0g",
	}
	userRepository.Mock.On("FindByEmail", ctx, email).Return(user, nil)

	// Call the Login method with invalid credentials
	req := dto.LoginRequest{
		Email:    email,
		Password: password,
	}
	res, err := userService.Login(ctx, req)

	// Assert that the response is as expected
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, "invalid credentials", err.Error())

	// Assert that the FindByEmail method was called with the correct arguments
	userRepository.Mock.AssertCalled(t, "FindByEmail", ctx, email)

	// Assert that bcrypt.CompareHashAndPassword was called with the correct arguments
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	assert.Error(t, err)
	assert.True(t, errors.Is(err, bcrypt.ErrMismatchedHashAndPassword))
}

func TestResetPassword(t *testing.T) {
	ctx := context.Background()

	user := &datastruct.User{ID: 1, Email: "test@example.com"}

	t.Run("success", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepositoryInterface)
		mockTokenRepo := new(mocks.TokenRepositoryInterface)
		userService := service.NewUserService(mockUserRepo, mockTokenRepo)

		mockUserRepo.On("FindByEmail", ctx, user.Email).Return(user, nil)
		mockTokenRepo.On("CreateToken", ctx, mock.AnythingOfType("*datastruct.Token")).Return(nil)

		resp, err := userService.ResetPassword(ctx, dto.ResetPasswordRequest{Email: user.Email})

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Token)
		mockUserRepo.AssertExpectations(t)
		mockTokenRepo.AssertExpectations(t)
	})

	t.Run("FindByEmail error", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepositoryInterface)
		mockTokenRepo := new(mocks.TokenRepositoryInterface)
		userService := service.NewUserService(mockUserRepo, mockTokenRepo)

		mockUserRepo.On("FindByEmail", ctx, user.Email).Return(nil, errors.New("user not found"))

		resp, err := userService.ResetPassword(ctx, dto.ResetPasswordRequest{Email: user.Email})

		assert.Error(t, err)
		assert.Nil(t, resp)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("CreateToken error", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepositoryInterface)
		mockTokenRepo := new(mocks.TokenRepositoryInterface)
		userService := service.NewUserService(mockUserRepo, mockTokenRepo)

		mockUserRepo.On("FindByEmail", ctx, user.Email).Return(user, nil)
		mockTokenRepo.On("CreateToken", ctx, mock.AnythingOfType("*datastruct.Token")).Return(errors.New("db error"))

		resp, err := userService.ResetPassword(ctx, dto.ResetPasswordRequest{Email: user.Email})

		assert.Error(t, err)
		assert.Nil(t, resp)
		mockUserRepo.AssertExpectations(t)
		mockTokenRepo.AssertExpectations(t)
	})
}
