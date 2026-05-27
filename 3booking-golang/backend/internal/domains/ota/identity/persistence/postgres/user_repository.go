package postgres

import (
	"context"

	"1booking-golang/backend/internal/domains/ota/identity/model"
	"1booking-golang/backend/internal/domains/ota/identity/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

// userRepository implements repository.UserRepository using PostgreSQL.
type userRepository struct {
	db *pgxpool.Pool
}

// NewUserRepository creates a new PostgreSQL user repository.
func NewUserRepository(db *pgxpool.Pool) repository.UserRepository {
	return &userRepository{db: db}
}

// FindByEmail returns a user by email.
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	query := `
		SELECT id, email, phone, password_hash, first_name, last_name, date_of_birth, avatar_url,
		       email_verified, phone_verified, status, blocked_reason, last_login_at,
		       created_at, updated_at, deleted_at
		FROM users
		WHERE email = $1 AND deleted_at IS NULL
	`
	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.Phone, &user.PasswordHash, &user.FirstName, &user.LastName,
		&user.DateOfBirth, &user.AvatarURL, &user.EmailVerified, &user.PhoneVerified,
		&user.Status, &user.BlockedReason, &user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID returns a user by ID.
func (r *userRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	query := `
		SELECT id, email, phone, password_hash, first_name, last_name, date_of_birth, avatar_url,
		       email_verified, phone_verified, status, blocked_reason, last_login_at,
		       created_at, updated_at, deleted_at
		FROM users
		WHERE id = $1 AND deleted_at IS NULL
	`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.Email, &user.Phone, &user.PasswordHash, &user.FirstName, &user.LastName,
		&user.DateOfBirth, &user.AvatarURL, &user.EmailVerified, &user.PhoneVerified,
		&user.Status, &user.BlockedReason, &user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Create inserts a new user.
func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	query := `
		INSERT INTO users (id, email, phone, password_hash, first_name, last_name, date_of_birth, avatar_url,
		                   email_verified, phone_verified, status, blocked_reason, last_login_at,
		                   created_at, updated_at, deleted_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
	`
	_, err := r.db.Exec(ctx, query,
		user.ID, user.Email, user.Phone, user.PasswordHash, user.FirstName, user.LastName,
		user.DateOfBirth, user.AvatarURL, user.EmailVerified, user.PhoneVerified,
		user.Status, user.BlockedReason, user.LastLoginAt, user.CreatedAt, user.UpdatedAt, user.DeletedAt,
	)
	return err
}

// Update updates an existing user.
func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	query := `
		UPDATE users
		SET email = $2, phone = $3, password_hash = $4, first_name = $5, last_name = $6,
		    date_of_birth = $7, avatar_url = $8, email_verified = $9, phone_verified = $10,
		    status = $11, blocked_reason = $12, last_login_at = $13, updated_at = $14
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(ctx, query,
		user.ID, user.Email, user.Phone, user.PasswordHash, user.FirstName, user.LastName,
		user.DateOfBirth, user.AvatarURL, user.EmailVerified, user.PhoneVerified,
		user.Status, user.BlockedReason, user.LastLoginAt, user.UpdatedAt,
	)
	return err
}

// Delete marks a user as deleted (soft delete).
func (r *userRepository) Delete(ctx context.Context, id string) error {
	query := `
		UPDATE users
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

// UpdateLastLoginAt updates the last login time for a user.
func (r *userRepository) UpdateLastLoginAt(ctx context.Context, id string) error {
	query := `
		UPDATE users
		SET last_login_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
