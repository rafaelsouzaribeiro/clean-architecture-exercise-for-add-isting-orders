package usecase

import (
	"github.com/rafaelsouzaribeiro/clean-architecture-exercise-for-add-isting-orders/internal/entity"
	"github.com/rafaelsouzaribeiro/clean-architecture-exercise-for-add-isting-orders/pkg/events"
)

type OrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderCreated    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *OrderUseCase {
	return &OrderUseCase{
		OrderRepository: OrderRepository,
		OrderCreated:    OrderCreated,
		EventDispatcher: EventDispatcher,
	}
}
