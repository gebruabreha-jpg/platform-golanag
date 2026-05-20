package domain

import (
	"time"
)

type User struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	Email        string    `json:"email" gorm:"uniqueIndex"`
	Phone        string    `json:"phone" gorm:"uniqueIndex"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	Bio          string    `json:"bio" gorm:"type:text"`
	AvatarURL    string    `json:"avatar_url"`
	Location     string    `json:"location"`
	Country      string    `json:"country"`
	City         string    `json:"city"`
	Role         string    `json:"role"` // DIASPORA, LOCAL, MERCHANT, ADMIN
	IsVerified   bool      `json:"is_verified"`
	VerificationLevel int `json:"verification_level"` // 0=none, 1=basic, 2=full
	TrustScore   float64   `json:"trust_score" gorm:"default:0"`
	TotalTransactions int `json:"total_transactions" gorm:"default:0"`
	PasswordHash string    `json:"-" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

type Community struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description" gorm:"type:text"`
	Category    string    `json:"category"` // SHIPPING, HOUSING, MARKETPLACE, JOBS, SCHOLARSHIPS, BUSINESS
	Location    string    `json:"location"` // City/Region
	Country     string    `json:"country"`
	IsPrivate   bool      `json:"is_private"`
	MemberCount int       `json:"member_count" gorm:"default:0"`
	ModeratorID string    `json:"moderator_id"`
	Rules       string    `json:"rules" gorm:"type:text"`
	ImageURL    string    `json:"image_url"`
	Tags        string    `json:"tags" gorm:"type:jsonb"` // JSON array
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Post struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	CommunityID string   `json:"community_id"`
	UserID     string    `json:"user_id"`
	Type       string    `json:"type"` // OFFER, REQUEST, INFO, DISCUSSION
	Title      string    `json:"title"`
	Content    string    `json:"content" gorm:"type:text"`
	MediaURLs  string    `json:"media_urls" gorm:"type:jsonb"`
	IsPinned   bool      `json:"is_pinned" gorm:"default:false"`
	IsClosed   bool      `json:"is_closed" gorm:"default:false"`
	ReplyCount int       `json:"reply_count" gorm:"default:0"`
	ViewCount  int       `json:"view_count" gorm:"default:0"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type MarketplaceItem struct {
	ID            string    `json:"id" gorm:"primaryKey"`
	SellerID      string    `json:"seller_id"`
	Title         string    `json:"title"`
	Description   string    `json:"description" gorm:"type:text"`
	Category      string    `json:"category"` // ELECTRONICS, CLOTHING, VEHICLES, HOUSEHOLD, etc.
	Subcategory   string    `json:"subcategory"`
	Price         float64   `json:"price"`
	Currency      string    `json:"currency" gorm:"default:'USD'"`
	Condition     string    `json:"condition"` // NEW, LIKE_NEW, GOOD, FAIR
	Location      string    `json:"location"`
	Country       string    `json:"country"`
	ShippingAvailable bool `json:"shipping_available" gorm:"default:false"`
	ShippingCost  float64   `json:"shipping_cost" gorm:"default:0"`
	ImageURLs     string    `json:"image_urls" gorm:"type:jsonb"`
	IsActive      bool      `json:"is_active" gorm:"default:true"`
	IsSold        bool      `json:"is_sold" gorm:"default:false"`
	ViewCount     int       `json:"view_count" gorm:"default:0"`
	InterestCount int       `json:"interest_count" gorm:"default:0"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type HousingListing struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	LandlordID   string    `json:"landlord_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description" gorm:"type:text"`
	PropertyType string    `json:"property_type"` // APARTMENT, HOUSE, ROOM
	RoomType     string    `json:"room_type"` // PRIVATE, SHARED, MASTER
	Bedrooms     int       `json:"bedrooms"`
	Bathrooms    int       `json:"bathrooms"`
	MonthlyRent  float64   `json:"monthly_rent"`
	Currency     string    `json:"currency" gorm:"default:'USD'"`
	Deposit      float64   `json:"deposit"`
	Address      string    `json:"address"`
	City         string    `json:"city"`
	Country      string    `json:"country"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	AvailableFrom time.Time `json:"available_from"`
	LeaseTerm    string    `json:"lease_term"` // MONTHLY, SHORT_TERM, LONG_TERM
	Furnished    bool      `json:"furnishished" gorm:"default:false"`
	IncludesUtilities bool `json:"includes_utilities" gorm:"default:false"`
	ImageURLs    string    `json:"image_urls" gorm:"type:jsonb"`
	IsActive     bool      `json:"is_active" gorm:"default:true"`
	ApplicationCount int `json:"application_count" gorm:"default:0"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Scholarship struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	Title        string    `json:"title"`
	Description  string    `json:"description" gorm:"type:text"`
	Provider     string    `json:"provider"`
	ProviderType string    `json:"provider_type"` // GOVERNMENT, UNIVERSITY, FOUNDATION
	Country      string    `json:"country"`
	City         string    `json:"city"`
	Level        string    `json:"level"` // UNDERGRADUATE, GRADUATE, PhD, POSTDOC
	Field        string    `json:"field"` // STEM, ARTS, BUSINESS, etc.
	Amount       float64   `json:"amount"`
	Currency     string    `json:"currency"`
	Covers       string    `json:"covers" gorm:"type:jsonb"` // ["TUITION", "LIVING", "TRAVEL"]
	Deadline     *time.Time `json:"deadline"`
	Eligibility  string    `json:"eligibility" gorm:"type:text"`
	Requirements string    `json:"requirements" gorm:"type:text"`
	ApplicationURL string  `json:"application_url"`
	IsActive     bool      `json:"is_active" gorm:"default:true"`
	IsFeatured   bool      `json:"is_featured" gorm:"default:false"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Job struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	EmployerID  string    `json:"employer_id"`
	Title       string    `json:"title"`
	Description string    `json:"description" gorm:"type:text"`
	JobType     string    `json:"job_type"` // FULL_TIME, PART_TIME, CONTRACT, FREELANCE
	Remote      bool      `json:"remote"`
	Location    string    `json:"location"`
	Country     string    `json:"country"`
	SalaryMin   float64   `json:"salary_min"`
	SalaryMax   float64   `json:"salary_max"`
	Currency    string    `json:"currency"`
	Industry    string    `json:"industry"`
	Skills      string    `json:"skills" gorm:"type:jsonb"`
	Benefits    string    `json:"benefits" gorm:"type:jsonb"`
	ApplicationURL string `json:"application_url"`
	ExpiresAt   *time.Time `json:"expires_at"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	ViewCount   int       `json:"view_count" gorm:"default:0"`
	ApplicationCount int `json:"application_count" gorm:"default:0"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CurrencyRate struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	FromCurrency string `json:"from_currency"`
	ToCurrency   string `json:"to_currency"`
	Rate      float64   `json:"rate"`
	Source    string    `json:"source"` // EXCHANGE_RATE_API, WISE, REMITLY
	UpdatedAt time.Time `json:"updated_at"`
}

type TrustReview struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	ReviewerID string    `json:"reviewer_id"`
	SubjectID  string    `json:"subject_id"`
	SubjectType string   `json:"subject_type"` // USER, LISTING, COMMUNITY
	Rating     int       `json:"rating"` // 1-5
	Comment    string    `json:"comment" gorm:"type:text"`
	IsVerified bool      `json:"is_verified"` // Verified transaction
	CreatedAt  time.Time `json:"created_at"`
}

type Transaction struct {
	ID            string    `json:"id" gorm:"primaryKey"`
	BuyerID       string    `json:"buyer_id"`
	SellerID      string    `json:"seller_id"`
	ItemID        *string   `json:"item_id,omitempty"`
	Type          string    `json:"type"` // MARKETPLACE_PURCHASE, SERVICE_PAYMENT, DONATION
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Status        string    `json:"status"` // PENDING, COMPLETED, CANCELLED, DISPUTED
	PaymentMethod string    `json:"payment_method"`
	EscrowID      string    `json:"escrow_id,omitempty"`
	ReleasedAt    *time.Time `json:"released_at"`
	CompletedAt  *time.Time `json:"completed_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Lawyer struct {
	ID             string    `json:"id" gorm:"primaryKey"`
	UserID         string    `json:"user_id"`
	Specialization string    `json:"specialization"` // IMMIGRATION, BUSINESS, FAMILY, PROPERTY, CRIMINAL
	Name           string    `json:"name"`
	Firm           string    `json:"firm"`
	YearsExperience int      `json:"years_experience"`
	Location       string    `json:"location"`
	Country        string    `json:"country"`
	ConsultationFee float64  `json:"consultation_fee"`
	Currency       string    `json:"currency" gorm:"default:'USD'"`
	IsVerified     bool      `json:"is_verified" gorm:"default:false"`
	Rating         float64   `json:"rating" gorm:"default:0"`
	ReviewCount    int       `json:"review_count" gorm:"default:0"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type LawyerBooking struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	LawyerID     string    `json:"lawyer_id"`
	UserID       string    `json:"user_id"`
	ScheduledAt  time.Time `json:"scheduled_at"`
	Duration     int       `json:"duration"` // minutes
	Type         string    `json:"type"` // IN_PERSON, REMOTE
	Status       string    `json:"status"` // PENDING, CONFIRMED, COMPLETED, CANCELLED
	Notes        string    `json:"notes" gorm:"type:text"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
