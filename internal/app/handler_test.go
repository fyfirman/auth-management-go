package app_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/fyfirman/auth-management-go/internal/app"
	"github.com/fyfirman/auth-management-go/internal/dto"
	"github.com/fyfirman/auth-management-go/internal/service/mocks"
	"github.com/fyfirman/auth-management-go/pkg"
	"github.com/stretchr/testify/mock"
)

func TestUserHandler_Register(t *testing.T) {
	mockUserService := new(mocks.UserServiceInterface)

	ctx := context.Background()
	// Create a new instance of UserHandler
	handler := app.NewUserHandler(mockUserService)

	response := &dto.RegisterResponse{
		ID:        12345,
		Username:  "john_doe",
		Email:     "john_doe@example.com",
		CreatedAt: time.Date(2024, 4, 8, 10, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2024, 4, 8, 12, 0, 0, 0, time.UTC),
	}

	mockUserService.Mock.On("RegisterUser", ctx, mock.AnythingOfType("*dto.RegisterRequest")).Return(response, nil)
	// Create a new HTTP request
	reqBody := []byte(`{"username": "testuser", "email": "testuser@gmail.com", "password": "testpassword", "role": "superadmin"}`)
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP response recorder
	recorder := httptest.NewRecorder()

	// Call the Register method
	handler.Register(recorder, req)

	mockUserService.Mock.AssertCalled(t, "RegisterUser", ctx, mock.AnythingOfType("*dto.RegisterRequest"))

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

	// Check the response body
	expectedResp := `{"id":12345,"username":"john_doe","email":"john_doe@example.com","created_at":"2024-04-08T10:00:00Z","updated_at":"2024-04-08T12:00:00Z"}`
	resultBody := recorder.Body.String()
	isSame, err := pkg.CompareJSONMaps(resultBody, expectedResp)
	if !isSame || err != nil {
		t.Errorf("expected response body %s, got %s", expectedResp, resultBody)
	}
}

func TestUserHandler_Register_InvalidRequest(t *testing.T) {
	mockUserService := new(mocks.UserServiceInterface)

	ctx := context.Background()
	// Create a new instance of UserHandler
	handler := app.NewUserHandler(mockUserService)

	response := &dto.RegisterResponse{
		ID:        12345,
		Username:  "john_doe",
		Email:     "john_doe@example.com",
		CreatedAt: time.Date(2024, 4, 8, 10, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2024, 4, 8, 12, 0, 0, 0, time.UTC),
	}

	mockUserService.Mock.On("RegisterUser", ctx, mock.AnythingOfType("*dto.RegisterRequest")).Return(response, nil)
	// Create a new HTTP request
	reqBody := []byte(`{"username": "testuser", "email": "testuser@gmail.com", "password": "testpassword", "role": "hack!!!"}`)
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP response recorder
	recorder := httptest.NewRecorder()

	// Call the Register method
	handler.Register(recorder, req)

	mockUserService.Mock.AssertNotCalled(t, "RegisterUser")
	// Check the response status code
	if recorder.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d, got %d", http.StatusBadRequest, recorder.Code)
	}

	// TODO: Check the response body
}

func TestResetPassword(t *testing.T) {
	t.Run("invalid request body", func(t *testing.T) {
		mockUserService := new(mocks.UserServiceInterface) // Reinitialize mock
		handler := app.NewUserHandler(mockUserService)     // Assume this correctly creates a handler

		req, _ := http.NewRequest("POST", "/reset-password", bytes.NewBufferString("{invalid json"))
		recorder := httptest.NewRecorder()

		handler.ResetPassword(recorder, req)

		if recorder.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, recorder.Code)
		}
	})

	t.Run("userService ResetPassword error", func(t *testing.T) {
		mockUserService := new(mocks.UserServiceInterface) // Reinitialize mock
		handler := app.NewUserHandler(mockUserService)     // Reset handler with a new mock

		reqBody, _ := json.Marshal(dto.ResetPasswordRequest{Email: "user@example.com"})
		req, _ := http.NewRequest("POST", "/reset-password", bytes.NewBuffer(reqBody))
		recorder := httptest.NewRecorder()

		mockUserService.On("ResetPassword", mock.Anything, mock.AnythingOfType("dto.ResetPasswordRequest")).Return(nil, errors.New("internal server error"))

		handler.ResetPassword(recorder, req)

		if recorder.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, recorder.Code)
		}
	})

	t.Run("successful reset password", func(t *testing.T) {
		mockUserService := new(mocks.UserServiceInterface) // Reinitialize mock for isolation
		handler := app.NewUserHandler(mockUserService)     // Reset handler with the new mock

		resetResponse := &dto.ResetPasswordResponse{Token: "Password reset successfully"}
		reqBody, _ := json.Marshal(dto.ResetPasswordRequest{Email: "user@example.com"})
		req, _ := http.NewRequest("POST", "/reset-password", bytes.NewBuffer(reqBody))
		recorder := httptest.NewRecorder()

		mockUserService.On("ResetPassword", mock.Anything, mock.AnythingOfType("dto.ResetPasswordRequest")).Return(resetResponse, nil)

		handler.ResetPassword(recorder, req)

		if recorder.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
		}

		var response dto.ResetPasswordResponse
		if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
			t.Fatal("failed to decode response")
		}

		if response.Token != resetResponse.Token {
			t.Errorf("expected message %q, got %q", resetResponse.Token, response.Token)
		}

		mockUserService.AssertExpectations(t) // Ensure all expected interactions were made
	})
}
