package config

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/MoMentalochka/HomeWork/inventory/internal/config/env"
)

var appConfig *config

type config struct {
	InventoryGRPC InventoryGRPCConfig
	Mongo         MongoConfig
	Logger        LoggerConfig
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		log.Println("failed to load env file")
		return err
	}
	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}
	mongoCfg, err := env.NewMongoConfig()
	if err != nil {
		log.Println("failed to load mongo config")
		return err
	}

	inventoryGRPCCfg, err := env.NewInventoryGRPCConfig()
	if err != nil {
		log.Println("failed to load grpc config")
		return err
	}

	appConfig = &config{
		Mongo:         mongoCfg,
		InventoryGRPC: inventoryGRPCCfg,
		Logger:        loggerCfg,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}
