package domain

import "database/sql"

// ProductRepository is a contract to manage data
type ProductRepository interface {
	GetByID(id string) (*sql.Rows, error)
	List() (*sql.Rows, error)
	Close() error
}
