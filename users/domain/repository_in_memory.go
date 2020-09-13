package domain

import (
	"errors"
	"time"

	"github.com/SamuelsSantos/product-discount-service/users/domain/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//InMemoryRepo repository
type InMemoryRepo struct {
	data map[string]*pb.User
}

//NewInMemoryRepository create new repository
func NewInMemoryRepository() *InMemoryRepo {
	data := make(map[string]*pb.User, 2)

	data["1"] = &pb.User{
		Id:          "1",
		FirstName:   "James",
		LastName:    "LeBron",
		DateOfBirth: timestamppb.New(time.Now()),
	}

	data["2"] = &pb.User{
		Id:          "2",
		FirstName:   "Oscar",
		LastName:    "Schmidt",
		DateOfBirth: timestamppb.New(time.Now()),
	}

	return &InMemoryRepo{data}
}

//Close database connection
func (r *InMemoryRepo) Close() error {
	r.data = make(map[string]*pb.User, 0)
	return nil
}

// GetByID fetch user by ID
func (r *InMemoryRepo) GetByID(id string) (*pb.User, error) {

	user := r.data[id]
	if user == nil {
		return nil, errors.New("Not found")
	}

	return user, nil
}
