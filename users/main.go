package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/SamuelsSantos/product-discount-service/users/pb"
	"github.com/SamuelsSantos/product-discount-service/users/service"
	"google.golang.org/grpc"
)

func main() {

	service.SetCache()

	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	grpc := grpc.NewServer()

	pb.RegisterUserServiceServer(grpc, service.NewServer{})
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
