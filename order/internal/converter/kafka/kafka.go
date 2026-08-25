package kafka

import "github.com/MoMentalochka/HomeWork/order/internal/model"

type OrderAssembledDecoder interface {
	Decode(data []byte) (model.OrderAssembledEvent, error)
}
