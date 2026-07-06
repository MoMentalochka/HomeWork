package converter

import (
	"github.com/MoMentalochka/HomeWork/payment/internal/model"
	payment_v1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/payment/v1"
)

func ProtoPayOrderRequestToModel(req *payment_v1.PayOrderRequest) *model.PayOrderRequest {
	return &model.PayOrderRequest{
		OrderUuid:     req.OrderUuid,
		UserUuid:      req.UserUuid,
		PaymentMethod: req.PaymentMethod.String(),
	}
}

func ModelPayOrderResponse(res *model.PayOrderResponse) *payment_v1.PayOrderResponse {
	return &payment_v1.PayOrderResponse{
		TransactionUuid: res.TransactionUuid,
	}
}
