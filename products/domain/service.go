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
func (s *ProductService) List(r *pb.Empty, stream pb.ProductService_ListServer) error {

	rows, err := s.repo.List()
	if err != nil {
		return errors.New("Could not find products")
	}

	for rows.Next() {
		var id string
		var title string
		var description string
		var priceInCents int64

		err = rows.Scan(&id, &title, &description, &priceInCents)
		if err != nil {
			return err
		}

		pbProduct := pb.Product{
			Id:           id,
			Title:        title,
			Description:  description,
			PriceInCents: priceInCents,
		}

		res := &pb.Response{Result: &pbProduct}
		err := stream.Send(res)
		if err != nil {
			return err
		}

		log.Printf("sent product with id: %v", id)
	}

	return nil
}
