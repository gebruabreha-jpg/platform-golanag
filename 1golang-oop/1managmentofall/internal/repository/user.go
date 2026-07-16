package repository

import (
	"context"
	"errors"
	"sync"
	"time"

	"task-manager-api/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, email, passwordHash string) (model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetByID(ctx context.Context, id int) (*model.User, error)
}

var ErrUserNotFound = errors.New("user not found")

type InMemoryUserRepository struct {
	mu     sync.Mutex
	users  []model.User
	nextID int
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{nextID: 1}
}

func (r *InMemoryUserRepository) Create(ctx context.Context, email, passwordHash string) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	user := model.User{ID: r.nextID, Email: email, Password: passwordHash, CreatedAt: time.Now()}
	r.nextID++
	r.users = append(r.users, user)
	return user, nil
}

func (r *InMemoryUserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i := range r.users {
		if r.users[i].Email == email {
			u := r.users[i]
			return &u, nil
		}
	}
	return nil, ErrUserNotFound
}

func (r *InMemoryUserRepository) GetByID(ctx context.Context, id int) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i := range r.users {
		if r.users[i].ID == id {
			u := r.users[i]
			return &u, nil
		}
	}
	return nil, ErrUserNotFound
}
