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

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*datastruct.User, error) {
	var user datastruct.User
	result := DB.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
