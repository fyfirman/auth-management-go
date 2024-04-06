package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fyfirman/auth-management-go/internal/dto"
	"github.com/fyfirman/auth-management-go/internal/service"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userService *service.UserService
	validator   *validator.Validate
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService, validator: validator.New()}
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrors []ValidationError

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(req); err != nil {
		errors := prepareValidationErrors(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
		return
	}

	resp, err := h.userService.RegisterUser(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.userService.Login(r.Context(), req)
	if err != nil {
		// Handle authentication error
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func prepareValidationErrors(err error) []map[string]string {
	var errors []map[string]string
	for _, err := range err.(validator.ValidationErrors) {
		var element = make(map[string]string)
		element["field"] = err.Field()
		element["message"] = fieldErrorMessage(err)
		errors = append(errors, element)
	}
	return errors
}

func fieldErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "email":
		return fmt.Sprintf("%s is not a valid email address", e.Field())
	case "min":
		// e.Param() contains the parameter of the validation tag, for min it's the minimum length/number
		return fmt.Sprintf("%s must be at least %s characters long", e.Field(), e.Param())
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s characters", e.Field(), e.Param())
	case "eqfield":
		return fmt.Sprintf("%s must be equal to %s", e.Field(), e.Param())
	case "alphanum":
		return fmt.Sprintf("%s must contain alphanumeric characters only", e.Field())
	// Add more cases as needed based on your validation rules
	default:
		return fmt.Sprintf("%s is not valid", e.Field())
	}
}
