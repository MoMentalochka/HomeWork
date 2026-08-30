package assembly_consumer

import (
	"context"

	"go.uber.org/zap"

	kafkaConverter "github.com/MoMentalochka/HomeWork/order/internal/converter/kafka"
	"github.com/MoMentalochka/HomeWork/order/internal/repository"
	def "github.com/MoMentalochka/HomeWork/order/internal/service"
	"github.com/MoMentalochka/HomeWork/platform/pkg/kafka"
	"github.com/MoMentalochka/HomeWork/platform/pkg/logger"
)

var _ def.OrderAssembledConsumerService = (*service)(nil)

type service struct {
	repository            repository.OrderRepository
	orderAssembleConsumer kafka.Consumer
	orderAssembleDecoder  kafkaConverter.OrderAssembledDecoder
}

func NewService(orderAssembleConsumer kafka.Consumer, repository repository.OrderRepository, orderAssembleDecoder kafkaConverter.OrderAssembledDecoder) *service {
	return &service{
		orderAssembleConsumer: orderAssembleConsumer,
		orderAssembleDecoder:  orderAssembleDecoder,
		repository:            repository,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	err := s.orderAssembleConsumer.Consume(ctx, s.OrderAssembledHandler)
	if err != nil {
		logger.Error(ctx, "Consume from order.assembled topic error", zap.Error(err))
		return err
	}

	return nil
}
