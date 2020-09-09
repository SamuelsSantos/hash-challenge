package main

import (
	"context"
	"log"

	products "github.com/SamuelsSantos/product-discount-service/products/domain/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	router := gin.Default()

	rg := router.Group("api/v1/product")
	{
		rg.GET("/:id", fetchProduct)
	}

	router.Run()
}

func fetchProduct(c *gin.Context) {
	sAddress := "localhost:8486"
	conn, e := grpc.Dial(sAddress, grpc.WithInsecure())
	if e != nil {
		log.Fatalf("Failed to connect to server %v", e)
	}
	defer conn.Close()

	client := products.NewProductServiceClient(conn)
	player, e := client.GetByID(context.Background(), &products.Request{
		Id: c.Param("id"),
	})
	if e != nil {
		log.Fatalf("Failed to get player data: %v", e)
	}

	c.JSON(200, &player)
}
