package datastruct

import (
	"time"
)

type UserRole int

const (
	SuperAdmin UserRole = iota
	Admin
	GeneralUser
)

func (r UserRole) String() string {
	return [...]string{"superadmin", "admin", "general-user"}[r]
}

type User struct {
	ID           uint     `gorm:"primaryKey"`
	Username     string   `gorm:"unique;not null"`
	Email        string   `gorm:"unique;not null"`
	Role         UserRole `gorm:"not null"`
	PasswordHash string   `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
