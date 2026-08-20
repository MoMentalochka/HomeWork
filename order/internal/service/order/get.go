package order

import (
	"context"

	"github.com/MoMentalochka/HomeWork/order/internal/model"
)

func (s *orderService) GetOrderById(_ context.Context, uuid string) (*model.OrderDto, error) {
	order, err := s.store.Get(uuid)
	if err != nil {
		return &model.OrderDto{}, err
	}
	if order == nil {
		return nil, model.ErrPartNotFound
	}

	return order, nil
}
