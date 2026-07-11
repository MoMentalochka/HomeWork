package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	inventoryApi "github.com/MoMentalochka/HomeWork/inventory/internal/api/inventory/v1"
	inventoryRepository "github.com/MoMentalochka/HomeWork/inventory/internal/repository/inventory"
	inventoryService "github.com/MoMentalochka/HomeWork/inventory/internal/service/inventory"
	inventoryV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	defer func() {
		err := lis.Close()
		if err != nil {
			log.Printf("Error closing listener: %s\n", err)
		}
	}()

	s := grpc.NewServer()

	repo := inventoryRepository.NewRepository()
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
