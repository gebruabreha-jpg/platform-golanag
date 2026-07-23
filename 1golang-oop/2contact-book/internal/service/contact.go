package service

import (
	"context"
	"errors"
	"strings"

	"contact-book-api/internal/model"
	"contact-book-api/internal/repository"
)

// Sentinel errors for service-layer failure modes.
// These are what the handler checks to return the correct HTTP status.
var (
	ErrInvalidContactID = errors.New("invalid contact id")
	ErrContactNotFound  = errors.New("contact not found")
	ErrEmailExists      = errors.New("email already exists")
)

// ContactServiceInterface defines the contract for contact business logic.
// The handler depends on this interface, not the concrete ContactService type.
// This allows swapping implementations (e.g., mock for tests) without changing the handler.
type ContactServiceInterface interface {
	Create(ctx context.Context, firstName, lastName, email, phone string) (model.Contact, error)
	GetByID(ctx context.Context, id int) (*model.Contact, error)
	GetAll(ctx context.Context) ([]model.Contact, error)
	Update(ctx context.Context, id int, firstName, lastName, email, phone string) (*model.Contact, error)
	Delete(ctx context.Context, id int) error
}

// ContactService holds the repository it delegates to.
// The service contains business logic and validation.
// It does NOT know about HTTP — it only works with models and errors.
type ContactService struct {
	contacts repository.ContactRepository
}

// NewContactService injects the repository dependency via constructor.
// Called in main.go when wiring layers together.
// The repository interface means we can swap Postgres for a mock in tests.
func NewContactService(contacts repository.ContactRepository) *ContactService {
	return &ContactService{contacts: contacts}
}

// validateContact checks that all required fields are non-empty.
// This is a helper used by Create and Update to avoid duplicated checks.
func validateContact(firstName, lastName, email, phone string) error {
	if strings.TrimSpace(firstName) == "" {
		return errors.New("first name is required")
	}
	if strings.TrimSpace(lastName) == "" {
		return errors.New("last name is required")
	}
	if strings.TrimSpace(email) == "" {
		return errors.New("email is required")
	}
	if strings.TrimSpace(phone) == "" {
		return errors.New("phone is required")
	}
	return nil
}

// Create validates input, checks email uniqueness, then saves via repository.
// Flow: validate → check uniqueness → save → return created contact.
func (s *ContactService) Create(ctx context.Context, firstName, lastName, email, phone string) (model.Contact, error) {
	// Step 1: validate required fields.
	if err := validateContact(firstName, lastName, email, phone); err != nil {
		return model.Contact{}, err
	}

	// Step 2: check email uniqueness (case-insensitive via LOWER() in SQL).
	// We discard the returned contact value with _ because we only care
	// whether the email exists or not — we don't need the contact data here.
	// GetByEmail returns ErrContactNotFound if no match — that's expected (email is free).
	// If it returns nil (no error), the email already exists — reject it.
	_, err := s.contacts.GetByEmail(ctx, strings.ToLower(email))
	if err == nil {
		return model.Contact{}, ErrEmailExists
	}
	if !errors.Is(err, repository.ErrContactNotFound) {
		return model.Contact{}, err
	}

	// Step 3: save to database.
	return s.contacts.Create(ctx, firstName, lastName, email, phone)
}

// GetByID retrieves a contact by ID.
// Returns ErrInvalidContactID if id is not positive.
// Returns ErrContactNotFound if no row matches.
func (s *ContactService) GetByID(ctx context.Context, id int) (*model.Contact, error) {
	if id <= 0 {
		return nil, ErrInvalidContactID
	}
	contact, err := s.contacts.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrContactNotFound) {
			return nil, ErrContactNotFound
		}
		return nil, err
	}
	return contact, nil
}

// GetAll returns all contacts.
func (s *ContactService) GetAll(ctx context.Context) ([]model.Contact, error) {
	return s.contacts.GetAll(ctx)
}

// Update validates input and updates an existing contact.
// Returns ErrInvalidContactID if id is not positive.
// Returns ErrContactNotFound if no row matches the id.
// Returns ErrEmailExists if the new email violates the unique constraint.
// Flow: validate → update via repo (repo handles existence check via RETURNING).
func (s *ContactService) Update(ctx context.Context, id int, firstName, lastName, email, phone string) (*model.Contact, error) {
	// Step 1: validate required fields.
	if err := validateContact(firstName, lastName, email, phone); err != nil {
		return nil, err
	}

	// Step 2: update via repository.Translate repo errors to service-level errors.
	contact, err := s.contacts.Update(ctx, id, firstName, lastName, email, phone)
	if err != nil {
		if errors.Is(err, repository.ErrContactNotFound) {
			return nil, ErrContactNotFound
		}
		if errors.Is(err, repository.ErrEmailExists) {
			return nil, ErrEmailExists
		}
		return nil, err
	}
	return contact, nil
}

// Delete removes a contact by ID.
// Returns ErrInvalidContactID if id is not positive.
// Returns ErrContactNotFound if no row was deleted.
func (s *ContactService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return ErrInvalidContactID
	}
	err := s.contacts.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrContactNotFound) {
			return ErrContactNotFound
		}
		return err
	}
	return nil
}
