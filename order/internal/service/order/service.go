package order

import (
	"context"
	"fmt"

	grpc "github.com/MoMentalochka/HomeWork/order/internal/client/grpc"
	"github.com/MoMentalochka/HomeWork/order/internal/converter"
	"github.com/MoMentalochka/HomeWork/order/internal/model"
	"github.com/MoMentalochka/HomeWork/order/internal/repository"
	ordersv1 "github.com/MoMentalochka/HomeWork/shared/pkg/openapi/order/v1"
	paymentv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/payment/v1"
	"github.com/google/uuid"
)

type orderService struct {
	store           repository.OrderRepository
	paymentClient   grpc.PaymentClient
	inventoryClient grpc.InventoryClient
}

func NewOrderService(rep repository.OrderRepository, paymentClient grpc.PaymentClient, inventoryClient grpc.InventoryClient) *orderService {
	return &orderService{
		rep,
		paymentClient,
		inventoryClient,
	}
}

func (s *orderService) CreateNewOrder(ctx context.Context, req model.CreateOrderRequest) (*ordersv1.CreateOrderResponse, error) {

	parts, err := s.inventoryClient.ListParts(ctx, model.PartsFilter{Uuids: req.PartUuids})

	if err != nil {
		return &ordersv1.CreateOrderResponse{}, err
	}

	total := 0.0
	for _, part := range parts {

		total += part.Price
	}

	order := &model.OrderDto{
		OrderUUID:  uuid.New().String(),
		PartUuids:  req.PartUuids,
		UserUUID:   req.UserUUID,
		TotalPrice: total,
		Status:     converter.ProtoStatusToModel(ordersv1.StatusPENDINGPAYMENT),
	}

	s.store.AddOrder(order)
	return &ordersv1.CreateOrderResponse{OrderUUID: order.OrderUUID, TotalPrice: total}, nil
}

func (s *orderService) GetOrderById(_ context.Context, uuid string) (*model.OrderDto, error) {

	order := s.store.GetOrder(uuid)
	if order == nil {
		return nil, model.ErrPartNotFound
	}

	return order, nil
}

func (s *orderService) OrderCancel(_ context.Context, uuid string) (ordersv1.OrderCancelRes, error) {
	order := s.store.GetOrder(uuid)
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
	return &ordersv1.OrderCancelResponse{TransactionUUID: order.TransactionUUID}, nil
}

func (s *orderService) OrderPay(ctx context.Context, req ordersv1.OptOrderPayRequest, uuid string) (ordersv1.OrderPayRes, error) {
	order := s.store.GetOrder(uuid)

	if order == nil {
		return &ordersv1.NotFound{
			Code:    404,
			Message: fmt.Sprintf("Order with uuid '%s' not found", uuid),
		}, nil
	}

	if order.Status == "CANCELLED" || order.Status == "PAID" {
		return &ordersv1.Conflict{
			Code:    409,
			Message: fmt.Sprintf("Order with uuid '%s' is already paid or cancelled", uuid),
		}, nil
	}

	method := paymentv1.PaymentMethod_value[string(req.Value.PaymentMethod)]
	payRequest := paymentv1.PayOrderRequest{
		UserUuid:      order.UserUUID,
		OrderUuid:     order.OrderUUID,
		PaymentMethod: paymentv1.PaymentMethod(method),
	}

	transactionUuid, err := s.paymentClient.PayOrder(ctx, &payRequest)

	if err != nil {
		return nil, err
	}
	order.Status = "PAID"
	order.PaymentMethod = string(req.Value.PaymentMethod)
	order.TransactionUUID = transactionUuid
	return &ordersv1.OrderPayResponse{TransactionUUID: order.TransactionUUID}, nil
}
