package postgres

import (
	"connectme/internal/domain"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func (r *PostRepository) Create(post *domain.Post) error {
	return r.db.Create(post).Error
}

func (r *PostRepository) GetByID(id string) (*domain.Post, error) {
	var post domain.Post
	err := r.db.First(&post, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) ListByCommunity(communityID string, limit, offset int) ([]*domain.Post, error) {
	var posts []*domain.Post
	err := r.db.Where("community_id = ?", communityID).
		Order("created_at DESC").Limit(limit).Offset(offset).Find(&posts).Error
	return posts, err
}

func (r *PostRepository) Update(post *domain.Post) error {
	return r.db.Save(post).Error
}

func (r *PostRepository) IncrementReplyCount(id string) error {
	return r.db.Model(&domain.Post{}).Where("id = ?", id).UpdateColumn("reply_count", gorm.Expr("reply_count + 1")).Error
}

func (r *PostRepository) IncrementViewCount(id string) error {
	return r.db.Model(&domain.Post{}).Where("id = ?", id).UpdateColumn("view_count", gorm.Expr("view_count + 1")).Error
}