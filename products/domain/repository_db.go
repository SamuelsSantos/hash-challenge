package domain

import (
	"database/sql"
)

//SQLRepo repository
type SQLRepo struct {
	db *sql.DB
}

//NewRepository create new repository
func NewRepository(db *sql.DB) *SQLRepo {
	return &SQLRepo{db}
}

// GetByID fetch product by ID
func (r *SQLRepo) GetByID(id string) (*sql.Rows, error) {
	stmt, err := r.db.Prepare(`select id, title, description, price_in_cents from public.products where id = $1`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// List all products
func (r *SQLRepo) List() (*sql.Rows, error) {
	rows, err := r.db.Query(`select id, title, description, price_in_cents from public.products`)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

//Close database connection
func (r *SQLRepo) Close() error {
	return r.db.Close()
}
