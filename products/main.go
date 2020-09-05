package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/SamuelsSantos/product-discount-service/products/pb"
	"github.com/SamuelsSantos/product-discount-service/products/service"
	"google.golang.org/grpc"
)

func main() {

	service.SetCache()

	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	server := grpc.NewServer()

	pb.RegisterProductServiceServer(server, service.NewServer())
	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	err = server.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
