package env

import (
	"net"

	"github.com/caarlos0/env/v11"
)

type inventoryGRPCEnvCongig struct {
	Host string `env:"GRPC_HOST,required"`
	Port string `env:"GRPC_PORT,required"`
}

type inventoryGRPCConfig struct {
	cfg inventoryGRPCEnvCongig
}

func NewInvenoryGRPCConfig() (*inventoryGRPCConfig, error) {
	var cfg inventoryGRPCEnvCongig
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &inventoryGRPCConfig{cfg: cfg}, nil
}

func (ic *inventoryGRPCConfig) Address() string {
	return net.JoinHostPort(ic.cfg.Host, ic.cfg.Port)
}
