package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:100"`
	Username string `gorm:"size:30"`
	Email    string `gorm:"unique"`
	Password string
}

type UserToken struct {
	ID        uint
	UserID    uint
	TokenHash string
	ExpiresAt time.Time
}
