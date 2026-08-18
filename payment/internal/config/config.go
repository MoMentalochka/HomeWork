package config

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/MoMentalochka/HomeWork/payment/internal/config/env"
)

var appConfig *config

type config struct {
	PaymentGRPC PaymentGRPCConfig
	Logger      LoggerConfig
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		log.Println("failed to load env file")
		return err
	}
	paymentGRPCCfg, err := env.NewPaymentGRPCConfig()
	if err != nil {
		log.Println("failed to load grpc config")
		return err
	}
	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		PaymentGRPC: paymentGRPCCfg,
		Logger:      loggerCfg,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}
