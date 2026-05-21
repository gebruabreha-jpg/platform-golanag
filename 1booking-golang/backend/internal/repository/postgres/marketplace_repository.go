package postgres

import (
	"connectme/internal/domain"
	"gorm.io/gorm"
)

type MarketplaceRepository struct {
	db *gorm.DB
}

func (r *MarketplaceRepository) CreateItem(item *domain.MarketplaceItem) error {
	return r.db.Create(item).Error
}

func (r *MarketplaceRepository) GetByID(id string) (*domain.MarketplaceItem, error) {
	var item domain.MarketplaceItem
	err := r.db.First(&item, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MarketplaceRepository) List(filter map[string]interface{}, limit, offset int) ([]*domain.MarketplaceItem, error) {
	query := r.db
	for key, value := range filter {
		query = query.Where(key+" = ?", value)
	}
	var items []*domain.MarketplaceItem
	err := query.Limit(limit).Offset(offset).Order("created_at DESC").Find(&items).Error
	return items, err
}

func (r *MarketplaceRepository) UpdateItem(item *domain.MarketplaceItem) error {
	return r.db.Save(item).Error
}

func (r *MarketplaceRepository) IncrementViewCount(id string) error {
	return r.db.Model(&domain.MarketplaceItem{}).Where("id = ?", id).UpdateColumn("view_count", gorm.Expr("view_count + 1")).Error
}

func (r *MarketplaceRepository) IncrementInterestCount(id string) error {
	return r.db.Model(&domain.MarketplaceItem{}).Where("id = ?", id).UpdateColumn("interest_count", gorm.Expr("interest_count + 1")).Error
}

func (r *MarketplaceRepository) MarkAsSold(id string) error {
	return r.db.Model(&domain.MarketplaceItem{}).Where("id = ?", id).Update("is_sold", true).Error
}

func (r *MarketplaceRepository) Count(filter map[string]interface{}) (int64, error) {
	var count int64
	query := r.db.Model(&domain.MarketplaceItem{})
	for key, value := range filter {
		query = query.Where(key+" = ?", value)
	}
	err := query.Count(&count).Error
	return count, err
}