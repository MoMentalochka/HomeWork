package env

import (
	"net"

	"github.com/caarlos0/env/v11"
)

type paymentGRPCEnvConfig struct {
	Host string `env:"GRPC_HOST,required"`
	Port string `env:"GRPC_PORT,required"`
}

type paymentGRPCConfig struct {
	cfg paymentGRPCEnvConfig
}

func NewPaymentGRPCConfig() (*paymentGRPCConfig, error) {
	var cfg paymentGRPCEnvConfig
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &paymentGRPCConfig{cfg: cfg}, nil
}

func (ic *paymentGRPCConfig) Address() string {
	return net.JoinHostPort(ic.cfg.Host, ic.cfg.Port)
}
