package models

import (
	"time"
)

// User model
type User struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	FullName     string    `json:"name"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	PasswordHash string    `json:"-"`
	Bio          string    `json:"bio"`
	AvatarURL    string    `json:"avatar_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	IsActive     bool      `json:"is_active"`
}
