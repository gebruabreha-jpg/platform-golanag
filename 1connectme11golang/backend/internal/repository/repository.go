package repository

import (
	"connectme/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	GetByPhone(phone string) (*domain.User, error)
	Update(user *domain.User) error
	UpdateTrustScore(id string, score float64) error
	List(filter map[string]interface{}, limit, offset int) ([]*domain.User, error)
	Search(query string, limit int) ([]*domain.User, error)
}

type CommunityRepository interface {
	Create(community *domain.Community) error
	GetByID(id string) (*domain.Community, error)
	List(category, location string, limit, offset int) ([]*domain.Community, error)
	Update(community *domain.Community) error
	IncrementMemberCount(id string) error
	DecrementMemberCount(id string) error
}

type PostRepository interface {
	Create(post *domain.Post) error
	GetByID(id string) (*domain.Post, error)
	ListByCommunity(communityID string, limit, offset int) ([]*domain.Post, error)
	Update(post *domain.Post) error
	IncrementReplyCount(id string) error
	IncrementViewCount(id string) error
}

type MarketplaceRepository interface {
	CreateItem(item *domain.MarketplaceItem) error
	GetByID(id string) (*domain.MarketplaceItem, error)
	List(filter map[string]interface{}, limit, offset int) ([]*domain.MarketplaceItem, error)
	UpdateItem(item *domain.MarketplaceItem) error
	IncrementViewCount(id string) error
	IncrementInterestCount(id string) error
	MarkAsSold(id string) error
}

type HousingRepository interface {
	CreateListing(listing *domain.HousingListing) error
	GetByID(id string) (*domain.HousingListing, error)
	List(filter map[string]interface{}, limit, offset int) ([]*domain.HousingListing, error)
	UpdateListing(listing *domain.HousingListing) error
	IncrementApplicationCount(id string) error
}

type ServiceRepository interface {
	CreateScholarship(scholarship *domain.Scholarship) error
	GetScholarships(filter map[string]interface{}, limit, offset int) ([]*domain.Scholarship, error)
	CreateJob(job *domain.Job) error
	GetJobs(filter map[string]interface{}, limit, offset int) ([]*domain.Job, error)
	UpdateCurrencyRate(rate *domain.CurrencyRate) error
	GetLatestRate(from, to string) (*domain.CurrencyRate, error)
	ListLawyers(filter map[string]interface{}, limit, offset int) ([]*domain.Lawyer, error)
	GetLawyerByID(id string) (*domain.Lawyer, error)
	CreateLawyerBooking(booking *domain.LawyerBooking) error
	GetCurrencyRates(from, to string) ([]*domain.CurrencyRate, error)
}

type TrustRepository interface {
	AddReview(review *domain.TrustReview) error
	GetUserReviews(userID string) ([]*domain.TrustReview, error)
	GetAverageRating(subjectID string, subjectType string) (float64, error)
}

type TransactionRepository interface {
	Create(tx *domain.Transaction) error
	GetByID(id string) (*domain.Transaction, error)
	GetByUser(userID string, limit, offset int) ([]*domain.Transaction, error)
	UpdateStatus(txID, status string) error
}
