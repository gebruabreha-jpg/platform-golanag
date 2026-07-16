//Interface + Postgres SQL queries
package repository
import (
	"context"
	"errors"
	"sync"
	"time"
	"contact-book-api/internal/model"
	"github.com/lib/pq"

)

//interface for all mtheods
type ContactRepository interface {
    Create(ctx context.Context, firstName, lastName, email, phone string) (model.Contact, error)
    GetByID(ctx context.Context, id int) (*model.Contact, error)
    GetByEmail(ctx context.Context, email string) (*model.Contact, error)
    GetAll(ctx context.Context) ([]model.Contact, error)
    Update(ctx context.Context, id int, firstName, lastName, email, phone string) (*model.Contact, error)
    Delete(ctx context.Context, id int) error
}

//custom error
var ErrContactNotFound = errors.New("contact not found")
var ErrEmailExists = errors.New("email already exists")

//Struct + Constructor
//This is the repository wrapper — it holds the database connection so all query methods can use it.
//With this struct, you inject the DB once at startup, then all methods share it:
// 1. Struct = "this object owns a DB connection"
type PostgresContactRepository struct {
    db *sql.DB   // unexported field: only this package can use it
}

// 2. Constructor = "give me a DB, I'll give you a ready repository"
func NewPostgresContactRepository(db *sql.DB) *PostgresContactRepository {
    return &PostgresContactRepository{db: db}  // store connection inside the struct
}


  //main.go
  //└─ sql.Open("postgres", dsn)  →  *sql.DB
  //        ↓
  // NewPostgresContactRepository(db)  →  repo.db = db
  //       ↓
  //  repo.Create(...) uses repo.db internally


//creat, update, get, delete and  other functions used to manipulate DB tables.
// Create inserts a new contact and returns the created record with generated ID.
func (r *PostgresContactRepository) Create(ctx context.Context, firstName, lastName, email, phone string) (model.Contact, error) {
	var c model.Contact
	row := r.db.QueryRowContext(ctx,
		`INSERT INTO contacts (first_name, last_name, email, phone, created_at)
		 VALUES ($1, $2, $3, $4, now())
		 RETURNING id, first_name, last_name, email, phone, created_at`,
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
func (r *PostgresContactRepository) GetByID(ctx context.Context, id int) (*model.Contact, error) {
	row := r.db.QueryRowContext(ctx,
		`SELECT id, first_name, last_name, email, phone, created_at
		 FROM contacts
		 WHERE id = $1`,
		id,
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

// GetByEmail retrieves a contact by email address.
func (r *PostgresContactRepository) GetByEmail(ctx context.Context, email string) (*model.Contact, error) {
	row := r.db.QueryRowContext(ctx,
		`SELECT id, first_name, last_name, email, phone, created_at
		 FROM contacts
		 WHERE email = $1`,
		email,
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
func (r *PostgresContactRepository) GetAll(ctx context.Context) ([]model.Contact, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT id, first_name, last_name, email, phone, created_at
		 FROM contacts
		 ORDER BY id`,
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
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

// Update modifies an existing contact by ID.
func (r *PostgresContactRepository) Update(ctx context.Context, id int, firstName, lastName, email, phone string) (*model.Contact, error) {
	var c model.Contact
	row := r.db.QueryRowContext(ctx,
		`UPDATE contacts
		 SET first_name = $1, last_name = $2, email = $3, phone = $4
		 WHERE id = $5
		 RETURNING id, first_name, last_name, email, phone, created_at`,
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