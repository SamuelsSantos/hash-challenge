package domain

import (
	"context"
	"log"
	"time"

	"github.com/SamuelsSantos/product-discount-service/discountcalculator/domain/pb"
	"google.golang.org/grpc"
)

// UserClient definition
type UserClient struct {
	host string
}

// UserService definition
type UserService interface {
	GetUserByID(id string) (*pb.User, error)
}

// NewUserClient create a new instance UserClient
func NewUserClient(host string) *UserClient {
	return &UserClient{host: host}
}

// GetUserByID fetch product by ID
func (s *UserClient) GetUserByID(id string) (*pb.User, error) {
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
