package config

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/MoMentalochka/HomeWork/order/internal/config/env"
)

var appConfig *config

type config struct {
	Http      OrderHttpConfig
	Postgres  PostgresConfig
	Payment   PaymentConfig
	Inventory InventoryConfig
}

func Load(path string) error {
	if err := godotenv.Load(path); err != nil {
		return err
	}

	orderCfg, err := env.NewOrderHTTPConfig()
	if err != nil {
		log.Println("failed to load order http config")
		return err
	}
	postgresCfg, err := env.NewPostgresConfig()
	if err != nil {
		log.Println("failed to load order http config")
		return err
	}
	inventoryCfg, err := env.NewInventoryGrpcConfig()
	if err != nil {
		log.Println("failed to load order http config")
		return err
	}
	paymentCfg, err := env.NewPaymentGrpcConfig()
	if err != nil {
		log.Println("failed to load order http config")
		return err
	}

	appConfig = &config{Http: orderCfg, Postgres: postgresCfg, Payment: paymentCfg, Inventory: inventoryCfg}

	return nil
}

func AppConfig() *config {
	return appConfig
}
