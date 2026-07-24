package env

import (
	"log"
	"net"
	"time"

	"github.com/caarlos0/env/v11"
)

type orderHttpEnvConfig struct {
	Host            string `env:"HTTP_HOST,required"`
	Port            string `env:"HTTP_PORT,required"`
	ReadTimeOut     string `env:"HTTP_READ_TIMEOUT,required"`
	ShutdownTimeout string `env:"HTTP_SHUTDOWN_TIMEOUT"`
}

type httpConfig struct {
	cfg orderHttpEnvConfig
}

func NewOrderHTTPConfig() (*httpConfig, error) {
	var cfg orderHttpEnvConfig

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &httpConfig{cfg: cfg}, nil
}

func (h *httpConfig) Address() string {
	return net.JoinHostPort(h.cfg.Host, h.cfg.Port)
}

func (h *httpConfig) Port() string {
	return h.cfg.Port
}

func (h *httpConfig) ReadTimeOut() time.Duration {
	dur, err := time.ParseDuration(h.cfg.ReadTimeOut)
	if err != nil {
		log.Println("Error parsing ReadTimeOut, used default 5s timeout", err)
		return 5 * time.Second
	}
	return dur
}
