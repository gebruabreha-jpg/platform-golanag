package postgres

import (
	"connectme/internal/domain"
	"gorm.io/gorm"
)

type HousingRepository struct {
	db *gorm.DB
}

func NewHousingRepository(db *gorm.DB) *HousingRepository {
	return &HousingRepository{db: db}
}

func (r *HousingRepository) CreateListing(listing *domain.HousingListing) error {
	return r.db.Create(listing).Error
}

func (r *HousingRepository) GetByID(id string) (*domain.HousingListing, error) {
	var listing domain.HousingListing
	err := r.db.First(&listing, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &listing, nil
}

func (r *HousingRepository) List(filter map[string]interface{}, limit, offset int) ([]*domain.HousingListing, error) {
	query := r.db
	for key, value := range filter {
		query = query.Where(key+" = ?", value)
	}
	var listings []*domain.HousingListing
	err := query.Limit(limit).Offset(offset).Order("created_at DESC").Find(&listings).Error
	return listings, err
}

func (r *HousingRepository) UpdateListing(listing *domain.HousingListing) error {
	return r.db.Save(listing).Error
}

func (r *HousingRepository) IncrementApplicationCount(id string) error {
	return r.db.Model(&domain.HousingListing{}).Where("id = ?", id).UpdateColumn("application_count", gorm.Expr("application_count + 1")).Error
}