package app

import (
	"context"
	"fmt"

	inventoryV1API "github.com/MoMentalochka/HomeWork/inventory/internal/api/inventory/v1"
	"github.com/MoMentalochka/HomeWork/inventory/internal/config"
	"github.com/MoMentalochka/HomeWork/inventory/internal/repository"
	inventoryRepository "github.com/MoMentalochka/HomeWork/inventory/internal/repository/inventory"
	"github.com/MoMentalochka/HomeWork/inventory/internal/service"
	inventoryService "github.com/MoMentalochka/HomeWork/inventory/internal/service/inventory"
	"github.com/MoMentalochka/HomeWork/platform/pkg/closer"
	inventoryV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type diContainer struct {
	inventoryV1API   inventoryV1.InventoryServiceServer
	inventoryService service.InventoryService
	inventoryRepo    repository.InventoryRepository

	mongoDBClient *mongo.Client
	mongoDBHandle *mongo.Database
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}
func (d *diContainer) InventoryV1API(ctx context.Context) inventoryV1.InventoryServiceServer {
	if d.inventoryV1API == nil {
		d.inventoryV1API = inventoryV1API.NewApi(d.InventoryService(ctx))
	}

	return d.inventoryV1API
}

func (d *diContainer) InventoryService(ctx context.Context) service.InventoryService {
	if d.inventoryService == nil {
		d.inventoryService = inventoryService.NewService(d.InventoryRepository(ctx))
	}
	return d.inventoryService
}

func (d *diContainer) InventoryRepository(ctx context.Context) repository.InventoryRepository {
	if d.inventoryRepo == nil {
		d.inventoryRepo = inventoryRepository.NewRepository(d.MongoDBClient(ctx))
	}
	return d.inventoryRepo
}

func (d *diContainer) MongoDBClient(ctx context.Context) *mongo.Client {
	if d.mongoDBClient == nil {
		client, err := mongo.Connect(options.Client().ApplyURI(config.AppConfig().Mongo.URI()))
		if err != nil {
			panic(fmt.Sprintf("failed to connect to MongoDB: %s\n", err.Error()))
		}

		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			panic(fmt.Sprintf("failed to ping MongoDB: %v\n", err))
		}

		closer.AddNamed("MongoDB client", func(ctx context.Context) error {
			return client.Disconnect(ctx)
		})

		d.mongoDBClient = client
	}

	return d.mongoDBClient
}
