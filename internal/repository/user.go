package repository

import (
	"context"

	"github.com/fyfirman/auth-management-go/internal/datastruct"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *datastruct.User) error {
	result := DB.WithContext(ctx).Create(user)
	return result.Error
}
