package v1

import (
	"context"
	"fmt"

	"github.com/MoMentalochka/HomeWork/payment/internal/converter"
	"github.com/MoMentalochka/HomeWork/payment/internal/service"
	payment_v1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/payment/v1"
)

type api struct {
	payment_v1.UnimplementedPaymentServiceServer

	paymentService service.PaymentService
}

func NewApi(service service.PaymentService) *api {
	return &api{
		paymentService: service,
	}
}

func (a *api) PayOrder(ctx context.Context, req *payment_v1.PayOrderRequest) (*payment_v1.PayOrderResponse, error) {
	res, err := a.paymentService.PayOrder(ctx, converter.ProtoPayOrderRequestToModel(req))
	if err != nil {
		return &payment_v1.PayOrderResponse{}, fmt.Errorf("error calling service.PayOrder: %w", err)
	}
	return res, nil
}
