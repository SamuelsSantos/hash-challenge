package service

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/SamuelsSantos/product-discount-service/products/domain/pb"
	"github.com/patrickmn/go-cache"
)

// Server that provides a User Service
type Server struct{}

// NewServer returns a new instance of Server
func NewServer() *Server {
	return &Server{}
}

var c *cache.Cache

// SetCache define cache
func SetCache() {

	c = cache.New(60*time.Minute, 70*time.Minute)

	for i := 1; i < 10; i++ {
		id := fmt.Sprintf("%d", i)
		c.Set(id, &pb.Product{
			Id:           id,
			Descricao:    "Produto Teste " + id,
			PriceInCents: 10090,
			Title:        "Produto " + id,
		}, cache.DefaultExpiration)

	}
}

// GetByID fetch product by ID
func (*Server) GetByID(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	id := req.GetId()

	product, found := c.Get(id)
	if found {
		value := product.(*pb.Product)
		return &pb.Response{
			Result: value,
		}, nil
	}

	return nil, fmt.Errorf("could not find product with id: %v", id)
}

// GetAll ... fetch all products by server stream
func (*Server) GetAll(req *pb.Empty, stream pb.ProductService_GetAllServer) error {

	for id := range c.Items() {
		product, found := c.Get(id)

		log.Printf("sent product with id: %v", product)
		if found {
			value := product.(*pb.Product)
			res := &pb.Response{Result: value}
			err := stream.Send(res)
			if err != nil {
				return err
			}

			log.Printf("sent product with id: %v", id)
		}
	}

	return nil
}
