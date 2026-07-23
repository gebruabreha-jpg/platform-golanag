// Package model defines the domain data structures for the contact book.
// These structs represent the core entities of the application.
// They are shared across all layers (handler, service, repository).
package model

import "time"

// Contact represents a single contact entry in the book.
// JSON tags control how the field names appear in API responses.
// The json:"-" tag on Password would prevent it from being serialized (not used here).
type Contact struct {
	// ID is the auto-generated primary key from PostgreSQL (SERIAL).
	// It is set by the database on INSERT and returned via RETURNING.
	ID int `json:"id"`

	// FirstName is the contact's first name. Required.
	FirstName string `json:"first_name"`

	// LastName is the contact's last name. Required.
	LastName string `json:"last_name"`

	// Email is the contact's email address. Must be unique in the database.
	// The UNIQUE constraint in SQL enforces this at the database level.
	Email string `json:"email"`

	// Phone is the contact's phone number. Required.
	Phone string `json:"phone"`

	// CreatedAt is set by the database (DEFAULT now()) on INSERT.
	// It records when the contact was first created and is not updated on edits.
	CreatedAt time.Time `json:"created_at"`
}
