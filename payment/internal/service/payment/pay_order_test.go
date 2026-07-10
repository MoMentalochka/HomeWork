package payment

import (
	"github.com/MoMentalochka/HomeWork/payment/internal/model"
	"github.com/google/uuid"
)

func (s *ServiceSuite) TestPayOrderValidUuid() {

	res := s.service.PayOrder(s.ctx, &model.PayOrderRequest{})

	s.Require().NoError(uuid.Validate(res.TransactionUuid))
}
