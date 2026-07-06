package payment

import (
	"context"
	"log"

	"github.com/MoMentalochka/HomeWork/payment/internal/model"
	payment_v1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/payment/v1"
	"github.com/google/uuid"
)

type service struct {
}

func NewService() *service {
	return &service{}
}

func (s *service) PayOrder(_ context.Context, _ *model.PayOrderRequest) (*payment_v1.PayOrderResponse, error) {
	id := uuid.New().String()
	log.Printf("Оплата прошла успешно, transaction_uuid: %s", id)

	return &payment_v1.PayOrderResponse{TransactionUuid: id}, nil
}
