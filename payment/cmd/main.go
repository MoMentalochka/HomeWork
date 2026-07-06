package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	paymentApi "github.com/MoMentalochka/HomeWork/payment/internal/api/payment/v1"
	paymentService "github.com/MoMentalochka/HomeWork/payment/internal/service/payment"
	paymentV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/payment/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50052

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	service := paymentService.NewService()
	api := paymentApi.NewApi(service)
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	paymentV1.RegisterPaymentServiceServer(grpcServer, api)

	go func() {
		log.Printf("Payment server run on %d", grpcPort)
		err := grpcServer.Serve(lis)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Shutting down payment server...")
	grpcServer.GracefulStop()
	log.Println("✅ Payment server gracefully stopped")
}
