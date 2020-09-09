package domain

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/SamuelsSantos/product-discount-service/users/config"
	"github.com/SamuelsSantos/product-discount-service/users/domain/pb"
	"github.com/golang/protobuf/ptypes"
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

func getUser(s *UserService, id string) (*pb.Response, error) {
	rows, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("could not find User with id: %v", id)
	}

	var user *pb.User

	for rows.Next() {
		var id string
		var firstName string
		var lastName string
		var dateOfBirth time.Time

		err = rows.Scan(&id, &firstName, &lastName, &dateOfBirth)
		if err != nil {
			return nil, err
		}

		dtProto, err := ptypes.TimestampProto(dateOfBirth)
		if err != nil {
			return nil, err
		}

		pbUser := pb.User{
			Id:          id,
			FirstName:   firstName,
			LastName:    lastName,
			DateOfBirth: dtProto,
		}

		user = &pbUser
	}

	return &pb.Response{
		Result: user,
	}, nil
}

// GetByID fetch user by ID
func (s *UserService) GetByID(ctx context.Context, r *pb.Request) (*pb.Response, error) {
	id := r.GetId()

	return getUser(s, id)
}
