package domain

import (
	"context"
	"errors"
	"time"

	"github.com/SamuelsSantos/product-discount-service/products/domain/pb"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/patrickmn/go-cache"
)

//ProductService  interface
type ProductService struct {
	repo  ProductRepository
	cache *cache.Cache
}

//NewProductService create new service
func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{
		repo:  repo,
		cache: cache.New(60*time.Second, 70*time.Second),
	}
}

// GetByID fetch product by ID
func (s *ProductService) GetByID(ctx context.Context, r *pb.RequestProduct) (*pb.Product, error) {
	id := r.GetId()

	productCache, found := s.cache.Get(id)
	if found {
		return productCache.(*pb.Product), nil
	}

	product, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("Not found")
	}

	s.cache.Set(product.GetId(), product, cache.DefaultExpiration)
	return product, nil
}

// List fetch products
func (s *ProductService) List(r *empty.Empty, stream pb.ProductService_ListServer) error {

	products, err := s.repo.List()
	if err != nil {
		return err
	}

	for _, product := range products {
		err := stream.Send(product)
		if err != nil {
			return err
		}
	}

	return nil
}
