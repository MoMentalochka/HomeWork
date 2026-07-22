package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type mongoEnvConfig struct {
	Host     string `env:"MONGO_HOST,required"`
	Port     string `env:"MONGO_PORT,required"`
	Database string `env:"MONGO_DATABASE,required"`
	AuthDB   string `env:"MONGO_AUTH_DB,required"`
	User     string `env:"MONGO_INITDB_ROOT_USERNAME,required"`
	Password string `env:"MONGO_INITDB_ROOT_PASSWORD,required"`
}

type mongoConfig struct {
	cfg mongoEnvConfig
}

func NewMongoConfig() (*mongoConfig, error) {
	var cfg mongoEnvConfig
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &mongoConfig{cfg: cfg}, nil
}

func (mc *mongoConfig) URI() string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/%s?authSource=%s",
		mc.cfg.User,
		mc.cfg.Password,
		mc.cfg.Host,
		mc.cfg.Port,
		mc.cfg.Database,
		mc.cfg.AuthDB,
	)
}

func (mc *mongoConfig) DB_Name() string {
	return mc.cfg.Database
}
