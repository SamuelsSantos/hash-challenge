package domain

import (
	"database/sql"

	"github.com/SamuelsSantos/product-discount-service/products/domain/pb"
)

// ProductRepository is a contract to manage data
type ProductRepository interface {
	GetByID(id string) (*pb.Product, error)
	List() ([]*pb.Product, error)
	GetDB() (*sql.DB, error)
}
