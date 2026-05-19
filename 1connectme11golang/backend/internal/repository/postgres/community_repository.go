package postgres

import (
	"connectme/internal/domain"
	"gorm.io/gorm"
)

type CommunityRepository struct {
	db *gorm.DB
}

func NewCommunityRepository(db *gorm.DB) *CommunityRepository {
	return &CommunityRepository{db: db}
}

func (r *CommunityRepository) Create(community *domain.Community) error {
	return r.db.Create(community).Error
}

func (r *CommunityRepository) GetByID(id string) (*domain.Community, error) {
	var community domain.Community
	err := r.db.First(&community, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &community, nil
}

func (r *CommunityRepository) List(category, location string, limit, offset int) ([]*domain.Community, error) {
	query := r.db
	if category != "" {
		query = query.Where("category = ?", category)
	}
	if location != "" {
		query = query.Where("location = ?", location)
	}
	var communities []*domain.Community
	err := query.Limit(limit).Offset(offset).Find(&communities).Error
	return communities, err
}

func (r *CommunityRepository) Update(community *domain.Community) error {
	return r.db.Save(community).Error
}

func (r *CommunityRepository) IncrementMemberCount(id string) error {
	return r.db.Model(&domain.Community{}).Where("id = ?", id).UpdateColumn("member_count", gorm.Expr("member_count + 1")).Error
}

func (r *CommunityRepository) DecrementMemberCount(id string) error {
	return r.db.Model(&domain.Community{}).Where("id = ?", id).UpdateColumn("member_count", gorm.Expr("member_count - 1")).Error
}

func (r *CommunityRepository) Count(filter map[string]interface{}) (int64, error) {
	var count int64
	query := r.db.Model(&domain.Community{})
	for key, value := range filter {
		query = query.Where(key+" = ?", value)
	}
	err := query.Count(&count).Error
	return count, err
}