package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/SamuelsSantos/product-discount-service/products/config"
	"github.com/SamuelsSantos/product-discount-service/products/domain"
	pb "github.com/SamuelsSantos/product-discount-service/products/domain/pb"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	cfg := config.NewConfig()
	repository := domain.NewSQLRepository(cfg)
	service := domain.NewProductService(repository)

	port := flag.Int("port", 50001, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	grpc := grpc.NewServer()

	pb.RegisterProductServiceServer(grpc, service)
	address := fmt.Sprintf("0.0.0.0:%d", *port)
	reflection.Register(grpc)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	err = grpc.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
