package decoder

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/MoMentalochka/HomeWork/assembly/internal/model"
	events_v1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/events/v1"
)

type decoder struct{}

func NewOrderPaidDecoder() *decoder {
	return &decoder{}
}

func (d *decoder) Decode(data []byte) (model.OrderPaidEvent, error) {
	var pb events_v1.OrderPaid
	if err := proto.Unmarshal(data, &pb); err != nil {
		return model.OrderPaidEvent{}, fmt.Errorf("failed to unmarshal protobuf: %w", err)
	}

	return model.OrderPaidEvent{
		OrderUUID:       pb.GetOrderUuid(),
		EventUUID:       pb.GetEventUuid(),
		UserUUID:        pb.GetUserUuid(),
		TransactionUUID: pb.GetTransactionUuid(),
		PaymentMethod:   pb.GetPaymentMethod(),
	}, nil
}
