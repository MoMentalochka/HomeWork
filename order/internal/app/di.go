package app

import (
	"context"
	"database/sql"
	"log"

	orderV1Api "github.com/MoMentalochka/HomeWork/order/internal/api/order/v1"
	orderGRPC "github.com/MoMentalochka/HomeWork/order/internal/client/grpc"
	inventoryClient "github.com/MoMentalochka/HomeWork/order/internal/client/grpc/inventory/v1"
	paymentClient "github.com/MoMentalochka/HomeWork/order/internal/client/grpc/payment/v1"
	"github.com/MoMentalochka/HomeWork/order/internal/config"
	"github.com/MoMentalochka/HomeWork/order/internal/migrator"
	"github.com/MoMentalochka/HomeWork/order/internal/repository"
	orderRepository "github.com/MoMentalochka/HomeWork/order/internal/repository/order"
	"github.com/MoMentalochka/HomeWork/order/internal/service"
	orderService "github.com/MoMentalochka/HomeWork/order/internal/service/order"
	"github.com/MoMentalochka/HomeWork/platform/pkg/closer"
	ordersV1 "github.com/MoMentalochka/HomeWork/shared/pkg/openapi/order/v1"
	inventoryv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
	paymentv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/payment/v1"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type diContainer struct {
	orderRepo    repository.OrderRepository
	orderService service.OrderService
	orderV1Api   ordersV1.Invoker

	inventoryClient orderGRPC.InventoryClient
	paymentClient   orderGRPC.PaymentClient

	postgresDB *sql.DB
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) OrderV1API(ctx context.Context) ordersV1.Invoker {
	if d.orderV1Api == nil {
		d.orderV1Api = orderV1Api.NewOrderApi(d.OrderService(ctx))
	}

	return d.orderV1Api
}

func (d *diContainer) OrderService(ctx context.Context) service.OrderService {
	if d.orderService == nil {
		d.orderService = orderService.NewOrderService(d.OrderRepo(ctx), d.PaymentClient(ctx), d.InventoryClient(ctx))
	}

	return d.orderService
}

func (d *diContainer) OrderRepo(ctx context.Context) repository.OrderRepository {
	if d.orderRepo == nil {
		d.orderRepo = orderRepository.NewOrderRepository(d.PostgresDB(ctx))
	}

	return d.orderRepo
}

func (d *diContainer) PostgresDB(ctx context.Context) *sql.DB {
	if d.postgresDB == nil {
		// Создаем соединение с базой данных
		con, err := pgx.Connect(ctx, config.AppConfig().Postgres.URI())
		if err != nil {
			log.Printf("failed to connect postgres: %v\n", err)
			return nil
		}
		// Проверяем, что соединение с базой установлено
		err = con.Ping(ctx)
		if err != nil {
			log.Printf("failed to ping postgres: %v\n", err)
			return nil
		}

		closer.AddNamed("PostgresDB client", func(ctx context.Context) error {
			return con.Close(ctx)
		})

		d.postgresDB = stdlib.OpenDB(*con.Config().Copy())
		// Инициализируем мигратор
		migrationRunner := migrator.NewMigrator(stdlib.OpenDB(*con.Config().Copy()), config.AppConfig().Postgres.MigrationsDir())
		// Раскатываем миграции
		err = migrationRunner.Up()
		if err != nil {
			log.Printf("failed to run migrations: %v\n", err)
			return nil
		}

	}

	return d.postgresDB
}

func (d *diContainer) PaymentClient(_ context.Context) orderGRPC.PaymentClient {
	if d.paymentClient == nil {
		paymentConn, err := grpc.NewClient(
			config.AppConfig().Payment.Address(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Printf("failed to connect: %v\n", err)
		}
		paymentConn.Connect()
		closer.AddNamed("Order payment client", func(ctx context.Context) error {
			return paymentConn.Close()
		})

		d.paymentClient = paymentClient.NewPaymentClient(paymentv1.NewPaymentServiceClient(paymentConn))
	}

	return d.paymentClient
}

func (d *diContainer) InventoryClient(_ context.Context) orderGRPC.InventoryClient {
	if d.inventoryClient == nil {
		inventoryConn, err := grpc.NewClient(
			config.AppConfig().Inventory.Address(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Printf("failed to connect: %v\n", err)
		}
		inventoryConn.Connect()
		closer.AddNamed("Order inventory client", func(ctx context.Context) error {
			return inventoryConn.Close()
		})
		d.inventoryClient = inventoryClient.NewInventoryClient(inventoryv1.NewInventoryServiceClient(inventoryConn))
	}

	return d.inventoryClient
}
