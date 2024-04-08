package datastruct

import (
	"time"
)

type Token struct {
	Token     string `gorm:"unique;not null"`
	ExpiredAt time.Time
	UserId    uint `gorm:"unique;not null"`
}
