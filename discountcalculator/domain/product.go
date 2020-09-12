package domain

import (
	"context"
	"errors"

	"github.com/SamuelsSantos/product-discount-service/discountcalculator/domain/pb"
	"google.golang.org/grpc"
)

// ProductService ...
type ProductService struct {
	host string
}

// NewProductService create a new instance ProductService
func NewProductService(host string) *ProductService {
	return &ProductService{host: host}
}

// GetProductByID fetch product by ID
func (s *ProductService) GetProductByID(id string) (*pb.Product, error) {
	return getProductByID(s.host, id)
}

func getProductByID(host string, id string) (*pb.Product, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, errors.New("Failed to fetch data from server")
	}
	defer conn.Close()

	product, err := pb.NewProductServiceClient(conn).GetByID(context.Background(), &pb.RequestProduct{Id: id})
	if err != nil {
		return nil, errors.New("Failed to connect to server")
	}

	return product, nil
}
