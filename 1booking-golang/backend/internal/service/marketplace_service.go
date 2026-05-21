package service

import (
	"connectme/internal/domain"
	"connectme/internal/repository"
	"encoding/json"
	"errors"
)

type MarketplaceService struct {
	itemRepo repository.MarketplaceRepository
}

func NewMarketplaceService(itemRepo repository.MarketplaceRepository) *MarketplaceService {
	return &MarketplaceService{itemRepo: itemRepo}
}

type CreateItemRequest struct {
	SellerID         string   `json:"seller_id"`
	Title            string   `json:"title" validate:"required,min=3,max=200"`
	Description      string   `json:"description" validate:"required,min=20"`
	Category         string   `json:"category" validate:"required"`
	Subcategory      string   `json:"subcategory"`
	Price            float64  `json:"price" validate:"required,min=0"`
	Currency         string   `json:"currency" validate:"required,oneof=USD ETB EUR GBP"`
	Condition        string   `json:"condition" validate:"required,oneof=NEW LIKE_NEW GOOD FAIR"`
	Location         string   `json:"location" validate:"required"`
	Country          string   `json:"country" validate:"required"`
	ShippingAvailable bool    `json:"shipping_available"`
	ShippingCost     float64 `json:"shipping_cost"`
	ImageURLs        []string `json:"image_urls"`
}

func (s *MarketplaceService) CreateItem(req CreateItemRequest) (*domain.MarketplaceItem, error) {
	imagesJSON, _ := json.Marshal(req.ImageURLs)
	item := &domain.MarketplaceItem{
		SellerID:          req.SellerID,
		Title:            req.Title,
		Description:      req.Description,
		Category:         req.Category,
		Subcategory:      req.Subcategory,
		Price:           req.Price,
		Currency:        req.Currency,
		Condition:       req.Condition,
		Location:        req.Location,
		Country:         req.Country,
		ShippingAvailable: req.ShippingAvailable,
		ShippingCost:    req.ShippingCost,
		ImageURLs:       string(imagesJSON),
		IsActive:        true,
	}

	if err := s.itemRepo.CreateItem(item); err != nil {
		return nil, err
	}

	return item, nil
}

type ListItemsFilter struct {
	Category    string `json:"category"`
	Country     string `json:"country"`
	City        string `json:"city"`
	MinPrice    float64 `json:"min_price"`
	MaxPrice    float64 `json:"max_price"`
	Condition   string `json:"condition"`
	Search      string `json:"search"`
	SortBy      string `json:"sort_by"` // price_asc, price_desc, newest, relevance
	Limit       int    `json:"limit"`
	Offset      int    `json:"offset"`
}

func (s *MarketplaceService) ListItems(filter ListItemsFilter) ([]*domain.MarketplaceItem, int, error) {
	filterMap := map[string]interface{}{}
	if filter.Category != "" {
		filterMap["category"] = filter.Category
	}
	if filter.Country != "" {
		filterMap["country"] = filter.Country
	}
	if filter.City != "" {
		filterMap["city"] = filter.City
	}
	if filter.Condition != "" {
		filterMap["condition"] = filter.Condition
	}

	items, err := s.itemRepo.List(filterMap, filter.Limit, filter.Offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.itemRepo.Count(filterMap)
	if err != nil {
		return nil, 0, err
	}

	return items, int(total), nil
}

func (s *MarketplaceService) GetItem(id string) (*domain.MarketplaceItem, error) {
	item, err := s.itemRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("item not found")
	}

	// Increment view count asynchronously
	go s.itemRepo.IncrementViewCount(id)

	return item, nil
}

type ExpressInterestRequest struct {
	ItemID         string `json:"item_id"`
	BuyerID        string `json:"buyer_id"`
	Message        string `json:"message"`
	 OfferedPrice  float64 `json:"offered_price"`
}

func (s *MarketplaceService) ExpressInterest(req ExpressInterestRequest) error {
	// Check item exists and is available
	item, err := s.itemRepo.GetByID(req.ItemID)
	if err != nil || !item.IsActive || item.IsSold {
		return errors.New("item is not available")
	}

	// Increment interest count
	if err := s.itemRepo.IncrementInterestCount(req.ItemID); err != nil {
		return err
	}

	// Create transaction record for escrow (future implementation)
	// Send notification to seller

	return nil
}
