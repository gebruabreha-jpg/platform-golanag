package service

import (
	"connectme/internal/domain"
	"connectme/internal/repository"
	"errors"
)

type HousingService struct {
	housingRepo repository.HousingRepository
}

func NewHousingService(housingRepo repository.HousingRepository) *HousingService {
	return &HousingService{housingRepo: housingRepo}
}

type CreateListingRequest struct {
	LandlordID         string   `json:"landlord_id"`
	Title              string   `json:"title" validate:"required,min=5,max=200"`
	Description        string   `json:"description" validate:"required,min=20"`
	PropertyType       string   `json:"property_type" validate:"required"`
	RoomType           string   `json:"room_type"`
	Bedrooms           int      `json:"bedrooms"`
	Bathrooms          int      `json:"bathrooms"`
	MonthlyRent        float64  `json:"monthly_rent" validate:"required,min=0"`
	Currency           string   `json:"currency" validate:"required"`
	Deposit            float64  `json:"deposit"`
	Address            string   `json:"address" validate:"required"`
	City               string   `json:"city" validate:"required"`
	Country            string   `json:"country" validate:"required"`
	Latitude           float64  `json:"latitude"`
	Longitude          float64  `json:"longitude"`
	AvailableFrom      string   `json:"available_from"`
	LeaseTerm          string   `json:"lease_term"`
	Furnished          bool     `json:"furnished"`
	IncludesUtilities  bool     `json:"includes_utilities"`
	ImageURLs          []string `json:"image_urls"`
}

func (s *HousingService) CreateListing(req CreateListingRequest) (*domain.HousingListing, error) {
	listing := &domain.HousingListing{
		LandlordID:        req.LandlordID,
		Title:             req.Title,
		Description:       req.Description,
		PropertyType:      req.PropertyType,
		RoomType:          req.RoomType,
		Bedrooms:          req.Bedrooms,
		Bathrooms:         req.Bathrooms,
		MonthlyRent:       req.MonthlyRent,
		Currency:          req.Currency,
		Deposit:           req.Deposit,
		Address:           req.Address,
		City:              req.City,
		Country:           req.Country,
		Latitude:          req.Latitude,
		Longitude:         req.Longitude,
		Furnished:         req.Furnished,
		IncludesUtilities: req.IncludesUtilities,
	}

	if err := s.housingRepo.CreateListing(listing); err != nil {
		return nil, err
	}

	return listing, nil
}

type ListListingsFilter struct {
	City        string `json:"city"`
	Country     string `json:"country"`
	PropertyType string `json:"property_type"`
	RoomType    string `json:"room_type"`
	MinRent     float64 `json:"min_rent"`
	MaxRent     float64 `json:"max_rent"`
	Limit       int   `json:"limit"`
	Offset      int   `json:"offset"`
}

func (s *HousingService) ListListings(filter ListListingsFilter) ([]*domain.HousingListing, int, error) {
	listings, total, err := s.housingRepo.List(filter, filter.Limit, filter.Offset)
	if err != nil {
		return nil, 0, err
	}
	// Count not implemented - would need to add to repository
	return listings, 0, nil
}

func (s *HousingService) GetListing(id string) (*domain.HousingListing, error) {
	listing, err := s.housingRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("listing not found")
	}
	return listing, nil
}

func (s *HousingService) SubmitApplication(listingID, userID string, message string) error {
	_, err := s.housingRepo.GetByID(listingID)
	if err != nil {
		return errors.New("listing not found")
	}

	_ = message // Would create application record

	return nil
}