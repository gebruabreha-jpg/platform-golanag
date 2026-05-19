package postgres

import (
	"connectme/internal/domain"
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

	// Auto-migrate domain models
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
