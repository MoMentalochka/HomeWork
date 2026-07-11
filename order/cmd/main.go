package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	orderApi "github.com/MoMentalochka/HomeWork/order/internal/api/order/v1"
	inventoryClient "github.com/MoMentalochka/HomeWork/order/internal/client/grpc/inventory/v1"
	paymentClient "github.com/MoMentalochka/HomeWork/order/internal/client/grpc/payment/v1"
	orderRepository "github.com/MoMentalochka/HomeWork/order/internal/repository/order"
	orderService "github.com/MoMentalochka/HomeWork/order/internal/service/order"
	ordersv1 "github.com/MoMentalochka/HomeWork/shared/pkg/openapi/order/v1"
	inventoryv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
	paymentv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/payment/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	httpPort          = "8080"
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second
	paymentPort       = "50052"
	inventoryPort     = "50051"
)

func main() {
	//	Payment Client
	paymentConn, err := grpc.NewClient(
		fmt.Sprintf("localhost:%s", paymentPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("failed to connect: %v\n", err)
	}
	paymentC := paymentClient.NewPaymentClient(paymentv1.NewPaymentServiceClient(paymentConn))

	//	Inventory Client
	inventoryConn, err := grpc.NewClient(
		fmt.Sprintf("localhost:%s", inventoryPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("failed to connect: %v\n", err)
	}
	inventoryC := inventoryClient.NewInventoryClient(inventoryv1.NewInventoryServiceClient(inventoryConn))

	defer func() {
		if err := inventoryConn.Close(); err != nil {
			log.Printf("failed to close inventoryConn: %v", err)
		}
		if err := paymentConn.Close(); err != nil {
			log.Printf("failed to close paymentConn: %v", err)
		}
	}()

	repo := orderRepository.NewOrderRepository()
	service := orderService.NewOrderService(repo, paymentC, inventoryC)
	api := orderApi.NewOrderApi(service)
	ordersServer, err := ordersv1.NewServer(api)
	if err != nil {
		fmt.Printf("ошибка создания сервера OpenAPI: %v", err)
	}

	r := chi.NewRouter()

	// Добавляем middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	// Монтируем обработчики OpenAPI
	r.Mount("/", ordersServer)

	server := &http.Server{
		Addr:              net.JoinHostPort("localhost", httpPort),
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout,
	}
	// Создаем контекст с таймаутом для остановки сервера
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	// Запускаем сервер в отдельной горутине
	go func() {
		log.Printf("🚀 HTTP-сервер запущен на порту %s\n", httpPort)
		err = server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("❌ Ошибка запуска сервера: %v\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Завершение работы сервера...")

	err = server.Shutdown(ctx)
	if err != nil {
		log.Printf("❌ Ошибка при остановке сервера: %v\n", err)
	}

	log.Println("✅ Сервер остановлен")
}
