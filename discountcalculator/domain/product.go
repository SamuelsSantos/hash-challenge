package domain

import (
	"context"
	"time"

	"github.com/SamuelsSantos/product-discount-service/discountcalculator/domain/pb"
	"google.golang.org/grpc"
)

// ProductClient definition
type ProductClient struct {
	host string
}

// ProductService definition
type ProductService interface {
	GetProductByID(id string) (*pb.Product, error)
}

// NewProductClient create a new instance ProductClient
func NewProductClient(host string) *ProductClient {
	return &ProductClient{host: host}
}

// GetProductByID fetch product by ID
func (s *ProductClient) GetProductByID(id string) (*pb.Product, error) {
	return getProductByID(s.host, id)
}

func getProductByID(host string, id string) (*pb.Product, error) {

	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	product, err := client.GetByID(ctx, &pb.RequestProduct{Id: id})
	if err != nil {
		return nil, err
	}

	return product, nil
}
