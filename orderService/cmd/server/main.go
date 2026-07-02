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
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

const (
	httpPort          = "8080"
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second
)

type OrderStore struct {
	mu     sync.Mutex
	orders map[string]*ordersv1.Order
}

func NewOrderStore() *OrderStore {
	return &OrderStore{
		orders: make(map[string]*ordersv1.Order),
	}
}

func (s *OrderStore) AddOrder(order *ordersv1.Order) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.orders[order.OrderUUID] = order
}

func (s *OrderStore) GetOrder(id string) *ordersv1.Order {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, ok := s.orders[id]
	if !ok {
		return nil
	}

	return order
}

type OrderHandler struct {
	store *OrderStore
}

func NewOrderHandler(store *OrderStore) *OrderHandler {
	return &OrderHandler{
		store: store,
	}
}

func (h *OrderHandler) CreateNewOrder(_ context.Context, req *ordersv1.CreateOrderRequest) (*ordersv1.CreateOrderResponse, error) {
	order := &ordersv1.Order{
		OrderUUID: uuid.New().String(),
		PartUuids: req.PartUuids,
		UserUUID:  req.UserUUID,
		Status:    ordersv1.OrderStatusPENDINGPAYMENT,
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

func (h *OrderHandler) OrderPay(_ context.Context, req ordersv1.OptOrderPayRequest, params ordersv1.OrderPayParams) (ordersv1.OrderPayRes, error) {
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

	order.Status = "PAID"
	order.PaymentMethod = req.Value.PaymentMethod
	order.TransactionUUID = uuid.New().String()
	return &ordersv1.OrderPayResponse{TransactionUUID: order.TransactionUUID}, nil
}

func main() {

	storage := NewOrderStore()
	orderHandler := NewOrderHandler(storage)

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

	// Создаем контекст с таймаутом для остановки сервера
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Printf("❌ Ошибка при остановке сервера: %v\n", err)
	}

	log.Println("✅ Сервер остановлен")

}
