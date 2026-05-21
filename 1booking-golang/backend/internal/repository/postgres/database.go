package postgres

import (
	"connectme/internal/domain"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func NewPostgresDB(databaseURL string) (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = DB.AutoMigrate(
		&domain.User{},
		&domain.Community{},
		&domain.Post{},
		&domain.MarketplaceItem{},
		&domain.HousingListing{},
		&domain.Scholarship{},
		&domain.Job{},
		&domain.CurrencyRate{},
		&domain.TrustReview{},
		&domain.Transaction{},
		&domain.Lawyer{},
		&domain.LawyerBooking{},
	)
	if err != nil {
		return nil, err
	}

	log.Println("Database migration completed successfully")
	return DB, nil
}

// --- Factory functions called by main.go (return concrete types that
//     satisfy the interfaces defined in repository/repository.go) ---

func NewUserRepository(db *gorm.DB) *UserRepository                                          { return &UserRepository{db: db} }
func NewCommunityRepository(db *gorm.DB) *CommunityRepository                                 { return &CommunityRepository{db: db} }
func NewPostRepository(db *gorm.DB) *PostRepository                                           { return &PostRepository{db: db} }
func NewMarketplaceRepository(db *gorm.DB) *MarketplaceRepository                             { return &MarketplaceRepository{db: db} }
func NewHousingRepository(db *gorm.DB) *HousingRepository                                     { return &HousingRepository{db: db} }
func NewServiceRepository(db *gorm.DB) *ServiceRepository                                     { return &ServiceRepository{db: db} }

// NewAuthRepository placeholder — extend domain with AuthToken entity if needed
func NewAuthRepository(db *gorm.DB, jwtSecret string) *AuthRepository {
	return &AuthRepository{db: db, jwtSecret: jwtSecret, uuid: uuid.NewString}
}

// AuthRepository is a sub-layer for token/blacklist operations.
type AuthRepository struct {
	db        *gorm.DB
	jwtSecret string
	uuid      func() string
}
