package main

import (
	"context"
	"io"
	"log"

	pb "github.com/SamuelsSantos/product-discount-service/products/domain/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	router := gin.Default()

	rg := router.Group("api/v1/product")
	{
		rg.GET("/:id", fetchProduct)
		rg.GET("/", fetchAll)
	}

	router.Run(":8080")
}

func fetchProduct(c *gin.Context) {
	sAddress := "localhost:8486"
	conn, e := grpc.Dial(sAddress, grpc.WithInsecure())
	if e != nil {
		log.Fatalf("Failed to connect to server %v", e)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)
	value, e := client.GetByID(context.Background(), &pb.Request{
		Id: c.Param("id"),
	})
	if e != nil {
		log.Fatalf("Failed to get data: %v", e)
	}

	c.JSON(200, &value)
}

func fetchAll(c *gin.Context) {
	sAddress := "localhost:8486"
	conn, e := grpc.Dial(sAddress, grpc.WithInsecure())
	if e != nil {
		log.Fatalf("Failed to connect to server %v", e)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)
	stream, err := client.List(context.Background(), &pb.Empty{})

	if err != nil {
		log.Fatal("cannot find products: ", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal("cannot receive response: ", err)
		}

		product := res.GetResult()
		c.JSON(200, &product)
	}
}
