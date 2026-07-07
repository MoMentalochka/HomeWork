package service

import (
	"context"

	ordersv1 "github.com/MoMentalochka/HomeWork/shared/pkg/openapi/order/v1"
)

type OrderService interface {
	CreateNewOrder(ctx context.Context, req *ordersv1.CreateOrderRequest) (*ordersv1.CreateOrderResponse, error)
	GetOrderById(_ context.Context, params ordersv1.GetOrderByIdParams) (ordersv1.GetOrderByIdRes, error)
	OrderCancel(_ context.Context, params ordersv1.OrderCancelParams) (ordersv1.OrderCancelRes, error)
	OrderPay(ctx context.Context, req ordersv1.OptOrderPayRequest, params ordersv1.OrderPayParams) (ordersv1.OrderPayRes, error)
}
