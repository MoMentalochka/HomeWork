package payment

import (
	"github.com/google/uuid"

	"github.com/MoMentalochka/HomeWork/payment/internal/model"
)

func (s *ServiceSuite) TestPayOrderValidUuid() {
	res := s.service.PayOrder(s.ctx, &model.PayOrderRequest{})

	s.Require().NoError(uuid.Validate(res.TransactionUuid))
}
