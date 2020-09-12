package domain

import (
	"context"
	"database/sql"
	"time"

	"github.com/SamuelsSantos/product-discount-service/users/config"
	"github.com/SamuelsSantos/product-discount-service/users/domain/pb"
	"github.com/patrickmn/go-cache"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var c *cache.Cache

// SetCache define cache
func SetCache() {
	c = cache.New(60*time.Minute, 70*time.Minute)
	c.Set("1", &pb.User{
		Id:          "1",
		FirstName:   "James",
		LastName:    "LeBron",
		DateOfBirth: timestamppb.New(time.Now()),
	}, cache.DefaultExpiration)
}

//UserService  interface
type UserService struct {
	repo *SQLRepo
}

//NewService create new service
func NewService(cfg *config.Config) *UserService {

	dB, err := sql.Open(cfg.Db.Driver, cfg.Db.ToURL())
	if err != nil {
		panic(err)
	}

	return &UserService{
		repo: NewRepository(dB),
	}
}

// GetByID fetch user by ID
func (s *UserService) GetByID(ctx context.Context, r *pb.RequestUser) (*pb.User, error) {
	id := r.GetId()
	return s.repo.GetByID(id)
}
