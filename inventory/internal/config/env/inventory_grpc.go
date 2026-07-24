package env

import (
	"net"

	"github.com/caarlos0/env/v11"
)

type inventoryGRPCEnvConfig struct {
	Host string `env:"GRPC_HOST,required"`
	Port string `env:"GRPC_PORT,required"`
}

type inventoryGRPCConfig struct {
	cfg inventoryGRPCEnvConfig
}

func NewInventoryGRPCConfig() (*inventoryGRPCConfig, error) {
	var cfg inventoryGRPCEnvConfig
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &inventoryGRPCConfig{cfg: cfg}, nil
}

func (ic *inventoryGRPCConfig) Address() string {
	return net.JoinHostPort(ic.cfg.Host, ic.cfg.Port)
}
