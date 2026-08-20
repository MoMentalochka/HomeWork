package order

import (
	"context"
	"fmt"

	"github.com/MoMentalochka/HomeWork/order/internal/converter"
	"github.com/MoMentalochka/HomeWork/order/internal/model"
	ordersv1 "github.com/MoMentalochka/HomeWork/shared/pkg/openapi/order/v1"
	"github.com/google/uuid"
)

func (s *orderService) CreateNewOrder(ctx context.Context, req model.CreateOrderRequest) (*ordersv1.CreateOrderResponse, error) {
	total := 0.0
	for _, id := range req.PartUuids {
		part, err := s.inventoryClient.GetPart(ctx, id)
		if err != nil {
			return &ordersv1.CreateOrderResponse{}, fmt.Errorf("failed to get part by uuid %s: %w", id, err)
		}
		total += part.Price
	}

	order := &model.OrderDto{
		OrderUUID:  uuid.New().String(),
		PartUuids:  req.PartUuids,
		UserUUID:   req.UserUUID,
		TotalPrice: total,
		Status:     converter.ProtoStatusToModel(ordersv1.StatusPENDINGPAYMENT),
	}

	err := s.store.Create(order)
	if err != nil {
		return &ordersv1.CreateOrderResponse{}, err
	}
	return &ordersv1.CreateOrderResponse{OrderUUID: order.OrderUUID, TotalPrice: total}, nil
}
