package service

import (
	"connectme/internal/domain"
	"connectme/internal/repository"
	"errors"
)

type ServiceService struct {
	serviceRepo repository.ServiceRepository
}

func NewServiceService(serviceRepo repository.ServiceRepository) *ServiceService {
	return &ServiceService{serviceRepo: serviceRepo}
}

type ListLawyersFilter struct {
	Specialization string `json:"specialization"`
	Country        string `json:"country"`
	Search         string `json:"search"`
	Limit          int    `json:"limit"`
	Offset         int    `json:"offset"`
}

func (s *ServiceService) ListLawyers(filter ListLawyersFilter) ([]*domain.Lawyer, int, error) {
	filterMap := map[string]interface{}{}
	if filter.Specialization != "" {
		filterMap["specialization"] = filter.Specialization
	}
	if filter.Country != "" {
		filterMap["country"] = filter.Country
	}
	if filter.Search != "" {
		filterMap["search"] = filter.Search
	}

	lawyers, total, err := s.serviceRepo.ListLawyers(filterMap, filter.Limit, filter.Offset)
	if err != nil {
		return nil, 0, err
	}
	return lawyers, int(total), nil
}

func (s *ServiceService) GetLawyer(id string) (*domain.Lawyer, error) {
	lawyer, err := s.serviceRepo.GetLawyerByID(id)
	if err != nil {
		return nil, errors.New("lawyer not found")
	}
	return lawyer, nil
}

type BookLawyerRequest struct {
	LawyerID    string `json:"lawyer_id"`
	UserID      string `json:"user_id"`
	ScheduledAt string `json:"scheduled_at"`
	Duration    int    `json:"duration"`
	Type        string `json:"type"`
	Notes       string `json:"notes"`
}

func (s *ServiceService) BookLawyer(req BookLawyerRequest) (*domain.LawyerBooking, error) {
	// Verify lawyer exists
	_, err := s.serviceRepo.GetLawyerByID(req.LawyerID)
	if err != nil {
		return nil, errors.New("lawyer not found")
	}

	booking := &domain.LawyerBooking{
		LawyerID:    req.LawyerID,
		UserID:      req.UserID,
		Duration:    req.Duration,
		Type:        req.Type,
		Status:      "PENDING",
		Notes:       req.Notes,
	}

	if err := s.serviceRepo.CreateLawyerBooking(booking); err != nil {
		return nil, err
	}

	return booking, nil
}

func (s *ServiceService) GetCurrencyRates(from, to string) ([]*domain.CurrencyRate, error) {
	rates, err := s.serviceRepo.GetCurrencyRates(from, to)
	if err != nil {
		return nil, err
	}
	return rates, nil
}