package grpc

import (
	"context"

	"github.com/MoMentalochka/HomeWork/order/internal/model"
	paymentv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/payment/v1"
)

type PaymentClient interface {
	PayOrder(ctx context.Context, req *paymentv1.PayOrderRequest) (string, error)
}

type InventoryClient interface {
	GetPart(ctx context.Context, uuid string) (model.Part, error)
}
