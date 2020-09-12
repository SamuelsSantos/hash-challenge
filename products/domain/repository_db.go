package domain

import (
	"database/sql"
	"errors"
	"log"

	"github.com/SamuelsSantos/product-discount-service/products/domain/pb"
)

//SQLRepo repository
type SQLRepo struct {
	db *sql.DB
}

//NewRepository create new repository
func NewRepository(db *sql.DB) *SQLRepo {
	return &SQLRepo{db}
}

//Close database connection
func (r *SQLRepo) Close() error {
	return r.db.Close()
}

// GetByID fetch product by ID
func (r *SQLRepo) GetByID(id string) (*pb.Product, error) {

	stmt, err := r.db.Prepare(`select id, title, description, price_in_cents from public.products where id = $1`)
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		pbProduct, err := transform(rows)
		if err != nil {
			return nil, err
		}

		return pbProduct, nil
	}

	return nil, errors.New("Not found")
}

// List all products
func (r *SQLRepo) List() ([]*pb.Product, error) {

	rows, err := r.db.Query(`select id, title, description, price_in_cents from public.products`)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	result := make([]*pb.Product, 0)
	for rows.Next() {
		pbProduct, err := transform(rows)
		if err != nil {
			return nil, err
		}

		result = append(result, pbProduct)
	}

	return result, nil
}

func transform(r *sql.Rows) (*pb.Product, error) {

	var id string
	var title string
	var description string
	var priceInCents int64

	err := r.Scan(&id, &title, &description, &priceInCents)
	if err != nil {
		return nil, err
	}

	return &pb.Product{
		Id:           id,
		Title:        title,
		Description:  description,
		PriceInCents: priceInCents,
	}, nil
}
