package domain

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/SamuelsSantos/product-discount-service/discountcalculator/config"
	"github.com/SamuelsSantos/product-discount-service/discountcalculator/domain/pb"
	"github.com/golang/protobuf/ptypes"
)

// CalculatorServer that provides a Discount Calculate Service
type CalculatorServer struct {
	userService    *UserService
	productService *ProductService
	cfg            *config.Config
}

// NewCalculatorServer returns a new instance of Server
func NewCalculatorServer(userService *UserService, productService *ProductService, cfg *config.Config) *CalculatorServer {
	return &CalculatorServer{
		userService:    userService,
		productService: productService,
		cfg:            cfg,
	}
}

// Process calculate discount
func (s *CalculatorServer) Process(ctx context.Context, req *pb.DiscountRequest) (*pb.DiscountResponse, error) {

	var dateOfBirth time.Time
	if strings.TrimSpace(req.GetUserId()) != "" {
		user, err := s.userService.GetUserByID(req.GetUserId())
		if err != nil {
			log.Println(err)
		} else {
			if dateOfBirth, err = ptypes.Timestamp(user.GetDateOfBirth()); err != nil {
				log.Println(err)
			}
			log.Printf("Calculate discount to user: %v", user)
		}
	}

	product, err := s.productService.GetProductByID(req.GetProductId())
	if err != nil {
		return nil, err
	}

	log.Printf("Calculate discount to product: %v", product)
	pct := GetDiscountPercentual(s.cfg.BlackFridayDate, dateOfBirth)

	product.Discount = &pb.Discount{
		Pct:          pct,
		ValueInCents: Calculate(pct, product.GetPriceInCents()),
	}

	return &pb.DiscountResponse{
		Result: product,
	}, nil
}
