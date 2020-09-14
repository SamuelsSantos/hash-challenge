package domain

import (
	"context"
	"io"
	"log"
	"net"
	"reflect"
	"testing"

	"github.com/SamuelsSantos/product-discount-service/products/domain/pb"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/patrickmn/go-cache"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func TestProductService_GetByID(t *testing.T) {
	type fields struct {
		repo  ProductRepository
		cache *cache.Cache
	}
	type args struct {
		ctx context.Context
		r   *pb.RequestProduct
	}

	repository := NewInMemoryRepository()
	service := NewProductService(repository)

	f := fields{
		repo:  service.repo,
		cache: service.cache,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "WhenGetByID_1",
			fields: f,
			args: args{
				ctx: context.Background(),
				r:   &pb.RequestProduct{Id: "1"},
			},
			want:    "1",
			wantErr: false,
		},
		{
			name:   "WhenGetByID_3",
			fields: f,
			args: args{
				ctx: context.Background(),
				r:   &pb.RequestProduct{Id: "3"},
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ProductService{
				repo:  tt.fields.repo,
				cache: tt.fields.cache,
			}
			got, err := s.GetByID(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.GetId(), tt.want) {
				t.Errorf("ProductService.GetByID() = %v, want %v", got.GetId(), tt.want)
			}
		})
	}
}

func serverTest(t *testing.T, service *ProductService) string {
	server := grpc.NewServer()
	pb.RegisterProductServiceServer(server, service)
	reflection.Register(server)
	listener, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	go server.Serve(listener)

	return listener.Addr().String()
}

func TestProductService_List(t *testing.T) {
	type fields struct {
		repo  ProductRepository
		cache *cache.Cache
	}
	repository := NewInMemoryRepository()
	service := NewProductService(repository)

	address := serverTest(t, service)

	conn, e := grpc.Dial(address, grpc.WithInsecure())
	if e != nil {
		log.Fatalf("Failed to connect to server %v", e)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)
	stream, err := client.List(context.Background(), &empty.Empty{})
	if err != nil {
		t.Errorf("ProductService.GetByID() error = %v", err)
		return
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			t.Errorf("ProductService.GetByID() error = %v", err)
		}
		if res == nil {
			t.Error("Want two records")
		}
	}
}
