package order

import (
	"context"
	"fmt"

	ordersv1 "github.com/MoMentalochka/HomeWork/shared/pkg/openapi/order/v1"
)

func (s *orderService) OrderCancel(_ context.Context, uuid string) (ordersv1.OrderCancelRes, error) {
	order, err := s.store.Get(uuid)
	if err != nil {
		return &ordersv1.OrderCancelResponse{}, err
	}
	if order == nil {
		return &ordersv1.NotFound{
			Code:    404,
			Message: fmt.Sprintf("Order with uuid '%s' not found", uuid),
		}, nil
	}
	if order.Status == "PAID" || order.Status == "CANCELLED" {
		return &ordersv1.Conflict{
			Code:    409,
			Message: fmt.Sprintf("Order with uuid '%s' is already paid or cancelled", uuid),
		}, nil
	}

	order.Status = "CANCELLED"
	err = s.store.Update(uuid, order)
	if err != nil {
		return &ordersv1.OrderCancelResponse{}, err
	}

	return &ordersv1.OrderCancelResponse{TransactionUUID: order.TransactionUUID}, nil
}
