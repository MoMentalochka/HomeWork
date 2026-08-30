package assembly_consumer

import (
	"context"

	"go.uber.org/zap"

	kafkaConverter "github.com/MoMentalochka/HomeWork/assembly/internal/converter/kafka"
	def "github.com/MoMentalochka/HomeWork/assembly/internal/service"
	"github.com/MoMentalochka/HomeWork/platform/pkg/kafka"
	"github.com/MoMentalochka/HomeWork/platform/pkg/logger"
)

var _ def.OrderPaidConsumerService = (*service)(nil)

type service struct {
	orderPaidConsumer     kafka.Consumer
	orderAssembleProducer def.AssemblyProducerService
	orderPaidDecoder      kafkaConverter.OrderPaidDecoder
}

func NewService(orderPaidConsumer kafka.Consumer, orderAssembleProducer def.AssemblyProducerService, orderPaidDecoder kafkaConverter.OrderPaidDecoder) *service {
	return &service{
		orderPaidConsumer:     orderPaidConsumer,
		orderAssembleProducer: orderAssembleProducer,
		orderPaidDecoder:      orderPaidDecoder,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	err := s.orderPaidConsumer.Consume(ctx, s.OrderPaidHandler)
	if err != nil {
		logger.Error(ctx, "Consume from order.paid topic error", zap.Error(err))
		return err
	}

	return nil
}
