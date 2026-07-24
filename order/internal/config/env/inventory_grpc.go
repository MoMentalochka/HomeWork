package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type inventoryGrpcEnvConfig struct {
	Host string `env:"GRPC_HOST,required"`
	Port string `env:"INVENTORY_PORT,required"`
}

type inventoryGrpcConfig struct {
	cfg inventoryGrpcEnvConfig
}

func NewInventoryGrpcConfig() (*inventoryGrpcConfig, error) {
	var cfg inventoryGrpcEnvConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &inventoryGrpcConfig{cfg: cfg}, nil
}

func (p *inventoryGrpcConfig) Address() string {
	return fmt.Sprintf(
		"%s:%s",
		p.cfg.Host,
		p.cfg.Port,
	)
}
