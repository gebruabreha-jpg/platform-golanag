package postgres

import (
	"connectme/internal/domain"
	"gorm.io/gorm"
)

type ServiceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{db: db}
}

func (r *ServiceRepository) CreateScholarship(scholarship *domain.Scholarship) error {
	return r.db.Create(scholarship).Error
}

func (r *ServiceRepository) GetScholarships(filter map[string]interface{}, limit, offset int) ([]*domain.Scholarship, error) {
	query := r.db
	for key, value := range filter {
		query = query.Where(key+" = ?", value)
	}
	var scholarships []*domain.Scholarship
	err := query.Limit(limit).Offset(offset).Find(&scholarships).Error
	return scholarships, err
}

func (r *ServiceRepository) CreateJob(job *domain.Job) error {
	return r.db.Create(job).Error
}

func (r *ServiceRepository) GetJobs(filter map[string]interface{}, limit, offset int) ([]*domain.Job, error) {
	query := r.db
	for key, value := range filter {
		query = query.Where(key+" = ?", value)
	}
	var jobs []*domain.Job
	err := query.Limit(limit).Offset(offset).Find(&jobs).Error
	return jobs, err
}

func (r *ServiceRepository) UpdateCurrencyRate(rate *domain.CurrencyRate) error {
	return r.db.Save(rate).Error
}

func (r *ServiceRepository) GetLatestRate(from, to string) (*domain.CurrencyRate, error) {
	var rate domain.CurrencyRate
	err := r.db.Where("from_currency = ? AND to_currency = ?", from, to).
		Order("updated_at DESC").First(&rate).Error
	if err != nil {
		return nil, err
	}
	return &rate, nil
}

func (r *ServiceRepository) ListLawyers(filter map[string]interface{}, limit, offset int) ([]*domain.Lawyer, error) {
	query := r.db
	for key, value := range filter {
		query = query.Where(key+" = ?", value)
	}
	var lawyers []*domain.Lawyer
	err := query.Limit(limit).Offset(offset).Find(&lawyers).Error
	return lawyers, err
}

func (r *ServiceRepository) GetLawyerByID(id string) (*domain.Lawyer, error) {
	var lawyer domain.Lawyer
	err := r.db.First(&lawyer, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &lawyer, nil
}

func (r *ServiceRepository) CreateLawyerBooking(booking *domain.LawyerBooking) error {
	return r.db.Create(booking).Error
}

func (r *ServiceRepository) GetCurrencyRates(from, to string) ([]*domain.CurrencyRate, error) {
	var rates []*domain.CurrencyRate
	err := r.db.Where("from_currency = ? AND to_currency = ?", from, to).
		Order("updated_at DESC").Find(&rates).Error
	return rates, err
}