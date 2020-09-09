package domain

import "database/sql"

// UserRepository is a contract to manage data
type UserRepository interface {
	GetByID(id string) (*sql.Rows, error)
	Close() error
}
