package domain

import (
	"context"
	"errors"

	"github.com/SamuelsSantos/product-discount-service/discountcalculator/domain/pb"
	"google.golang.org/grpc"
)

// UserService definition
type UserService struct {
	host string
}

// NewUserService create a new instance UserService
func NewUserService(host string) *UserService {
	return &UserService{host: host}
}

// GetUserByID fetch product by ID
func (s *UserService) GetUserByID(id string) (*pb.User, error) {
	return getUserByID(s.host, id)
}

func getUserByID(host string, id string) (*pb.User, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, errors.New("Failed to fetch data from server")
	}
	defer conn.Close()

	user, err := pb.NewUserServiceClient(conn).GetByID(context.Background(), &pb.RequestUser{Id: id})
	if err != nil {
		return nil, errors.New("Failed to connect to server")
	}

	return user, nil
}
