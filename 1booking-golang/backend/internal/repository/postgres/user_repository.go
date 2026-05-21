package postgres

import (
	"connectme/internal/domain"
	"errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByID(id string) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByPhone(phone string) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, "phone = ?", phone).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) UpdateTrustScore(id string, score float64) error {
	return r.db.Model(&domain.User{}).Where("id = ?", id).Update("trust_score", score).Error
}

func (r *UserRepository) List(filter map[string]interface{}, limit, offset int) ([]*domain.User, error) {
	query := r.db
	for key, value := range filter {
		query = query.Where(key+" = ?", value)
	}
	var users []*domain.User
	err := query.Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}

func (r *UserRepository) Search(query string, limit int) ([]*domain.User, error) {
	var users []*domain.User
	err := r.db.Where("first_name ILIKE ? OR last_name ILIKE ? OR email ILIKE ?",
		"%"+query+"%", "%"+query+"%", "%"+query+"%").
		Limit(limit).Find(&users).Error
	return users, err
}
