package domain

import (
	"context"
	"errors"
	"time"

	"github.com/SamuelsSantos/product-discount-service/users/domain/pb"
	"github.com/patrickmn/go-cache"
)

//UserService  interface
type UserService struct {
	repo  UserRepository
	cache *cache.Cache
}

//NewUserService create new service
func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo:  repo,
		cache: cache.New(60*time.Second, 70*time.Second),
	}
}

// GetByID fetch user by ID
func (s *UserService) GetByID(ctx context.Context, r *pb.RequestUser) (*pb.User, error) {
	id := r.GetId()

	userCache, found := s.cache.Get(id)
	if found {
		return userCache.(*pb.User), nil
	}

	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("Not found")
	}

	s.cache.Set(user.GetId(), user, cache.DefaultExpiration)
	return user, nil
}
