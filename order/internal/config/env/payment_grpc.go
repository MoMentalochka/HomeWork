package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type paymentGrpcEnvConfig struct {
	Host string `env:"GRPC_HOST,required"`
	Port string `env:"PAYMENT_PORT,required"`
}

type paymentGrpcConfig struct {
	cfg paymentGrpcEnvConfig
}

func NewPaymentGrpcConfig() (*paymentGrpcConfig, error) {
	var cfg paymentGrpcEnvConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &paymentGrpcConfig{cfg: cfg}, nil
}

func (p *paymentGrpcConfig) Address() string {
	return fmt.Sprintf(
		"%s:%s",
		p.cfg.Host,
		p.cfg.Port,
	)
}
