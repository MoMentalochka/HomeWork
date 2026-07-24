package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	inventoryApi "github.com/MoMentalochka/HomeWork/inventory/internal/api/inventory/v1"
	"github.com/MoMentalochka/HomeWork/inventory/internal/config"
	inventoryRepository "github.com/MoMentalochka/HomeWork/inventory/internal/repository/inventory"
	inventoryService "github.com/MoMentalochka/HomeWork/inventory/internal/service/inventory"
	inventoryV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
)

const configPath = "../../deploy/compose/inventory/.env"

func main() {
	if err := config.Load(configPath); err != nil {
		log.Printf("failed to load config: %v", err)
		return
	}
	grpcAddress := config.AppConfig().InventoryGRPC.Address()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// подключение к mongo
	mongoClient, err := mongo.Connect(options.Client().ApplyURI(config.AppConfig().Mongo.URI()))
	defer func() {
		err = mongoClient.Disconnect(ctx)
		if err != nil {
			log.Printf("Error closing connection: %s\n", err)
		}
	}()

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Printf("Error pinging database: %s\n", err)
		return
	}

	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return
	}
	defer func() {
		cerr := lis.Close()
		if cerr != nil {
			log.Printf("Error closing listener: %s\n", cerr)
		}
	}()

	s := grpc.NewServer()

	repo := inventoryRepository.NewRepository(mongoClient)
	service := inventoryService.NewService(repo)
	api := inventoryApi.NewApi(service)

	inventoryV1.RegisterInventoryServiceServer(s, api)

	reflection.Register(s)

	go func() {
		log.Printf("Starting Inventory gRPC server on port %s\n", grpcAddress)
		err = s.Serve(lis)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("🛑 Gracefully shutdown order server %s\n", grpcAddress)
	s.GracefulStop()
	log.Printf("✅ Inventory server stopped")
}
