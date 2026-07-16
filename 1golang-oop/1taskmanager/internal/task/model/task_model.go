package model

import "time"

type Task struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Done        bool      `json:"done" db:"done"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
