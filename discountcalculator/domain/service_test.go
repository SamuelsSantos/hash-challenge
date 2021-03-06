package domain

import (
	"context"
	"fmt"
	"log"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/SamuelsSantos/product-discount-service/discountcalculator/config"
	"github.com/SamuelsSantos/product-discount-service/discountcalculator/domain/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func serverTest(t *testing.T, service *CalculatorServer, port string) string {
	server := grpc.NewServer()
	pb.RegisterDiscountCalculatorServiceServer(server, service)
	reflection.Register(server)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	go server.Serve(listener)

	return listener.Addr().String()
}

func TestCalculatorServerWhenBlackFriday_Process(t *testing.T) {

	cfg := getConfig(time.Now(), "50010", "50011")
	server := configServer(t, cfg)
	serverTest(t, server, "50009")

	type args struct {
		ctx context.Context
		req *pb.DiscountRequest
	}

	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "When Is BlackFriday and Is BirthDay then 10% discount (MAX)",
			args: args{
				ctx: context.Background(),
				req: &pb.DiscountRequest{
					ProductId: "1",
					UserId:    "2",
				},
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "When Is BlackFriday and Isn't BirthDay then 10% discount (MAX)",
			args: args{
				ctx: context.Background(),
				req: &pb.DiscountRequest{
					ProductId: "1",
					UserId:    "1",
				},
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "When Is BlackFriday and user is undefined  then 10% discount (MAX)",
			args: args{
				ctx: context.Background(),
				req: &pb.DiscountRequest{
					ProductId: "1",
				},
			},
			want:    10,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := server
			got, err := s.Process(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculatorServer.Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			discount := got.GetResult().GetDiscount().GetValueInCents()
			if !reflect.DeepEqual(discount, tt.want) {
				t.Errorf("CalculatorServer.Process() = %v, want %v", discount, tt.want)
			}
		})
	}
}

func getConfig(blackFriday time.Time, userPort, productPort string) *config.Config {

	cfg := config.NewConfig()
	cfg.BlackFridayDate = blackFriday
	cfg.UserConfig.Host = "Localhost"
	cfg.UserConfig.Port = userPort
	cfg.ProductConfig.Host = "Localhost"
	cfg.ProductConfig.Port = productPort
	return cfg
}

func configServer(t *testing.T, cfg *config.Config) *CalculatorServer {
	serviceUser := newUserService()
	addressUser := serverUserTest(t, serviceUser, cfg.UserConfig.Port)
	serviceProduct := newProductService()
	addressProduct := serverProductTest(t, serviceProduct, cfg.ProductConfig.Port)

	clienteUser := NewUserClient(addressUser)
	clientProduct := NewProductClient(addressProduct)

	return NewCalculatorServer(clienteUser, clientProduct, cfg)
}

func TestCalculatorServerWhenIsJustBirthDay_Process(t *testing.T) {

	cfg := getConfig(time.Now().AddDate(0, 0, -1), "50012", "50013")
	server := configServer(t, cfg)
	serverTest(t, server, "50014")

	type args struct {
		ctx context.Context
		req *pb.DiscountRequest
	}

	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "When Is BirthDay then 5% discount",
			args: args{
				ctx: context.Background(),
				req: &pb.DiscountRequest{
					ProductId: "2",
					UserId:    "2",
				},
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "When Is BirthDay and user is undefined then no discount",
			args: args{
				ctx: context.Background(),
				req: &pb.DiscountRequest{
					ProductId: "1",
				},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := server
			got, err := s.Process(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculatorServer.Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			discount := got.GetResult().GetDiscount().GetValueInCents()
			if !reflect.DeepEqual(discount, tt.want) {
				t.Errorf("CalculatorServer.Process() = %v, want %v", discount, tt.want)
			}
		})
	}
}
