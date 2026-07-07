package order

import (
	"context"
	"fmt"

	"github.com/MoMentalochka/HomeWork/order/internal/repository"
	ordersv1 "github.com/MoMentalochka/HomeWork/shared/pkg/openapi/order/v1"
	inventoryv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
	paymentv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/payment/v1"
	"github.com/google/uuid"
)

type OrderService struct {
	store repository.OrderRepository
}

func NewOrderService(rep repository.OrderRepository) *OrderService {
	return &OrderService{
		store: rep,
	}
}

func (s *OrderService) CreateNewOrder(ctx context.Context, req *ordersv1.CreateOrderRequest) (*ordersv1.CreateOrderResponse, error) {

	for _, partUuid := range req.PartUuids {
		_, err := s.inventoryClient.GetPart(ctx, &inventoryv1.GetPartRequest{Uuid: partUuid})
		if err != nil {
			return nil, err
		}
	}

	order := &ordersv1.OrderDto{
		OrderUUID: uuid.New().String(),
		PartUuids: req.PartUuids,
		UserUUID:  req.UserUUID,
		Status:    ordersv1.StatusPENDINGPAYMENT,
	}

	s.store.AddOrder(order)
	return &ordersv1.CreateOrderResponse{OrderUUID: order.OrderUUID, TotalPrice: 0.0}, nil
}

func (s *OrderService) GetOrderById(_ context.Context, params ordersv1.GetOrderByIdParams) (ordersv1.GetOrderByIdRes, error) {

	order := s.store.GetOrder(params.OrderUUID)

	if order == nil {
		return &ordersv1.NotFound{
			Code:    404,
			Message: fmt.Sprintf("Order with uuid '%s' not found", params.OrderUUID),
		}, nil
	}

	return order, nil
}

func (s *OrderService) OrderCancel(_ context.Context, params ordersv1.OrderCancelParams) (ordersv1.OrderCancelRes, error) {
	order := s.store.GetOrder(params.OrderUUID)
	if order == nil {
		return &ordersv1.NotFound{
			Code:    404,
			Message: fmt.Sprintf("Order with uuid '%s' not found", params.OrderUUID),
		}, nil
	}
	if order.Status == "PAID" || order.Status == "CANCELLED" {
		return &ordersv1.Conflict{
			Code:    409,
			Message: fmt.Sprintf("Order with uuid '%s' is already paid or cancelled", params.OrderUUID),
		}, nil
	}
	order.Status = "CANCELLED"
	return &ordersv1.OrderCancelResponse{TransactionUUID: order.TransactionUUID}, nil
}

func (s *OrderService) OrderPay(ctx context.Context, req ordersv1.OptOrderPayRequest, params ordersv1.OrderPayParams) (ordersv1.OrderPayRes, error) {
	order := s.store.GetOrder(params.OrderUUID)

	if order == nil {
		return &ordersv1.NotFound{
			Code:    404,
			Message: fmt.Sprintf("Order with uuid '%s' not found", params.OrderUUID),
		}, nil
	}

	if order.Status == "CANCELLED" || order.Status == "PAID" {
		return &ordersv1.Conflict{
			Code:    404,
			Message: fmt.Sprintf("Order with uuid '%s' is already paid or cancelled", params.OrderUUID),
		}, nil
	}

	method := paymentv1.PaymentMethod_value[string(req.Value.PaymentMethod)]
	payRequest := paymentv1.PayOrderRequest{
		UserUuid:      order.UserUUID,
		OrderUuid:     order.OrderUUID,
		PaymentMethod: paymentv1.PaymentMethod(method),
	}
	res, err := s.paymentClient.PayOrder(ctx, &payRequest)
	if err != nil {
		return nil, err
	}
	order.Status = "PAID"
	order.PaymentMethod = req.Value.PaymentMethod
	order.TransactionUUID = res.TransactionUuid
	return &ordersv1.OrderPayResponse{TransactionUUID: order.TransactionUUID}, nil
}
