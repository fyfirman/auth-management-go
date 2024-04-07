package dto

import (
	"errors"
	"time"

	"github.com/fyfirman/auth-management-go/internal/datastruct"
)

type RegisterRequest struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=25"`
	Email    string `json:"email"    validate:"required,email"`
	Role     string `json:"role"     validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

func (r *RegisterRequest) Validate() error {
	if r.Role != datastruct.Admin.String() &&
		r.Role != datastruct.GeneralUser.String() &&
		r.Role != datastruct.SuperAdmin.String() {
		return errors.New("Invalid Role. Received: " + r.Role)
	}
	return nil
}

type RegisterResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
