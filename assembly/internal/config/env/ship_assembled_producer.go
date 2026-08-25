package env

import (
	"github.com/IBM/sarama"
	"github.com/caarlos0/env/v11"
)

type ShipAssembledProducerEnvConfig struct {
	Topic string `env:"ORDER_ASSEMBLED_TOPIC_NAME,required"`
}

type ShipAssembledProducerConfig struct {
	raw ShipAssembledProducerEnvConfig
}

func NewShipAssembledProducer() (*ShipAssembledProducerConfig, error) {
	var raw ShipAssembledProducerEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &ShipAssembledProducerConfig{raw: raw}, nil
}

func (cfg *ShipAssembledProducerConfig) Topic() string {
	return cfg.raw.Topic
}

func (cfg *ShipAssembledProducerConfig) Config() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V4_0_0_0
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	return config
}
