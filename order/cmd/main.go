package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	orderApi "github.com/MoMentalochka/HomeWork/order/internal/api/order/v1"
	inventoryClient "github.com/MoMentalochka/HomeWork/order/internal/client/grpc/inventory/v1"
	paymentClient "github.com/MoMentalochka/HomeWork/order/internal/client/grpc/payment/v1"
	"github.com/MoMentalochka/HomeWork/order/internal/config"
	"github.com/MoMentalochka/HomeWork/order/internal/migrator"
	orderRepository "github.com/MoMentalochka/HomeWork/order/internal/repository/order"
	orderService "github.com/MoMentalochka/HomeWork/order/internal/service/order"
	ordersv1 "github.com/MoMentalochka/HomeWork/shared/pkg/openapi/order/v1"
	inventoryv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
	paymentv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/payment/v1"
)

const configPath = "../../deploy/compose/order/.env"

func main() {
	// Init DB Connection
	ctx := context.Background()

	if err := config.Load(configPath); err != nil {
		log.Printf("failed to load config: %v", err)
		return
	}
	// Создаем соединение с базой данных
	con, err := pgx.Connect(ctx, config.AppConfig().Postgres.URI())
	if err != nil {
		log.Printf("failed to connect: %v\n", err)
		return
	}
	defer func() {
		err = con.Close(ctx)
		if err != nil {
			log.Printf("failed to close connection: %v", err)
		}
	}()

	// Проверяем, что соединение с базой установлено
	err = con.Ping(ctx)
	if err != nil {
		log.Printf("failed to ping: %v\n", err)
		return
	}
	// Инициализируем мигратор
	migrationRunner := migrator.NewMigrator(stdlib.OpenDB(*con.Config().Copy()), config.AppConfig().Postgres.MigrationsDir())
	// Раскатываем миграции
	err = migrationRunner.Up()
	if err != nil {
		log.Printf("failed to run migrations: %v\n", err)
		return
	}

	//	Payment Client
	paymentConn, err := grpc.NewClient(
		config.AppConfig().Payment.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("failed to connect: %v\n", err)
	}
	paymentConn.Connect()
	paymentC := paymentClient.NewPaymentClient(paymentv1.NewPaymentServiceClient(paymentConn))

	//	Inventory Client
	inventoryConn, err := grpc.NewClient(
		config.AppConfig().Inventory.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("failed to connect: %v\n", err)
	}
	inventoryConn.Connect()
	inventoryC := inventoryClient.NewInventoryClient(inventoryv1.NewInventoryServiceClient(inventoryConn))

	defer func() {
		if err := inventoryConn.Close(); err != nil {
			log.Printf("failed to close inventoryConn: %v", err)
		}
		if err := paymentConn.Close(); err != nil {
			log.Printf("failed to close paymentConn: %v", err)
		}
	}()

	// Init Service
	repo := orderRepository.NewOrderRepository(stdlib.OpenDB(*con.Config().Copy()))
	service := orderService.NewOrderService(repo, paymentC, inventoryC)
	api := orderApi.NewOrderApi(service)
	ordersServer, err := ordersv1.NewServer(api)
	if err != nil {
		log.Printf("ошибка создания сервера OpenAPI: %v", err)
	}

	r := chi.NewRouter()

	// Добавляем middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	// Монтируем обработчики OpenAPI
	r.Mount("/", ordersServer)

	server := &http.Server{
		Addr:              config.AppConfig().Http.Address(),
		Handler:           r,
		ReadHeaderTimeout: config.AppConfig().Http.ReadTimeOut(),
	}
	// Создаем контекст с таймаутом для остановки сервера
	ctx, cancel := context.WithTimeout(context.Background(), config.AppConfig().Http.ReadTimeOut())
	defer cancel()

	// Запускаем сервер в отдельной горутине
	go func() {
		log.Printf("🚀 HTTP-сервер запущен на порту %s\n", config.AppConfig().Http.Port())
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
