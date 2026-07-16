package service

import (
	"context"
	"errors"

	"task-manager-api/internal/model"
	"task-manager-api/internal/repository"
)

var ErrEmailExists = errors.New("email already registered")

type UserService struct {
	users repository.UserRepository
}

func NewUserService(users repository.UserRepository) *UserService {
	return &UserService{users: users}
}

func (s *UserService) Register(ctx context.Context, email, passwordHash string) (model.User, error) {
	_, err := s.users.GetByEmail(ctx, email)
	if err == nil {
		return model.User{}, ErrEmailExists
	}
	if !errors.Is(err, repository.ErrUserNotFound) {
		return model.User{}, err
	}
	return s.users.Create(ctx, email, passwordHash)
}

func (s *UserService) Login(ctx context.Context, email, passwordHash string) (*model.User, error) {
	user, err := s.users.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetByID(ctx context.Context, id int) (*model.User, error) {
	return s.users.GetByID(ctx, id)
}
