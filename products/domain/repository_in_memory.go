package domain

import (
	"errors"

	"github.com/SamuelsSantos/product-discount-service/products/domain/pb"
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
