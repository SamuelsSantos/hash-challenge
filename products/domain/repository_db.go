package domain

import (
	"database/sql"
	"errors"
	"log"

	"github.com/SamuelsSantos/product-discount-service/products/config"
	"github.com/SamuelsSantos/product-discount-service/products/domain/pb"
)

//SQLRepo repository
type SQLRepo struct {
	Cfg *config.Config
}

//NewSQLRepository create new repository
func NewSQLRepository(cfg *config.Config) *SQLRepo {
	return &SQLRepo{cfg}
}

func newDBConnection(cfg *config.Config) (*sql.DB, error) {
	return sql.Open(cfg.Db.Driver, cfg.Db.ToURL())
}

// GetDB new db connection
func (r *SQLRepo) GetDB() (*sql.DB, error) {
	return newDBConnection(r.Cfg)
}

// GetByID fetch product by ID
func (r *SQLRepo) GetByID(id string) (*pb.Product, error) {

	db, err := r.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`select id, title, description, price_in_cents from public.products where id = $1`)
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

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

	db, err := r.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`select id, title, description, price_in_cents from public.products`)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
