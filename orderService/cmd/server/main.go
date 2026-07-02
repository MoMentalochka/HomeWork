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
)

const (
	httpPort          = "8080"
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second
)

type OrderStore struct {
	mu     sync.Mutex
	orders map[string]*ordersv1.CreateOrderRequest
}

func NewOrderStore() *OrderStore {
	return &OrderStore{
		orders: make(map[string]*ordersv1.CreateOrderRequest),
	}
}

func (s *OrderStore) AddOrder(order *ordersv1.CreateOrderRequest) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.orders["13223"] = order
	fmt.Println("Записал", s.orders)
}

func (s *OrderStore) GetOrder(id string) *ordersv1.CreateOrderRequest {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, ok := s.orders["13223"]
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

func (h *OrderHandler) CreateNewOrder(ctx context.Context, req *ordersv1.CreateOrderRequest) (*ordersv1.Order, error) {
	return &ordersv1.Order{}, nil
}

func (h *OrderHandler) GetOrderById(ctx context.Context, params ordersv1.GetOrderByIdParams) (ordersv1.GetOrderByIdRes, error) {

	order := h.store.GetOrder(params.OrderUUID)

	if order == nil {
		return &ordersv1.NotFound{
			Code:    404,
			Message: fmt.Sprintf("Order with uuid '%s' not found", params.OrderUUID),
		}, nil
	}

	return nil, nil
}

func (h *OrderHandler) OrderCancel(ctx context.Context, params ordersv1.OrderCancelParams) (ordersv1.OrderCancelRes, error) {
	return nil, nil
}

func (h *OrderHandler) OrderPay(ctx context.Context, req *ordersv1.OrderPayRequest, params ordersv1.OrderPayParams) (ordersv1.OrderPayRes, error) {
	return nil, nil
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
