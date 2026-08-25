package config

import (
	"github.com/MoMentalochka/HomeWork/assembly/internal/config/env"
	"github.com/joho/godotenv"
)

var appConfig *config

type config struct {
	Logger                LoggerConfig
	Kafka                 KafkaConfig
	OrderPaidConsumer     OrderPaidConsumerConfig
	ShipAssembledProducer ShipAssembledProducerConfig
}

func Load(path string) error {
	if err := godotenv.Load(path); err != nil {
		return err
	}
	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	kafkaCfg, err := env.NewKafkaConfig()
	if err != nil {
		return err
	}

	orderPaidConsumerCfg, err := env.NewOrderPaidConsumerConfig()
	if err != nil {
		return err
	}

	shipAssembledProducerCfg, err := env.NewShipAssembledProducer()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:                loggerCfg,
		Kafka:                 kafkaCfg,
		OrderPaidConsumer:     orderPaidConsumerCfg,
		ShipAssembledProducer: shipAssembledProducerCfg,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}
