package v1

import (
	"context"

	"github.com/MoMentalochka/HomeWork/payment/internal/converter"
	payment_v1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/payment/v1"
)

func (a *api) PayOrder(ctx context.Context, req *payment_v1.PayOrderRequest) (*payment_v1.PayOrderResponse, error) {
	res := a.paymentService.PayOrder(ctx, converter.ProtoPayOrderRequestToModel(req))
	return res, nil
}
