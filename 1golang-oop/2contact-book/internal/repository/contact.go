package repository

import (
	"context"
	"database/sql"
	"errors"

	"contact-book-api/internal/model"
	"github.com/lib/pq"
)

// ContactRepository defines the contract for contact data operations.
// This interface allows us to swap implementations (Postgres, mock for tests)
// without changing the service or handler code.
type ContactRepository interface {
	Create(ctx context.Context, firstName, lastName, email, phone string) (model.Contact, error)
	GetByID(ctx context.Context, id int) (*model.Contact, error)
	GetByEmail(ctx context.Context, email string) (*model.Contact, error)
	GetAll(ctx context.Context) ([]model.Contact, error)
	Update(ctx context.Context, id int, firstName, lastName, email, phone string) (*model.Contact, error)
	Delete(ctx context.Context, id int) error
}

// Sentinel errors for repository-layer failure modes.
// These are checked by the service and handler to return the correct HTTP status.
var (
	ErrContactNotFound = errors.New("contact not found")
	ErrEmailExists     = errors.New("email already exists")
)

// PostgresContactRepository implements ContactRepository using PostgreSQL.
// It holds a *sql.DB connection pool and uses it to run queries.
type PostgresContactRepository struct {
	db *sql.DB
}

// NewPostgresContactRepository injects the database connection.
// Called in main.go: repository.NewPostgresContactRepository(db)
func NewPostgresContactRepository(db *sql.DB) *PostgresContactRepository {
	return &PostgresContactRepository{db: db}
}

// Create inserts a new contact and returns the created record with generated ID.
// Uses INSERT ... RETURNING to get the auto-generated id and timestamps in one round trip.
func (r *PostgresContactRepository) Create(ctx context.Context, firstName, lastName, email, phone string) (model.Contact, error) {
	var c model.Contact
	row := r.db.QueryRowContext(ctx,
		`INSERT INTO contacts (first_name, last_name, email, phone, created_at) VALUES ($1,$2,$3,$4,now()) RETURNING id, first_name, last_name, email, phone, created_at`,
		firstName, lastName, email, phone,
	)
	if err := row.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email, &c.Phone, &c.CreatedAt); err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return model.Contact{}, ErrEmailExists
		}
		return model.Contact{}, err
	}
	return c, nil
}

// GetByID retrieves a contact by primary key.
// Returns ErrContactNotFound if no row matches the id.
func (r *PostgresContactRepository) GetByID(ctx context.Context, id int) (*model.Contact, error) {
	row := r.db.QueryRowContext(ctx,
		`SELECT id, first_name, last_name, email, phone, created_at FROM contacts WHERE id = $1`, id,
	)
	var c model.Contact
	if err := row.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email, &c.Phone, &c.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrContactNotFound
		}
		return nil, err
	}
	return &c, nil
}

// GetByEmail retrieves a contact by email address (case-insensitive).
// Uses LOWER() in SQL so "John@Example.com" matches "john@example.com".
func (r *PostgresContactRepository) GetByEmail(ctx context.Context, email string) (*model.Contact, error) {
	row := r.db.QueryRowContext(ctx,
		`SELECT id, first_name, last_name, email, phone, created_at FROM contacts WHERE LOWER(email) = LOWER($1)`, email,
	)
	var c model.Contact
	if err := row.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email, &c.Phone, &c.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrContactNotFound
		}
		return nil, err
	}
	return &c, nil
}

// GetAll returns all contacts ordered by ID.
// Uses QueryContext (not QueryRowContext) because it returns multiple rows.
// defer rows.Close() ensures the result set is freed even if we return early from the loop.
func (r *PostgresContactRepository) GetAll(ctx context.Context) ([]model.Contact, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT id, first_name, last_name, email, phone, created_at FROM contacts ORDER BY id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []model.Contact
	for rows.Next() {
		var c model.Contact
		if err := rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email, &c.Phone, &c.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	// rows.Err() catches any error that occurred during iteration (e.g. network failure).
	return out, rows.Err()
}

// Update modifies an existing contact by ID.
// Uses UPDATE ... RETURNING to get the updated row in one round trip.
// Returns ErrContactNotFound if no row matches the id.
// Returns ErrEmailExists if the new email violates the unique constraint.
func (r *PostgresContactRepository) Update(ctx context.Context, id int, firstName, lastName, email, phone string) (*model.Contact, error) {
	var c model.Contact
	row := r.db.QueryRowContext(ctx,
		`UPDATE contacts SET first_name=$1, last_name=$2, email=$3, phone=$4 WHERE id=$5 RETURNING id, first_name, last_name, email, phone, created_at`,
		firstName, lastName, email, phone, id,
	)
	if err := row.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email, &c.Phone, &c.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrContactNotFound
		}
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, ErrEmailExists
		}
		return nil, err
	}
	return &c, nil
}

// Delete removes a contact by ID.
// Returns ErrContactNotFound if no row was deleted (id doesn't exist).
func (r *PostgresContactRepository) Delete(ctx context.Context, id int) error {
	result, err := r.db.ExecContext(ctx, `DELETE FROM contacts WHERE id = $1`, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrContactNotFound
	}
	return nil
}
