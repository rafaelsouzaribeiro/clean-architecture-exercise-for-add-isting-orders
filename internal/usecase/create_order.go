package usecase

import (
	"github.com/rafaelsouzaribeiro/clean-architecture-exercise-for-add-isting-orders/internal/dto"
	"github.com/rafaelsouzaribeiro/clean-architecture-exercise-for-add-isting-orders/internal/entity"
)

func (c *OrderUseCase) Execute(input dto.OrderInputDTO) (dto.OrderOutputDTO, error) {
	order := entity.Order{
		ID:    input.ID,
		Price: input.Price,
		Tax:   input.Tax,
	}
	order.CalculateFinalPrice()
	if err := c.OrderRepository.Save(&order); err != nil {
		return dto.OrderOutputDTO{}, err
	}

	dto := dto.OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.Price + order.Tax,
	}

	c.OrderCreated.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.OrderCreated)

	return dto, nil
}
