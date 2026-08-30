package assembly_consumer

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/MoMentalochka/HomeWork/assembly/internal/model"
	"github.com/MoMentalochka/HomeWork/platform/pkg/kafka"
	"github.com/MoMentalochka/HomeWork/platform/pkg/logger"
)

func (s *service) OrderPaidHandler(ctx context.Context, msg kafka.Message) error {
	event, err := s.orderPaidDecoder.Decode(msg.Value)
	if err != nil {
		logger.Error(ctx, "Failed to decode OrderPaid", zap.Error(err))
		return err
	}

	logger.Info(ctx, "Processing message",
		zap.String("topic", msg.Topic),
		zap.Any("partition", msg.Partition),
		zap.Any("offset", msg.Offset),
		zap.String("event_uuid", event.EventUUID),
		zap.String("order_uuid", event.OrderUUID),
		zap.String("user_uuid", event.UserUUID),
		zap.String("payment_method", event.PaymentMethod),
		zap.String("transaction_uuid", event.TransactionUUID),
	)

	go func() {
		timer := time.NewTimer(10 * time.Second)
		<-timer.C

		err = s.orderAssembleProducer.ProduceShipAssembled(ctx, model.ShipAssembledEvent{
			OrderUUID:    event.OrderUUID,
			UserUUID:     event.UserUUID,
			EventUUID:    event.EventUUID,
			BuildTimeSec: 10,
		},
		)
		if err != nil {
			logger.Error(ctx, "Failed to assembly ship", zap.Error(err))
		}
	}()

	return nil
}
