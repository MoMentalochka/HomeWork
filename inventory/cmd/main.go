package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	inventoryApi "github.com/MoMentalochka/HomeWork/inventory/internal/api/inventory/v1"
	inventoryRepository "github.com/MoMentalochka/HomeWork/inventory/internal/repository/inventory"
	inventoryService "github.com/MoMentalochka/HomeWork/inventory/internal/service/inventory"
	inventoryV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
)

const grpcPort = 50051

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
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
	// подключение к mongo
	ctx := context.Background()

	err = godotenv.Load("../.env")
	if err != nil {
		log.Println("Error with loading env file")
		return
	}

	dbUri := os.Getenv("MONGO_URI")

	client, err := mongo.Connect(options.Client().ApplyURI(dbUri))

	defer func() {
		err = client.Disconnect(ctx)
		if err != nil {
			log.Printf("Error closing connection: %s\n", err)
		}
	}()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Error pinging database: %s\n", err)
		return
	}

	repo := inventoryRepository.NewRepository(client.Database("inventory"))
	service := inventoryService.NewService(repo)
	api := inventoryApi.NewApi(service)

	inventoryV1.RegisterInventoryServiceServer(s, api)

	reflection.Register(s)

	go func() {
		log.Printf("Starting Inventory gRPC server on port %d\n", grpcPort)
		err = s.Serve(lis)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("🛑 Gracefully shutdown order server %d\n", grpcPort)
	s.GracefulStop()
	log.Printf("✅ Inventory server stopped")
}
