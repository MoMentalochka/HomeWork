package service

import (
	"context"

	"github.com/MoMentalochka/HomeWork/payment/internal/model"
	payment_v1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/payment/v1"
)

//go:generate ../../../bin/mockery --case=underscore --all

type PaymentService interface {
	PayOrder(ctx context.Context, data *model.PayOrderRequest) (*payment_v1.PayOrderResponse, error)
}
