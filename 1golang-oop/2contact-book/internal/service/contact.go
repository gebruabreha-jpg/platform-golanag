package service

import (
	"context"
	"errors"
	"strings"

	"contact-book-api/internal/model"
	"contact-book-api/internal/repository"

)


//custome error


//
type ContactService struct {
	contacts repository.ContactRepository
}

// NewContactService creates a new service instance.
func NewContactService(contacts repository.ContactRepository) *ContactService {
	return &ContactService{contacts: contacts}
}

// all end pnit  bussines logic like creat, update, get  and delete
// they need repository obect intailzed  with consutrcutor

// Create validates input and creates a new contact.
func (s *ContactService) Create(ctx context.Context, firstName, lastName, email, phone string) (model.Contact, error) {
	// 1. Validate required fields
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)
	email = strings.TrimSpace(email)
	phone = strings.TrimSpace(phone)

	if firstName == "" {
		return model.Contact{}, errors.New("first name is required")
	}
	if lastName == "" {
		return model.Contact{}, errors.New("last name is required")
	}
	if email == "" {
		return model.Contact{}, errors.New("email is required")
	}
	if phone == "" {
		return model.Contact{}, errors.New("phone is required")
	}

	// 2. Check email uniqueness (case-insensitive)
	existing, err := s.contacts.GetByEmail(ctx, strings.ToLower(email))
	if err == nil {
		return model.Contact{}, repository.ErrEmailExists
	}
	if !errors.Is(err, repository.ErrContactNotFound) {
		return model.Contact{}, err
	}

	// 3. Create via repository
	return s.contacts.Create(ctx, firstName, lastName, email, phone)
}

// GetByID retrieves a contact by ID.
func (s *ContactService) GetByID(ctx context.Context, id int) (*model.Contact, error) {
	if id <= 0 {
		return nil, errors.New("invalid contact id")
	}
	return s.contacts.GetByID(ctx, id)
}

// GetAll returns all contacts.
func (s *ContactService) GetAll(ctx context.Context) ([]model.Contact, error) {
	return s.contacts.GetAll(ctx)
}

// Update validates input and updates an existing contact.
func (s *ContactService) Update(ctx context.Context, id int, firstName, lastName, email, phone string) (*model.Contact, error) {
	// 1. Validate required fields
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)
	email = strings.TrimSpace(email)
	phone = strings.TrimSpace(phone)

	if firstName == "" {
		return nil, errors.New("first name is required")
	}
	if lastName == "" {
		return nil, errors.New("last name is required")
	}
	if email == "" {
		return nil, errors.New("email is required")
	}
	if phone == "" {
		return nil, errors.New("phone is required")
	}

	// 2. Check if contact exists
	_, err := s.contacts.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 3. Check email uniqueness (allow same email if it belongs to same contact)
	existing, err := s.contacts.GetByEmail(ctx, strings.ToLower(email))
	if err == nil && existing.ID != id {
		return nil, repository.ErrEmailExists
	}
	if err != nil && !errors.Is(err, repository.ErrContactNotFound) {
		return nil, err
	}

	// 4. Update via repository
	return s.contacts.Update(ctx, id, firstName, lastName, email, phone)
}

// Delete removes a contact by ID.
func (s *ContactService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid contact id")
	}
	return s.contacts.Delete(ctx, id)
}



