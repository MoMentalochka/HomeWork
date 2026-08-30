package app

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"

	"github.com/MoMentalochka/HomeWork/assembly/internal/config"
	kafkaConverter "github.com/MoMentalochka/HomeWork/assembly/internal/converter/kafka"
	"github.com/MoMentalochka/HomeWork/assembly/internal/converter/kafka/decoder"
	"github.com/MoMentalochka/HomeWork/assembly/internal/service"
	assemblyConsumer "github.com/MoMentalochka/HomeWork/assembly/internal/service/consumer/assembly_consumer"
	assemblyProducer "github.com/MoMentalochka/HomeWork/assembly/internal/service/producer/assembly_producer"
	"github.com/MoMentalochka/HomeWork/platform/pkg/closer"
	wrappedKafka "github.com/MoMentalochka/HomeWork/platform/pkg/kafka"
	wrappedKafkaConsumer "github.com/MoMentalochka/HomeWork/platform/pkg/kafka/consumer"
	wrappedKafkaProducer "github.com/MoMentalochka/HomeWork/platform/pkg/kafka/producer"
	"github.com/MoMentalochka/HomeWork/platform/pkg/logger"
	kafkaMiddleware "github.com/MoMentalochka/HomeWork/platform/pkg/middleware/kafka"
)

type diContainer struct {
	assemblyConsumerService service.OrderPaidConsumerService
	orderPaidConsumer       wrappedKafka.Consumer

	shipAssembledProducer   wrappedKafka.Producer
	assemblyProducerService service.AssemblyProducerService

	syncProducer  sarama.SyncProducer
	consumerGroup sarama.ConsumerGroup

	orderDecoder kafkaConverter.OrderPaidDecoder
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) AssemblyProducerService() service.AssemblyProducerService {
	if d.assemblyProducerService == nil {
		d.assemblyProducerService = assemblyProducer.NewService(d.ShipAssembledProducer())
	}

	return d.assemblyProducerService
}

func (d *diContainer) SyncProducer() sarama.SyncProducer {
	if d.syncProducer == nil {
		p, err := sarama.NewSyncProducer(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().OrderPaidConsumer.Config(),
		)
		if err != nil {
			panic(fmt.Sprintf("failed to create ASSEMBLY sync producer: %s\n", err.Error()))
		}
		closer.AddNamed("Kafka sync producer", func(ctx context.Context) error {
			return p.Close()
		})

		d.syncProducer = p
	}

	return d.syncProducer
}

func (d *diContainer) ShipAssembledProducer() wrappedKafka.Producer {
	if d.shipAssembledProducer == nil {
		d.shipAssembledProducer = wrappedKafkaProducer.NewProducer(
			d.SyncProducer(),
			config.AppConfig().ShipAssembledProducer.Topic(),
			logger.Logger(),
		)
	}

	return d.shipAssembledProducer
}

func (d *diContainer) ConsumerGroup() sarama.ConsumerGroup {
	if d.consumerGroup == nil {
		consumerGroup, err := sarama.NewConsumerGroup(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().OrderPaidConsumer.GroupID(),
			config.AppConfig().ShipAssembledProducer.Config(),
		)
		if err != nil {
			panic(fmt.Sprintf("failed to create consumer group: %s\n", err.Error()))
		}
		closer.AddNamed("Kafka consumer group", func(ctx context.Context) error {
			return d.consumerGroup.Close()
		})

		d.consumerGroup = consumerGroup
	}

	return d.consumerGroup
}

func (d *diContainer) OrderPaidConsumerService() service.OrderPaidConsumerService {
	if d.assemblyConsumerService == nil {
		d.assemblyConsumerService = assemblyConsumer.NewService(d.OrderPaidConsumer(), d.AssemblyProducerService(), d.OrderDecoder())
	}

	return d.assemblyConsumerService
}

func (d *diContainer) OrderPaidConsumer() wrappedKafka.Consumer {
	if d.orderPaidConsumer == nil {
		d.orderPaidConsumer = wrappedKafkaConsumer.NewConsumer(
			d.ConsumerGroup(),
			[]string{
				config.AppConfig().OrderPaidConsumer.Topic(),
			},
			logger.Logger(),
			kafkaMiddleware.Logging(logger.Logger()),
		)
	}

	return d.orderPaidConsumer
}

func (d *diContainer) OrderDecoder() kafkaConverter.OrderPaidDecoder {
	if d.orderDecoder == nil {
		d.orderDecoder = decoder.NewOrderPaidDecoder()
	}

	return d.orderDecoder
}
