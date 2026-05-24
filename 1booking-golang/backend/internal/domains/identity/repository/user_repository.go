package repository

import (
	"context"

	"1booking-golang/backend/internal/domains/identity/model"
)

// UserRepository defines the interface for user data access.
type UserRepository interface {
	// FindByEmail returns a user by email.
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	// FindByID returns a user by ID.
	FindByID(ctx context.Context, id string) (*model.User, error)
	// Create inserts a new user.
	Create(ctx context.Context, user *model.User) error
	// Update updates an existing user.
	Update(ctx context.Context, user *model.User) error
	// Delete marks a user as deleted (soft delete).
	Delete(ctx context.Context, id string) error
	// UpdateLastLoginAt updates the last login time for a user.
	UpdateLastLoginAt(ctx context.Context, id string) error
}