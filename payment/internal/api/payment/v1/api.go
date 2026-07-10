package v1

import (
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
