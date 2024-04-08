package service

import (
	"context"
	"crypto/rand"
	"encoding/base32"
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

type UserServiceInterface interface {
	RegisterUser(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error)
	ResetPassword(ctx context.Context, req dto.ResetPasswordRequest) (*dto.ResetPasswordResponse, error)
}

type UserService struct {
	userRepository  repository.UserRepositoryInterface
	tokenRepository repository.TokenRepositoryInterface
}

func NewUserService(
	userRepository repository.UserRepositoryInterface,
	tokenRepository repository.TokenRepositoryInterface,
) *UserService {
	return &UserService{userRepository: userRepository, tokenRepository: tokenRepository}
}

func (s *UserService) RegisterUser(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &datastruct.User{
		Username:     req.Username,
		Email:        req.Email,
		Role:         req.Role,
		PasswordHash: hashedPassword,
	}

	err = s.userRepository.CreateUser(ctx, user)
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
	user, err := s.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := generateJWT(user)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{Token: token}, nil
}

func (s *UserService) ResetPassword(
	ctx context.Context,
	req dto.ResetPasswordRequest,
) (*dto.ResetPasswordResponse, error) {
	user, err := s.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	expiryTimeInSeconds := 60 * 60

	token := generateResetPasswordToken()

	err = s.tokenRepository.CreateToken(ctx, &datastruct.Token{
		Token:     token,
		UserId:    user.ID,
		ExpiredAt: time.Now().Add(time.Duration(expiryTimeInSeconds) * time.Second),
	})

	if err != nil {
		return nil, err
	}

	return &dto.ResetPasswordResponse{Token: token}, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func generateJWT(user *datastruct.User) (string, error) {
	var jwtSecretKey = []byte(os.Getenv("JWT_SECRET"))
	expiryTimeInSecondsStr := os.Getenv("JWT_EXPIRY_TIME")
	expiryTimeInSeconds, err := strconv.Atoi(expiryTimeInSecondsStr)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":       time.Now().Add(time.Duration(expiryTimeInSeconds) * time.Second).Unix(),
		"user_id":   user.ID,
		"user_role": user.Role,
	})

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func generateResetPasswordToken() string {
	bytes := make([]byte, 15)
	rand.Read(bytes)
	token := base32.StdEncoding.EncodeToString(bytes)

	return token
}
