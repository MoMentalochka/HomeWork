package env

import (
	"github.com/IBM/sarama"
	"github.com/caarlos0/env/v11"
)

type OrderPaidConsumerEnvConfig struct {
	TopicName string `env:"ORDER_PAID_TOPIC_NAME,required"`
	GroupID   string `env:"ORDER_PAID_CONSUMER_GROUP_ID,required"`
}

type OrderPaidConsumerConfig struct {
	raw OrderPaidConsumerEnvConfig
}

func NewOrderPaidConsumerConfig() (*OrderPaidConsumerConfig, error) {
	var raw OrderPaidConsumerEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &OrderPaidConsumerConfig{raw: raw}, nil
}

func (cfg *OrderPaidConsumerConfig) Topic() string {
	return cfg.raw.TopicName
}

func (cfg *OrderPaidConsumerConfig) GroupID() string {
	return cfg.raw.GroupID
}

// Config возвращает конфигурацию для sarama consumer
func (cfg *OrderPaidConsumerConfig) Config() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V4_0_0_0
	config.Producer.Return.Successes = true

	return config
}
