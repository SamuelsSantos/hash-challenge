package domain

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/SamuelsSantos/product-discount-service/products/config"
	"github.com/SamuelsSantos/product-discount-service/products/domain/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// var c *cache.Cache

//ProductService  interface
type ProductService struct {
	repo *SQLRepo
}

//NewService create new service
func NewService(cfg *config.Config) *ProductService {

	dB, err := sql.Open(cfg.Db.Driver, cfg.Db.ToURL())
	if err != nil {
		panic(err)
	}

	return &ProductService{
		repo: NewRepository(dB),
	}
}

func getProduct(s *ProductService, id string) (*pb.Response, error) {
	rows, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("could not find Product with id: %v", id)
	}

	var product *pb.Product

	for rows.Next() {
		var id string
		var title string
		var description string
		var priceInCents int64

		err = rows.Scan(&id, &title, &description, &priceInCents)
		if err != nil {
			return nil, err
		}

		pbProduct := pb.Product{
			Id:           id,
			Title:        title,
			Description:  description,
			PriceInCents: priceInCents,
		}

		product = &pbProduct
	}

	return &pb.Response{
		Result: product,
	}, nil
}

// GetByID fetch product by ID
func (s *ProductService) GetByID(ctx context.Context, r *pb.Request) (*pb.Response, error) {
	id := r.GetId()
	return getProduct(s, id)
}

// List fetch products
func (s *ProductService) List(r *pb.Empty, list pb.ProductService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
