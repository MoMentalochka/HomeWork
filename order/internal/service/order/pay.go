package order

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/MoMentalochka/HomeWork/order/internal/model"
	ordersv1 "github.com/MoMentalochka/HomeWork/shared/pkg/openapi/order/v1"
	paymentv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/payment/v1"
)

func (s *orderService) OrderPay(ctx context.Context, req ordersv1.OptOrderPayRequest, uuid string) (ordersv1.OrderPayRes, error) {
	order, err := s.store.Get(uuid)
	if err != nil {
		return &ordersv1.OrderPayResponse{}, err
	}

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

	err = s.store.Update(uuid, order)
	if err != nil {
		return &ordersv1.OrderPayResponse{}, err
	}

	err = s.producerService.ProduceOrderPaid(ctx, model.OrderPaidEvent{
		OrderUUID:       order.OrderUUID,
		UserUUID:        order.UserUUID,
		TransactionUUID: transactionUuid,
		EventUUID:       gofakeit.UUID(),
		PaymentMethod:   string(req.Value.PaymentMethod),
	})
	if err != nil {
		return &ordersv1.OrderPayResponse{}, err
	}

	return &ordersv1.OrderPayResponse{TransactionUUID: order.TransactionUUID}, nil
}
