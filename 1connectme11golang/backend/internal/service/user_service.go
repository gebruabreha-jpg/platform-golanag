package service

import (
	"connectme/internal/domain"
	"connectme/internal/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

type RegisterRequest struct {
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=8"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	Phone        string `json:"phone"`
	Country      string `json:"country"`
	Role         string `json:"role"` // DIASPORA, LOCAL, MERCHANT
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID              string    `json:"id"`
	Email           string    `json:"email"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Phone           string    `json:"phone,omitempty"`
	AvatarURL       string    `json:"avatar_url,omitempty"`
	Location        string    `json:"location,omitempty"`
	Country         string    `json:"country,omitempty"`
	Role            string    `json:"role"`
	IsVerified      bool      `json:"is_verified"`
	VerificationLevel int     `json:"verification_level"`
	TrustScore      float64   `json:"trust_score"`
	CreatedAt       time.Time `json:"created_at"`
}

func (s *UserService) Register(req RegisterRequest) (*domain.User, error) {
	// Check if user already exists
	existing, _ := s.userRepo.GetByEmail(req.Email)
	if existing != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:       req.Email,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Phone:       req.Phone,
		Country:     req.Country,
		Role:        req.Role,
		IsVerified:  false,
		TrustScore:  0.5, // Initial trust score
	}
	// Store hashed password separately in real implementation
	_ = string(hashedPassword)

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(req LoginRequest) (*domain.User, error) {
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil || user == nil {
		return nil, errors.New("invalid credentials")
	}

	// Verify password (implement proper password hash check)
	// if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
	//     return nil, errors.New("invalid credentials")
	// }

	return user, nil
}

func (s *UserService) GetProfile(userID string) (*UserResponse, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	return &UserResponse{
		ID:              user.ID,
		Email:           user.Email,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Phone:           user.Phone,
		AvatarURL:       user.AvatarURL,
		Location:        user.Location,
		Country:         user.Country,
		Role:            user.Role,
		IsVerified:      user.IsVerified,
		VerificationLevel: user.VerificationLevel,
		TrustScore:      user.TrustScore,
		CreatedAt:       user.CreatedAt,
	}, nil
}

func (s *UserService) UpdateTrustScore(userID string, newScore float64) error {
	if newScore < 0 || newScore > 5 {
		return errors.New("trust score must be between 0 and 5")
	}
	return s.userRepo.UpdateTrustScore(userID, newScore)
}
