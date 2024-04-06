package service

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/fyfirman/auth-management-go/internal/datastruct"
	"github.com/fyfirman/auth-management-go/internal/dto"
	"github.com/fyfirman/auth-management-go/internal/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository repository.UserRepositoryInterface
}

func NewUserService(userRepository repository.UserRepositoryInterface) *UserService {
	return &UserService{UserRepository: userRepository}
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

	err = s.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	response := &dto.RegisterResponse{
		ID:        int64(user.ID),
		Username:  req.Username,
		Email:     req.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

func (s *UserService) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.UserRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := generateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{Token: token}, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func generateJWT(userID uint) (string, error) {
	var jwtSecretKey = []byte(os.Getenv("JWT_SECRET"))
	expiryTimeInSecondsStr := os.Getenv("JWT_EXPIRY_TIME")
	expiryTimeInSeconds, err := strconv.Atoi(expiryTimeInSecondsStr)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Duration(expiryTimeInSeconds) * time.Second).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
