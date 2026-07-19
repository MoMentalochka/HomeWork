package v1

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/MoMentalochka/HomeWork/payment/internal/converter"
	payment_v1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/payment/v1"
)

func (s *APISuite) TestPayOrderCorrectUuid() {
	var (
		request = &payment_v1.PayOrderRequest{
			PaymentMethod: payment_v1.PaymentMethod_PAYMENT_METHOD_CARD,
		}
		model  = converter.ProtoPayOrderRequestToModel(request)
		fakeId = gofakeit.UUID()
	)

	s.service.On("PayOrder", s.ctx, model).Return(&payment_v1.PayOrderResponse{TransactionUuid: fakeId})

	res, _ := s.api.PayOrder(s.ctx, request)

	s.Require().Equal(fakeId, res.TransactionUuid)
}
