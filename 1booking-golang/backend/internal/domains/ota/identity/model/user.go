package model

import (
	"time"
)

// User represents the core user entity
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone,omitempty"`
	PasswordHash string  `json:"-"` // Never expose password hash
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth,omitempty"`
	AvatarURL string    `json:"avatar_url,omitempty"`

	EmailVerified bool   `json:"email_verified"`
	PhoneVerified bool   `json:"phone_verified"`

	Status       string    `json:"status"` // active, inactive, suspended, etc.
	BlockedReason string   `json:"blocked_reason,omitempty"`

	LastLoginAt time.Time `json:"last_login_at,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
}