package converter

import (
	"github.com/MoMentalochka/HomeWork/order/internal/model"
	orders_v1 "github.com/MoMentalochka/HomeWork/shared/pkg/openapi/order/v1"
)

func ProtoStatusToModel(status orders_v1.Status) string {
	return string(status)
}

func ModelDtoToProtoDto(dto *model.OrderDto) *orders_v1.OrderDto {
	return &orders_v1.OrderDto{
		OrderUUID:       dto.OrderUUID,
		UserUUID:        dto.UserUUID,
		PartUuids:       dto.PartUuids,
		TotalPrice:      dto.TotalPrice,
		TransactionUUID: dto.TransactionUUID,
		PaymentMethod:   orders_v1.PaymentMethod(dto.PaymentMethod),
		Status:          orders_v1.Status(dto.Status),
	}
}
