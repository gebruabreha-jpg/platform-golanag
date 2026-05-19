package service

import (
	"connectme/internal/domain"
	"connectme/internal/repository"
	"errors"
)

type CommunityService struct {
	communityRepo repository.CommunityRepository
	postRepo      repository.PostRepository
}

func NewCommunityService(
	communityRepo repository.CommunityRepository,
	postRepo repository.PostRepository,
) *CommunityService {
	return &CommunityService{
		communityRepo: communityRepo,
		postRepo:      postRepo,
	}
}

type CreateCommunityRequest struct {
	Name        string   `json:"name" validate:"required,min=3,max=100"`
	Description string   `json:"description" validate:"required"`
	Category    string   `json:"category" validate:"required"`
	Location    string   `json:"location"`
	Country     string   `json:"country"`
	IsPrivate   bool     `json:"is_private"`
	Rules       string   `json:"rules"`
	ImageURL    string   `json:"image_url"`
	Tags        []string `json:"tags"`
	ModeratorID string   `json:"moderator_id"`
}

func (s *CommunityService) CreateCommunity(req CreateCommunityRequest) (*domain.Community, error) {
	community := &domain.Community{
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,
		Location:    req.Location,
		Country:     req.Country,
		IsPrivate:   req.IsPrivate,
		Rules:       req.Rules,
		ImageURL:    req.ImageURL,
		ModeratorID: req.ModeratorID,
		MemberCount: 1, // Creator is first member
	}

	if err := s.communityRepo.Create(community); err != nil {
		return nil, err
	}

	return community, nil
}

type ListCommunitiesFilter struct {
	Category  string `json:"category"`
	Location  string `json:"location"`
	Country   string `json:"country"`
	IsPrivate bool   `json:"is_private"`
	Search    string `json:"search"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
}

func (s *CommunityService) ListCommunities(filter ListCommunitiesFilter) ([]*domain.Community, int, error) {
	communities, err := s.communityRepo.List(filter.Category, filter.Location, filter.Limit, filter.Offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.communityRepo.Count(filter) // Implement Count in repo
	if err != nil {
		return nil, 0, err
	}

	return communities, total, nil
}

func (s *CommunityService) GetCommunity(id string) (*domain.Community, error) {
	return s.communityRepo.GetByID(id)
}

func (s *CommunityService) JoinCommunity(communityID, userID string) error {
	// Check if community exists
	community, err := s.communityRepo.GetByID(communityID)
	if err != nil || community == nil {
		return errors.New("community not found")
	}

	// Check if user already member (implement in repo)
	// Add user to community members (many-to-many relationship)

	// Increment member count
	return s.communityRepo.IncrementMemberCount(communityID)
}

func (s *CommunityService) LeaveCommunity(communityID, userID string) error {
	// Remove user from members
	// Decrement member count
	return s.communityRepo.DecrementMemberCount(communityID)
}

type CreatePostRequest struct {
	CommunityID string   `json:"community_id" validate:"required"`
	UserID      string   `json:"user_id" validate:"required"`
	Type        string   `json:"type" validate:"required,oneof=OFFER REQUEST INFO DISCUSSION"`
	Title       string   `json:"title" validate:"required,min=5,max=200"`
	Content     string   `json:"content" validate:"required,min=10"`
	MediaURLs   []string `json:"media_urls"`
}

func (s *CommunityService) CreatePost(req CreatePostRequest) (*domain.Post, error) {
	// Check user permissions to post in community
	post := &domain.Post{
		CommunityID: req.CommunityID,
		UserID:      req.UserID,
		Type:        req.Type,
		Title:       req.Title,
		Content:     req.Content,
		MediaURLs:   req.MediaURLs,
	}

	if err := s.postRepo.Create(post); err != nil {
		return nil, err
	}

	return post, nil
}

func (s *CommunityService) GetPosts(communityID string, limit, offset int) ([]*domain.Post, error) {
	return s.postRepo.ListByCommunity(communityID, limit, offset)
}
