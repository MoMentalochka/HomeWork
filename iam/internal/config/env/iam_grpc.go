package env

import (
	"net"

	"github.com/caarlos0/env/v11"
)

type iamGRPCEnvConfig struct {
	Host string `env:"GRPC_HOST,required"`
	Port string `env:"GRPC_PORT,required"`
}

type iamGRPCConfig struct {
	cfg iamGRPCEnvConfig
}

func NewIamGRPCConfig() (*iamGRPCConfig, error) {
	var cfg iamGRPCEnvConfig
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &iamGRPCConfig{cfg: cfg}, nil
}

func (ic *iamGRPCConfig) Address() string {
	return net.JoinHostPort(ic.cfg.Host, ic.cfg.Port)
}
