package domain

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/SamuelsSantos/product-discount-service/products/config"
	"github.com/SamuelsSantos/product-discount-service/products/domain/pb"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/patrickmn/go-cache"
)

var c *cache.Cache

func setCache() {

	c = cache.New(60*time.Minute, 70*time.Minute)

	for i := 1; i < 10; i++ {
		id := fmt.Sprintf("%d", i)
		c.Set(id, &pb.Product{
			Id:           id,
			Description:  "Produto Teste " + id,
			PriceInCents: 10090,
			Title:        "Produto " + id,
		}, cache.DefaultExpiration)

	}
}

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

	repository := NewRepository(dB)
	//defer repository.Close()
	return &ProductService{
		repo: repository,
	}
}

func getProduct(s *ProductService, id string) (*pb.Product, error) {
	return s.repo.GetByID(id)
}

// GetByID fetch product by ID
func (s *ProductService) GetByID(ctx context.Context, r *pb.RequestProduct) (*pb.Product, error) {
	id := r.GetId()
	return getProduct(s, id)
}

// List fetch products
func (s *ProductService) List(r *empty.Empty, stream pb.ProductService_ListServer) error {

	products, err := s.repo.List()
	if err != nil {
		return errors.New("Could not find products")
	}

	for _, product := range products {
		err := stream.Send(product)
		if err != nil {
			return err
		}

		log.Printf("sent product with id: %v", product.GetId())
	}

	return nil
}
