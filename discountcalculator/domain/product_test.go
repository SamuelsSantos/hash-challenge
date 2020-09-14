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
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/patrickmn/go-cache"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type productService struct {
	cache *cache.Cache
}

func newProductService() *productService {

	c := cache.New(cache.NoExpiration, 15*time.Second)

	c.Set("1", &pb.Product{
		Id:           "1",
		Description:  "Produto Teste 01",
		PriceInCents: 100,
		Title:        "Produto 01"},
		cache.NoExpiration,
	)

	c.Set("2", &pb.Product{
		Id:           "2",
		Description:  "Produto Teste 02",
		PriceInCents: 200,
		Title:        "Produto 02"},
		cache.NoExpiration,
	)

	return &productService{
		cache: c,
	}
}

func (s *productService) GetByID(ctx context.Context, r *pb.RequestProduct) (*pb.Product, error) {
	id := r.GetId()

	productCache, found := s.cache.Get(id)
	if found {
		return productCache.(*pb.Product), nil
	}
	return nil, errors.New("Not found")
}

func (s *productService) List(r *empty.Empty, stream pb.ProductService_ListServer) error {
	var products []*pb.Product

	items := s.cache.Items()

	for id := range items {
		productCache, found := s.cache.Get(id)
		if found {
			product := productCache.(*pb.Product)
			products = append(products, product)
		}

	}

	for _, product := range products {
		err := stream.Send(product)
		if err != nil {
			return err
		}
	}

	return nil
}

func serverProductTest(t *testing.T, service *productService, port string) string {
	server := grpc.NewServer()
	pb.RegisterProductServiceServer(server, service)
	reflection.Register(server)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	go server.Serve(listener)

	return listener.Addr().String()
}

func TestProductService_GetProductByID(t *testing.T) {

	service := newProductService()
	address := serverProductTest(t, service, "50111")

	type fields struct {
		host string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "WhenGetByID_1",
			fields:  fields{host: address},
			args:    args{id: "1"},
			want:    "1",
			wantErr: false,
		},
		{
			name:    "WhenGetByID_3",
			fields:  fields{host: address},
			args:    args{id: "3"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ProductClient{
				host: tt.fields.host,
			}
			got, err := s.GetProductByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.GetProductByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.GetId(), tt.want) {
				t.Errorf("ProductService.GetProductByID() = %v, want %v", got.GetId(), tt.want)
			}
		})
	}
}
