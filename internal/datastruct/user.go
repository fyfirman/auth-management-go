package datastruct

import "time"

type User struct {
	ID           int64
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
