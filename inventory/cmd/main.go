package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	inventoryApi "github.com/MoMentalochka/HomeWork/inventory/internal/api/inventoty/v1"
	inventoryRepository "github.com/MoMentalochka/HomeWork/inventory/internal/repository/inventory"
	inventoryService "github.com/MoMentalochka/HomeWork/inventory/internal/service/inventory"
	inventoryV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

//func filterSlice(slice []*inventoryV1.Part, predicate func(*inventoryV1.Part) bool) []*inventoryV1.Part {
//	result := make([]*inventoryV1.Part, 0, len(slice))
//	for _, part := range slice {
//		if predicate(part) {
//			result = append(result, part)
//		}
//	}
//	return result
//}

//type inventoryService struct {
//	inventoryV1.UnimplementedInventoryServiceServer
//
//	mu    sync.RWMutex
//	parts map[string]*inventoryV1.Part
//}
//
//func (s *inventoryService) GetPart(_ context.Context, req *inventoryV1.GetPartRequest) (*inventoryV1.GetPartResponse, error) {
//	part, ok := s.parts[req.Uuid]
//	if !ok {
//		return nil, status.Error(codes.NotFound, fmt.Sprintf("Part with id: \"%s\" not found", req.Uuid))
//	}
//
//	return &inventoryV1.GetPartResponse{Part: part}, nil
//}
//func (s *inventoryService) ListParts(_ context.Context, req *inventoryV1.ListPartsRequest) (*inventoryV1.ListPartsResponse, error) {
//	filteredParts := make([]*inventoryV1.Part, 0, len(s.parts))
//	if req.Filter == nil {
//		for _, v := range s.parts {
//			filteredParts = append(filteredParts, v)
//		}
//		return &inventoryV1.ListPartsResponse{
//			Parts: filteredParts,
//		}, nil
//	}
//	// If uuid's included fast filter
//	if len(req.Filter.Uuids) == 0 {
//		for _, v := range s.parts {
//			filteredParts = append(filteredParts, v)
//		}
//	} else {
//		seen := make(map[string]struct{})
//		for _, v := range req.Filter.Uuids {
//			part, ok := s.parts[v]
//			_, hadSeen := seen[v]
//			if ok && !hadSeen {
//				seen[v] = struct{}{}
//				filteredParts = append(filteredParts, part)
//			}
//		}
//	}
//
//	if len(filteredParts) == 0 {
//		return nil, status.Error(codes.NotFound, fmt.Sprintf("Parts not found"))
//	}
//	for _, v := range req.Filter.Categories {
//		filteredParts = filterSlice(filteredParts, func(part *inventoryV1.Part) bool {
//			return part.Category == v
//		})
//	}
//	if len(filteredParts) == 0 {
//		return nil, status.Error(codes.NotFound, fmt.Sprintf("Parts not found"))
//	}
//	for _, v := range req.Filter.Names {
//		filteredParts = filterSlice(filteredParts, func(part *inventoryV1.Part) bool {
//			return part.Name == v
//		})
//	}
//	if len(filteredParts) == 0 {
//		return nil, status.Error(codes.NotFound, fmt.Sprintf("Parts not found"))
//	}
//	for _, v := range req.Filter.ManufacturerCountries {
//		filteredParts = filterSlice(filteredParts, func(part *inventoryV1.Part) bool {
//			return part.Manufacturer.Country == v
//		})
//	}
//	if len(filteredParts) == 0 {
//		return nil, status.Error(codes.NotFound, fmt.Sprintf("Parts not found"))
//	}
//	for _, v := range req.Filter.Tags {
//		filteredParts = filterSlice(filteredParts, func(part *inventoryV1.Part) bool {
//			for _, tag := range part.Tags {
//				if tag == v {
//					return true
//				}
//			}
//			return false
//		})
//	}
//	if len(filteredParts) == 0 {
//		return nil, status.Error(codes.NotFound, fmt.Sprintf("Parts not found"))
//	}
//	return &inventoryV1.ListPartsResponse{
//		Parts: filteredParts,
//	}, nil
//
//	//return nil, status.Errorf(codes.Unimplemented, "method ListParts in progress")
//}

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
	log.Printf("🛑 Gracefully shutdown inventory server %d\n", grpcPort)
	s.GracefulStop()
	log.Printf("✅ Inventory server stopped")
}
