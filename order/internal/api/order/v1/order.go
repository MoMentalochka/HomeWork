package v1

import (
	"context"
	"fmt"
	"log"

	"github.com/MoMentalochka/HomeWork/order/internal/service"
	ordersv1 "github.com/MoMentalochka/HomeWork/shared/pkg/openapi/order/v1"
	inventoryv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
	paymentv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/payment/v1"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	paymentPort   = "50052"
	inventoryPort = "50051"
)

type orderApi struct {
	orderService    service.OrderService
	paymentClient   paymentv1.PaymentServiceClient
	inventoryClient inventoryv1.InventoryServiceClient
}

func NewOrderApi(service service.OrderService) (*orderApi, error) {

	paymentConn, err := grpc.NewClient(
		fmt.Sprintf("localhost:%s", paymentPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("failed to connect: %v\n", err)
		return &orderApi{}, fmt.Errorf("failed to connect: %v\n", err)
	}

	inventoryConn, err := grpc.NewClient(
		fmt.Sprintf("localhost:%s", inventoryPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("failed to connect: %v\n", err)
		return &orderApi{}, fmt.Errorf("failed to connect: %v\n", err)
	}

	defer func() {
		if err := inventoryConn.Close(); err != nil {
			log.Printf("failed to close inventoryConn: %v", err)
		}
		if err := paymentConn.Close(); err != nil {
			log.Printf("failed to close paymentConn: %v", err)
		}
	}()

	return &orderApi{
		orderService:    service,
		paymentClient:   paymentv1.NewPaymentServiceClient(paymentConn),
		inventoryClient: inventoryv1.NewInventoryServiceClient(inventoryConn),
	}, nil
}

func (a *orderApi) CreateNewOrder(ctx context.Context, req *ordersv1.CreateOrderRequest) (*ordersv1.CreateOrderResponse, error) {

	for _, partUuid := range req.PartUuids {
		_, err := a.inventoryClient.GetPart(ctx, &inventoryv1.GetPartRequest{Uuid: partUuid})
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

	a.store.AddOrder(order)
	return &ordersv1.CreateOrderResponse{OrderUUID: order.OrderUUID, TotalPrice: 0.0}, nil
}

func (a *orderApi) GetOrderById(_ context.Context, params ordersv1.GetOrderByIdParams) (ordersv1.GetOrderByIdRes, error) {

	order := a.store.GetOrder(params.OrderUUID)

	if order == nil {
		return &ordersv1.NotFound{
			Code:    404,
			Message: fmt.Sprintf("Order with uuid '%s' not found", params.OrderUUID),
		}, nil
	}

	return order, nil
}

func (a *orderApi) OrderCancel(_ context.Context, params ordersv1.OrderCancelParams) (ordersv1.OrderCancelRes, error) {
	order := a.store.GetOrder(params.OrderUUID)
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

func (a *orderApi) OrderPay(ctx context.Context, req ordersv1.OptOrderPayRequest, params ordersv1.OrderPayParams) (ordersv1.OrderPayRes, error) {
	order := a.store.GetOrder(params.OrderUUID)

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
	res, err := a.paymentClient.PayOrder(ctx, &payRequest)
	if err != nil {
		return nil, err
	}
	order.Status = "PAID"
	order.PaymentMethod = req.Value.PaymentMethod
	order.TransactionUUID = res.TransactionUuid
	return &ordersv1.OrderPayResponse{TransactionUUID: order.TransactionUUID}, nil
}
