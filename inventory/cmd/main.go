package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	inventoryV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

const grpcPort = 50051

type inventoryService struct {
	inventoryV1.UnimplementedInventoryServiceServer

	mu    sync.RWMutex
	parts map[string]*inventoryV1.Part
}

func (s *inventoryService) GetPart(context.Context, *inventoryV1.GetPartRequest) (*inventoryV1.GetPartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPart in progress")
}
func (s *inventoryService) ListParts(context.Context, *inventoryV1.ListPartsRequest) (*inventoryV1.ListPartsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListParts in progress")
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	defer func() {
		err := lis.Close()
		if err != nil {
			fmt.Printf("Error closing listener: %s\n", err)
		}
	}()
	serv := grpc.NewServer()

	service := &inventoryService{
		parts: make(map[string]*inventoryV1.Part),
	}

	inventoryV1.RegisterInventoryServiceServer(serv, service)

	reflection.Register(serv)

	go func() {
		log.Printf("Starting Inventory gRPC server on port %d\n", grpcPort)
		err = serv.Serve(lis)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("🛑 Gracefully shutdown inventory server %d\n", grpcPort)
	serv.GracefulStop()
	log.Printf("✅ Inventory server stopped")
}
