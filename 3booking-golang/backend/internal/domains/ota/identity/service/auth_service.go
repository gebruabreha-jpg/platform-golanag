package service

import (
	"context"

	"1booking-golang/backend/internal/domains/ota/identity/model"
	"1booking-golang/backend/internal/domains/ota/identity/repository"
	"1booking-golang/backend/internal/platform/utils"
)

// AuthService handles authentication business logic.
type AuthService struct {
	userRepo repository.UserRepository
	// tokenService *token.Service // Will be implemented later
}

// NewAuthService creates a new AuthService.
func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		// tokenService: tokenService,
	}
}

// Login authenticates a user with email and password.
func (s *AuthService) Login(ctx context.Context, email, password string) (*model.User, error) {
	// Find user by email
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// TODO: Verify password hash
	// Update last login
	err = s.userRepo.UpdateLastLoginAt(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Register creates a new user account.
func (s *AuthService) Register(ctx context.Context, user *model.User) error {
	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.PasswordHash = hashedPassword

	return s.userRepo.Create(ctx, user)
}

// Logout handles user logout (token invalidation would happen here).
func (s *AuthService) Logout(ctx context.Context, userID string) error {
	// Token invalidation logic would go here
	// For now, just return nil as a placeholder
	return nil
}

// RefreshToken generates a new access token from a refresh token.
func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*model.User, error) {
	// TODO: Validate refresh token and return associated user
	return nil, nil
}
