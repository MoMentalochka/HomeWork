package service

import (
	"context"

	"github.com/MoMentalochka/HomeWork/order/internal/model"
	ordersv1 "github.com/MoMentalochka/HomeWork/shared/pkg/openapi/order/v1"
)

type OrderService interface {
	CreateNewOrder(ctx context.Context, req model.CreateOrderRequest) (*ordersv1.CreateOrderResponse, error)
	GetOrderById(_ context.Context, uuid string) (*model.OrderDto, error)
	OrderCancel(_ context.Context, uuid string) (ordersv1.OrderCancelRes, error)
	OrderPay(ctx context.Context, method ordersv1.OptOrderPayRequest, uuid string) (ordersv1.OrderPayRes, error)
}

type OrderProducerService interface {
	ProduceOrderPaid(ctx context.Context, event model.OrderPaidEvent) error
}
