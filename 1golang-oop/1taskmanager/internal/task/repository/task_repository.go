package repository

import (
	"sync"

	"task-manager-api/internal/task/model"
)

// TaskRepository defines storage operations for tasks.
// Any backend (in-memory, Postgres, etc.) can implement this.
type TaskRepository interface {
	Add(title, description string) model.Task
	Get(id int) *model.Task
	List() []model.Task
	Update(id int, title *string, done *bool) *model.Task
	Delete(id int) bool
}

// InMemoryTaskRepository is a thread-safe in-memory implementation.
type InMemoryTaskRepository struct {
	mu      sync.Mutex
	tasks   []model.Task
	counter int
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks:   []model.Task{},
		counter: 1,
	}
}

func (r *InMemoryTaskRepository) Add(title, description string) model.Task {
	r.mu.Lock()
	defer r.mu.Unlock()
	task := model.Task{ID: r.counter, Title: title, Description: description, Done: false}
	r.counter++
	r.tasks = append(r.tasks, task)
	return task
}

func (r *InMemoryTaskRepository) Get(id int) *model.Task {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i := range r.tasks {
		if r.tasks[i].ID == id {
			t := r.tasks[i]
			return &t
		}
	}
	return nil
}

func (r *InMemoryTaskRepository) List() []model.Task {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := make([]model.Task, len(r.tasks))
	copy(out, r.tasks)
	return out
}

func (r *InMemoryTaskRepository) Update(id int, title *string, done *bool) *model.Task {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i := range r.tasks {
		if r.tasks[i].ID == id {
			if title != nil {
				r.tasks[i].Title = *title
			}
			if done != nil {
				r.tasks[i].Done = *done
			}
			t := r.tasks[i]
			return &t
		}
	}
	return nil
}

func (r *InMemoryTaskRepository) Delete(id int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, task := range r.tasks {
		if task.ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return true
		}
	}
	return false
}
