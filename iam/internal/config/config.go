package config

import (
	"log"

	"github.com/MoMentalochka/HomeWork/iam/internal/config/env"

	"github.com/joho/godotenv"
)

var appConfig *config

type config struct {
	IamGRPC  IamGRPCConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Logger   LoggerConfig
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		log.Println("failed to load env file")
		return err
	}

	iamGRPCCfg, err := env.NewIamGRPCConfig()
	if err != nil {
		return err
	}

	postgresCfg, err := env.NewPostgresConfig()
	if err != nil {
		log.Println("failed to load order http config")
		return err
	}

	redisCfg, err := env.NewRedisConfig()
	if err != nil {
		return err
	}

	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	appConfig = &config{IamGRPC: iamGRPCCfg, Postgres: postgresCfg, Redis: redisCfg, Logger: loggerCfg}

	return nil
}

func AppConfig() *config {
	return appConfig
}
