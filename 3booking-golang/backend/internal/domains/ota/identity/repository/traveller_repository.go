package repository

import (
	"context"

	"1booking-golang/backend/internal/domains/ota/identity/model"
)

// TravellerRepository defines the interface for traveller data access.
type TravellerRepository interface {
	// FindByID returns a traveller by ID.
	FindByID(ctx context.Context, id string) (*model.Traveller, error)
	// FindByUserID returns travellers for a given user ID.
	FindByUserID(ctx context.Context, userID string) ([]*model.Traveller, error)
	// Create inserts a new traveller.
	Create(ctx context.Context, traveller *model.Traveller) error
	// Update updates an existing traveller.
	Update(ctx context.Context, traveller *model.Traveller) error
	// Delete removes a traveller by ID.
	Delete(ctx context.Context, id string) error
}
