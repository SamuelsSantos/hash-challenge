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

// GetByID fetch user by ID
func (r *SQLRepo) GetByID(id string) (*sql.Rows, error) {
	stmt, err := r.db.Prepare(`select id, first_name, last_name, date_of_birth from public.user where id = $1`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

//Close database connection
func (r *SQLRepo) Close() error {
	return r.db.Close()
}
