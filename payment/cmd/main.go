package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	paymentV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/payment/v1"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50052

type PaymentService struct {
	paymentV1.UnimplementedPaymentServiceServer
}

func (s *PaymentService) PayOrder(_ context.Context, req *paymentV1.PayOrderRequest) (*paymentV1.PayOrderResponse, error) {
	id := uuid.New().String()
	log.Printf("Оплата прошла успешно, transaction_uuid: %s", id)
	return &paymentV1.PayOrderResponse{TransactionUuid: id}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	serv := &PaymentService{}
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	paymentV1.RegisterPaymentServiceServer(grpcServer, serv)

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
