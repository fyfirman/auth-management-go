package repository

import (
	"context"

	"github.com/fyfirman/auth-management-go/internal/datastruct"
)

type TokenRepositoryInterface interface {
	CreateToken(ctx context.Context, user *datastruct.Token) error
}

type TokenRepository struct{}

func NewTokenRepository() *TokenRepository {
	return &TokenRepository{}
}

func (r *TokenRepository) CreateToken(ctx context.Context, user *datastruct.Token) error {
	result := DB.WithContext(ctx).Create(user)
	return result.Error
}
