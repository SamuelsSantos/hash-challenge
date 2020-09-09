package main

import (
	"context"
	"log"

	users "github.com/SamuelsSantos/product-discount-service/users/domain/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	router := gin.Default()

	rg := router.Group("api/v1/user")
	{
		rg.GET("/:id", fetchUser)
	}

	router.Run(":8081")
}

func fetchUser(c *gin.Context) {
	sAddress := "localhost:8485"
	conn, e := grpc.Dial(sAddress, grpc.WithInsecure())
	if e != nil {
		log.Fatalf("Failed to connect to server %v", e)
	}
	defer conn.Close()

	client := users.NewUserServiceClient(conn)
	player, e := client.GetByID(context.Background(), &users.Request{
		Id: c.Param("id"),
	})
	if e != nil {
		log.Fatalf("Failed to get player data: %v", e)
	}

	c.JSON(200, &player)
}
