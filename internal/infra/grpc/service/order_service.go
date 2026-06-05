package service

import (
	"context"

	"github.com/rafaelsouzaribeiro/clean-architecture-exercise-for-add-isting-orders/internal/dto"
	"github.com/rafaelsouzaribeiro/clean-architecture-exercise-for-add-isting-orders/internal/infra/grpc/pb"
	"github.com/rafaelsouzaribeiro/clean-architecture-exercise-for-add-isting-orders/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.OrderUseCase
}

func NewOrderService(createOrderUseCase usecase.OrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := dto.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	output, err := s.CreateOrderUseCase.GetAll()
	if err != nil {
		return nil, err
	}
	var orders []*pb.Order
	for _, order := range *output {
		orders = append(orders, &pb.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}
	return &pb.ListOrdersResponse{
		Orders: orders,
	}, nil
}
