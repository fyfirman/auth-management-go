package datastruct

import (
	"strings"
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

var (
	userRoleMap = map[string]UserRole{
		"superadmin":   SuperAdmin,
		"admin":        Admin,
		"general-user": GeneralUser,
	}
)

func ParseUserRole(str string) (UserRole, bool) {
	c, ok := userRoleMap[strings.ToLower(str)]
	return c, ok
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
