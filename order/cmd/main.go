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
	"sync"
	"syscall"
	"time"

	ordersv1 "github.com/MoMentalochka/HomeWork/shared/pkg/openapi/order/v1"
	inventoryv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
	paymentv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/payment/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	httpPort          = "8080"
	paymentPort       = "50052"
	inventoryPort     = "50051"
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second
)

type OrderStore struct {
	mu     sync.Mutex
	orders map[string]*ordersv1.OrderDto
}

func NewOrderStore() *OrderStore {
	return &OrderStore{
		orders: make(map[string]*ordersv1.OrderDto),
	}
}

func (s *OrderStore) AddOrder(order *ordersv1.OrderDto) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.orders[order.OrderUUID] = order
}

func (s *OrderStore) GetOrder(id string) *ordersv1.OrderDto {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, ok := s.orders[id]
	if !ok {
		return nil
	}

	return order
}

type OrderHandler struct {
	store           *OrderStore
	paymentClient   paymentv1.PaymentServiceClient
	inventoryClient inventoryv1.InventoryServiceClient
}

func NewOrderHandler(store *OrderStore, payment paymentv1.PaymentServiceClient, inventory inventoryv1.InventoryServiceClient) *OrderHandler {
	return &OrderHandler{
		store:           store,
		paymentClient:   payment,
		inventoryClient: inventory,
	}
}

func (h *OrderHandler) CreateNewOrder(ctx context.Context, req *ordersv1.CreateOrderRequest) (*ordersv1.CreateOrderResponse, error) {

	for _, partUuid := range req.PartUuids {
		_, err := h.inventoryClient.GetPart(ctx, &inventoryv1.GetPartRequest{Uuid: partUuid})
		if err != nil {
			return nil, err
		}
	}

	order := &ordersv1.OrderDto{
		OrderUUID: uuid.New().String(),
		PartUuids: req.PartUuids,
		UserUUID:  req.UserUUID,
		Status:    ordersv1.StatusPENDINGPAYMENT,
	}

	h.store.AddOrder(order)
	return &ordersv1.CreateOrderResponse{OrderUUID: order.OrderUUID, TotalPrice: 0.0}, nil
}

func (h *OrderHandler) GetOrderById(_ context.Context, params ordersv1.GetOrderByIdParams) (ordersv1.GetOrderByIdRes, error) {

	order := h.store.GetOrder(params.OrderUUID)

	if order == nil {
		return &ordersv1.NotFound{
			Code:    404,
			Message: fmt.Sprintf("Order with uuid '%s' not found", params.OrderUUID),
		}, nil
	}

	return order, nil
}

func (h *OrderHandler) OrderCancel(_ context.Context, params ordersv1.OrderCancelParams) (ordersv1.OrderCancelRes, error) {
	order := h.store.GetOrder(params.OrderUUID)
	if order == nil {
		return &ordersv1.NotFound{
			Code:    404,
			Message: fmt.Sprintf("Order with uuid '%s' not found", params.OrderUUID),
		}, nil
	}
	if order.Status == "PAID" || order.Status == "CANCELLED" {
		return &ordersv1.Conflict{
			Code:    409,
			Message: fmt.Sprintf("Order with uuid '%s' is already paid or cancelled", params.OrderUUID),
		}, nil
	}
	order.Status = "CANCELLED"
	return &ordersv1.OrderCancelResponse{TransactionUUID: order.TransactionUUID}, nil
}

func (h *OrderHandler) OrderPay(ctx context.Context, req ordersv1.OptOrderPayRequest, params ordersv1.OrderPayParams) (ordersv1.OrderPayRes, error) {
	order := h.store.GetOrder(params.OrderUUID)

	if order == nil {
		return &ordersv1.NotFound{
			Code:    404,
			Message: fmt.Sprintf("Order with uuid '%s' not found", params.OrderUUID),
		}, nil
	}

	if order.Status == "CANCELLED" || order.Status == "PAID" {
		return &ordersv1.Conflict{
			Code:    404,
			Message: fmt.Sprintf("Order with uuid '%s' is already paid or cancelled", params.OrderUUID),
		}, nil
	}

	method := paymentv1.PaymentMethod_value[string(req.Value.PaymentMethod)]
	payRequest := paymentv1.PayOrderRequest{
		UserUuid:      order.UserUUID,
		OrderUuid:     order.OrderUUID,
		PaymentMethod: paymentv1.PaymentMethod(method),
	}
	res, err := h.paymentClient.PayOrder(ctx, &payRequest)
	if err != nil {
		return nil, err
	}
	order.Status = "PAID"
	order.PaymentMethod = req.Value.PaymentMethod
	order.TransactionUUID = res.TransactionUuid
	return &ordersv1.OrderPayResponse{TransactionUUID: order.TransactionUUID}, nil
}

func main() {

	paymentConn, err := grpc.NewClient(
		fmt.Sprintf("localhost:%s", paymentPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("failed to connect: %v\n", err)
		return
	}

	inventoryConn, err := grpc.NewClient(
		fmt.Sprintf("localhost:%s", inventoryPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("failed to connect: %v\n", err)
		return
	}

	defer func() {
		if err := inventoryConn.Close(); err != nil {
			log.Printf("failed to close inventoryConn: %v", err)
		}
		if err := paymentConn.Close(); err != nil {
			log.Printf("failed to close paymentConn: %v", err)
		}
	}()

	storage := NewOrderStore()
	orderHandler := NewOrderHandler(storage, paymentv1.NewPaymentServiceClient(paymentConn), inventoryv1.NewInventoryServiceClient(inventoryConn))

	ordersServer, err := ordersv1.NewServer(orderHandler)
	if err != nil {
		log.Fatalf("ошибка создания сервера OpenAPI: %v", err)
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
