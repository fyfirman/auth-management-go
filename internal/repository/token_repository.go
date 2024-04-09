package repository

import (
	"context"

	"github.com/fyfirman/auth-management-go/internal/datastruct"
)

type TokenRepositoryInterface interface {
	CreateToken(ctx context.Context, user *datastruct.Token) error
	FindByToken(ctx context.Context, token string) (*datastruct.Token, error)
}

type TokenRepository struct{}

func NewTokenRepository() *TokenRepository {
	return &TokenRepository{}
}

func (r *TokenRepository) CreateToken(ctx context.Context, user *datastruct.Token) error {
	result := DB.WithContext(ctx).Create(user)
	return result.Error
}

func (r *TokenRepository) FindByToken(ctx context.Context, token string) (*datastruct.Token, error) {
	var tokenData datastruct.Token
	result := DB.WithContext(ctx).Where("token = ?", token).First(&tokenData)
	if result.Error != nil {
		return nil, result.Error
	}
	return &tokenData, nil
}
