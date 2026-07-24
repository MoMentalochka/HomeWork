package config

import (
	"log"

	"github.com/MoMentalochka/HomeWork/inventory/internal/config/env"
	"github.com/joho/godotenv"
)

var appConfig *config

type config struct {
	InventoryGRPC InventoryGRPCConfig
	Mongo         MongoConfig
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		log.Println("failed to load env file")
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
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}
