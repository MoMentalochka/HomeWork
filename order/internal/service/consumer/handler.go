package assembly_consumer

import (
	"context"

	"go.uber.org/zap"

	"github.com/MoMentalochka/HomeWork/platform/pkg/kafka"
	"github.com/MoMentalochka/HomeWork/platform/pkg/logger"
)

func (s *service) OrderAssembledHandler(ctx context.Context, msg kafka.Message) error {
	event, err := s.orderAssembleDecoder.Decode(msg.Value)
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
		zap.Int64("build_time_sec", event.BuildTimeSec),
	)
	order, err := s.repository.Get(event.OrderUUID)
	if err != nil {
		logger.Error(ctx, "Failed to Get order", zap.Error(err))
		return err
	}

	if order == nil {
		logger.Info(ctx, "Order not Found")
		return err
	}

	if order.Status == "PAID" {
		order.Status = "COMPLETED"
		err = s.repository.Update(event.OrderUUID, order)
		return err
	}

	return nil
}
