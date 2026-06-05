package usecase

import "github.com/rafaelsouzaribeiro/clean-architecture-exercise-for-add-isting-orders/internal/dto"

func (l *OrderUseCase) GetAll() (*[]dto.OrderOutputDTO, error) {
	orders, err := l.OrderRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var output []dto.OrderOutputDTO
	for _, order := range *orders {
		output = append(output, dto.OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		})
	}
	return &output, nil
}
