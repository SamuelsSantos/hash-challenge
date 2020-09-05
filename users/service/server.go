package service

import (
	"context"
	"fmt"
	"time"

	pb "github.com/SamuelsSantos/product-discount-service/users/pb"
	"github.com/patrickmn/go-cache"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	c.Set("1", &pb.User{
		Id:          "1",
		FirstName:   "James",
		LastName:    "LeBron",
		DataOfBirth: timestamppb.New(time.Now()),
	}, cache.DefaultExpiration)

}

// GetByID ... fetch user by id
func (*Server) GetByID(ctx context.Context, r *pb.Request) (*pb.Response, error) {
	id := r.GetId()

	user, found := c.Get(id)
	if found {
		value := user.(*pb.User)
		return &pb.Response{
			Result: value,
		}, nil
	}

	return nil, fmt.Errorf("could not find User with id: %v", id)
}
