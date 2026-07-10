package v1

import (
	"context"
	"errors"
	"fmt"

	"github.com/MoMentalochka/HomeWork/order/internal/converter"
	"github.com/MoMentalochka/HomeWork/order/internal/model"
	"github.com/MoMentalochka/HomeWork/order/internal/service"
	ordersv1 "github.com/MoMentalochka/HomeWork/shared/pkg/openapi/order/v1"
)

type orderApi struct {
	orderService service.OrderService
}

func NewOrderApi(service service.OrderService) *orderApi {
	return &orderApi{
		orderService: service,
	}
}

func (a *orderApi) CreateNewOrder(ctx context.Context, req *ordersv1.CreateOrderRequest) (*ordersv1.CreateOrderResponse, error) {

	res, err := a.orderService.CreateNewOrder(ctx,
		model.CreateOrderRequest{
			UserUUID:  req.UserUUID,
			PartUuids: req.PartUuids,
		},
	)
	if err != nil {
		return &ordersv1.CreateOrderResponse{}, err
	}
	return res, nil
}

func (a *orderApi) GetOrderById(ctx context.Context, params ordersv1.GetOrderByIdParams) (ordersv1.GetOrderByIdRes, error) {
	order, err := a.orderService.GetOrderById(ctx, params.OrderUUID)
	if errors.Is(err, model.ErrPartNotFound) {
		return &ordersv1.NotFound{
			Code:    404,
			Message: fmt.Sprintf("order with uuid %s not found", params.OrderUUID),
		}, nil
	}
	return converter.ModelDtoToProtoDto(order), err
}

func (a *orderApi) OrderCancel(ctx context.Context, params ordersv1.OrderCancelParams) (ordersv1.OrderCancelRes, error) {
	res, err := a.orderService.OrderCancel(ctx, params.OrderUUID)
	if err != nil {
		return &ordersv1.OrderCancelResponse{}, nil
	}
	return res, nil
}

func (a *orderApi) OrderPay(ctx context.Context, req ordersv1.OptOrderPayRequest, params ordersv1.OrderPayParams) (ordersv1.OrderPayRes, error) {
	res, err := a.orderService.OrderPay(ctx, req, params.OrderUUID)
	if err != nil {
		return &ordersv1.OrderPayResponse{}, nil
	}
	return res, nil
}
