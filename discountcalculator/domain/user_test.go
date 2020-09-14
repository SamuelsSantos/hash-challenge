package domain

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/SamuelsSantos/product-discount-service/discountcalculator/domain/pb"
	"github.com/patrickmn/go-cache"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type userClient struct {
	cache *cache.Cache
}

func newUserService() *userClient {

	c := cache.New(cache.NoExpiration, 15*time.Second)

	c.Set("1", &pb.User{
		Id:          "1",
		FirstName:   "James",
		LastName:    "LeBron",
		DateOfBirth: timestamppb.New(time.Now().AddDate(0, 0, -1))},
		cache.NoExpiration,
	)

	c.Set("2", &pb.User{
		Id:          "2",
		FirstName:   "Oscar",
		LastName:    "Schmidt",
		DateOfBirth: timestamppb.New(time.Now())},
		cache.NoExpiration,
	)

	return &userClient{
		cache: c,
	}
}

func (s *userClient) GetByID(ctx context.Context, r *pb.RequestUser) (*pb.User, error) {
	id := r.GetId()

	userCache, found := s.cache.Get(id)
	if found {
		return userCache.(*pb.User), nil
	}
	return nil, errors.New("Not found")
}

func serverUserTest(t *testing.T, service *userClient, port string) string {
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, service)
	reflection.Register(server)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	go server.Serve(listener)

	return listener.Addr().String()
}

func Test_getUserByID(t *testing.T) {

	service := newUserService()
	address := serverUserTest(t, service, "50110")

	type args struct {
		host string
		id   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "When GetById 1",
			args: args{
				host: address,
				id:   "1",
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "When GetById 3",
			args: args{
				host: address,
				id:   "3",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getUserByID(tt.args.host, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("getUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.GetId(), tt.want) {
				t.Errorf("getUserByID() = %v, want %v", got.GetId(), tt.want)
			}
		})
	}
}
