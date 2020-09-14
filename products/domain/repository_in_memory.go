package domain

import (
	"database/sql"
	"errors"

	"github.com/SamuelsSantos/product-discount-service/products/domain/pb"
	_ "github.com/mattn/go-sqlite3"
)

//InMemoryRepo repository
type InMemoryRepo struct {
	data map[string]*pb.Product
}

//NewInMemoryRepository create new repository
func NewInMemoryRepository() *InMemoryRepo {
	data := make(map[string]*pb.Product, 2)

	data["1"] = &pb.Product{
		Id:           "1",
		Description:  "Produto Teste 01",
		PriceInCents: 10090,
		Title:        "Produto 01"}

	data["2"] = &pb.Product{
		Id:           "2",
		Description:  "Produto Teste 02",
		PriceInCents: 10090,
		Title:        "Produto 02"}

	return &InMemoryRepo{data}
}

//GetDB database connection
func (r *InMemoryRepo) GetDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic("failed to create in-memory SQLite database")
	}
	return db, nil
}

//Close database connection
func (r *InMemoryRepo) Close() error {
	r.data = make(map[string]*pb.Product, 0)
	return nil
}

// GetByID fetch product by ID
func (r *InMemoryRepo) GetByID(id string) (*pb.Product, error) {

	product := r.data[id]
	if product == nil {
		return nil, errors.New("Not found")
	}

	return product, nil
}

// List ...
func (r *InMemoryRepo) List() ([]*pb.Product, error) {

	var products []*pb.Product

	for _, product := range r.data {
		products = append(products, product)
	}

	return products, nil
}
