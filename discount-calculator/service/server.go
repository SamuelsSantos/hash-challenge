package service

import (
	"context"
	"fmt"

	"github.com/SamuelsSantos/product-discount-service/discount-calculator/pb"
)

// Server that provides a Discount Calculate Service
type Server struct{}

// NewServer returns a new instance of Server
func NewServer() *Server {
	return &Server{}
}

// Process calculate discount
func (server *Server) Process(ctx context.Context, req *pb.DiscountRequest) (*pb.DiscountResponse, error) {
	res := &pb.DiscountResponse{
		Msg: fmt.Sprintf("{\"Data\": %v - %v}", req.GetProductId(), req.GetUserId()),
	}

	fmt.Println(res.Msg)
	return res, nil
}
