package model

import (
	"time"
)

// Traveller represents a traveler profile associated with a user
type Traveller struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"` // FK to User

	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name"`

	Gender      string    `json:"gender,omitempty"` // e.g., male, female, other
	DateOfBirth time.Time `json:"date_of_birth,omitempty"`

	PassportNumber  string    `json:"passport_number,omitempty"`
	PassportCountry string    `json:"passport_country,omitempty"`
	PassportExpiry  time.Time `json:"passport_expiry,omitempty"`

	Nationality string `json:"nationality,omitempty"`

	FrequentFlyerNumber string `json:"frequent_flyer_number,omitempty"`
	KnownTravelerNumber string `json:"known_traveler_number,omitempty"`
	RedressNumber       string `json:"redress_number,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
