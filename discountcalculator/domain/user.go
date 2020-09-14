package domain

import (
	"context"
	"log"
	"time"

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
		log.Fatalln(err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	user, err := client.GetByID(ctx, &pb.RequestUser{Id: id})
	if err != nil {
		return nil, err
	}

	return user, nil
}
