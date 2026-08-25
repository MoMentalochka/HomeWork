package env

import (
	"github.com/IBM/sarama"
	"github.com/caarlos0/env/v11"
)

type OrderAssembledConsumerEnvConfig struct {
	TopicName string `env:"ORDER_ASSEMBLED_TOPIC_NAME,required"`
	GroupID   string `env:"ORDER_ASSEMBLED_CONSUMER_GROUP_ID,required"`
}

type OrderAssembledConsumerConfig struct {
	raw OrderAssembledConsumerEnvConfig
}

func NewOrderAssembledConsumerConfig() (*OrderAssembledConsumerConfig, error) {
	var raw OrderAssembledConsumerEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &OrderAssembledConsumerConfig{raw: raw}, nil
}

func (cfg *OrderAssembledConsumerConfig) Topic() string {
	return cfg.raw.TopicName
}

func (cfg *OrderAssembledConsumerConfig) GroupID() string {
	return cfg.raw.GroupID
}

// Config возвращает конфигурацию для sarama consumer
func (cfg *OrderAssembledConsumerConfig) Config() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V4_0_0_0
	config.Producer.Return.Successes = true

	return config
}
