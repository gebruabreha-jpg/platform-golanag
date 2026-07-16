package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"task-manager-api/internal/task/model"

	_ "github.com/lib/pq"
)

// schema is applied on startup so the API is self-contained.
var schema = `
CREATE TABLE IF NOT EXISTS tasks (
    id          SERIAL PRIMARY KEY,
    title       TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    done        BOOLEAN NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);
`

// PostgresTaskRepository implements TaskRepository on top of PostgreSQL.
type PostgresTaskRepository struct {
	db *sql.DB
}

// NewPostgresTaskRepository opens a connection pool and ensures the schema exists.
func NewPostgresTaskRepository(databaseURL string) (*PostgresTaskRepository, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}
	if _, err := db.Exec(schema); err != nil {
		return nil, fmt.Errorf("create schema: %w", err)
	}
	return &PostgresTaskRepository{db: db}, nil
}

// Close releases the connection pool.
func (r *PostgresTaskRepository) Close() error {
	return r.db.Close()
}

func (r *PostgresTaskRepository) Add(title, description string) model.Task {
	var t model.Task
	err := r.db.QueryRow(
		`INSERT INTO tasks (title, description) VALUES ($1, $2)
         RETURNING id, title, description, done, created_at`,
		title, description,
	).Scan(&t.ID, &t.Title, &t.Description, &t.Done, &t.CreatedAt)
	if err != nil {
		return model.Task{}
	}
	return t
}

func (r *PostgresTaskRepository) Get(id int) *model.Task {
	var t model.Task
	err := r.db.QueryRow(
		`SELECT id, title, description, done, created_at
         FROM tasks WHERE id = $1`,
		id,
	).Scan(&t.ID, &t.Title, &t.Description, &t.Done, &t.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return nil
	}
	return &t
}

func (r *PostgresTaskRepository) List() []model.Task {
	rows, err := r.db.Query(
		`SELECT id, title, description, done, created_at FROM tasks ORDER BY id`,
	)
	if err != nil {
		return []model.Task{}
	}
	defer rows.Close()

	out := []model.Task{}
	for rows.Next() {
		var t model.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Done, &t.CreatedAt); err != nil {
			return out
		}
		out = append(out, t)
	}
	return out
}

func (r *PostgresTaskRepository) Update(id int, title *string, done *bool) *model.Task {
	var (
		currentTitle string
		currentDone  bool
	)
	err := r.db.QueryRow(
		`SELECT title, done FROM tasks WHERE id = $1`, id,
	).Scan(&currentTitle, &currentDone)
	if err != nil {
		return nil
	}

	newTitle := currentTitle
	newDone := currentDone
	if title != nil {
		newTitle = *title
	}
	if done != nil {
		newDone = *done
	}

	res, err := r.db.Exec(
		`UPDATE tasks SET title = $1, done = $2 WHERE id = $3`,
		newTitle, newDone, id,
	)
	if err != nil {
		return nil
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return nil
	}

	t := model.Task{ID: id, Title: newTitle, Done: newDone}
	_ = r.db.QueryRow(
		`SELECT description, created_at FROM tasks WHERE id = $1`, id,
	).Scan(&t.Description, &t.CreatedAt)
	return &t
}

func (r *PostgresTaskRepository) Delete(id int) bool {
	res, err := r.db.Exec(`DELETE FROM tasks WHERE id = $1`, id)
	if err != nil {
		return false
	}
	n, _ := res.RowsAffected()
	return n > 0
}
