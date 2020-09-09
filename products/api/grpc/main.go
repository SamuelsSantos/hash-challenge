package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/SamuelsSantos/product-discount-service/products/config"
	"github.com/SamuelsSantos/product-discount-service/products/domain"
	pb "github.com/SamuelsSantos/product-discount-service/products/domain/pb"
	"google.golang.org/grpc"
)

func main() {

	cfg := config.NewConfig()
	service := domain.NewService(cfg)

	port := flag.Int("port", 50001, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	grpc := grpc.NewServer()

	pb.RegisterProductServiceServer(grpc, service)
	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	err = grpc.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
