package domain

import (
	"database/sql"

	"github.com/SamuelsSantos/product-discount-service/users/domain/pb"
)

// UserRepository is a contract to manage data
type UserRepository interface {
	GetByID(id string) (*pb.User, error)
	GetDB() (*sql.DB, error)
}
