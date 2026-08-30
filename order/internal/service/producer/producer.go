package producer

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/MoMentalochka/HomeWork/order/internal/model"
	"github.com/MoMentalochka/HomeWork/platform/pkg/kafka"
	"github.com/MoMentalochka/HomeWork/platform/pkg/logger"
	events_v1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/events/v1"
)

type service struct {
	orderRecordedProducer kafka.Producer
}

func NewService(orderRecordedProducer kafka.Producer) *service {
	return &service{
		orderRecordedProducer: orderRecordedProducer,
	}
}

func (p *service) ProduceOrderPaid(ctx context.Context, event model.OrderPaidEvent) error {
	msg := &events_v1.OrderPaid{
		OrderUuid:       event.OrderUUID,
		EventUuid:       event.EventUUID,
		UserUuid:        event.UserUUID,
		PaymentMethod:   event.PaymentMethod,
		TransactionUuid: event.TransactionUUID,
	}

	payload, err := proto.Marshal(msg)
	if err != nil {
		logger.Error(ctx, "failed to marshal UFORecorded", zap.Error(err))
		return err
	}

	err = p.orderRecordedProducer.Send(ctx, []byte(event.EventUUID), payload)
	if err != nil {
		logger.Error(ctx, "failed to publish UFORecorded", zap.Error(err))
		return err
	}

	return nil
}
