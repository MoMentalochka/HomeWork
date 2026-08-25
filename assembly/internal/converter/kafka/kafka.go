package kafka

import "github.com/MoMentalochka/HomeWork/assembly/internal/model"

type OrderPaidDecoder interface {
	Decode(data []byte) (model.OrderPaidEvent, error)
}
