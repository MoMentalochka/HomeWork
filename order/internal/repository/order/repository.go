package order

import (
	"sync"

	ordersv1 "github.com/MoMentalochka/HomeWork/shared/pkg/openapi/order/v1"
)

type orderRepository struct {
	mu     sync.Mutex
	orders map[string]*ordersv1.OrderDto
}

func NewOrderRepository() *orderRepository {
	return &orderRepository{
		orders: make(map[string]*ordersv1.OrderDto),
	}
}

func (r *orderRepository) AddOrder(order *ordersv1.OrderDto) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.orders[order.OrderUUID] = order
}

func (r *orderRepository) GetOrder(id string) *ordersv1.OrderDto {
	r.mu.Lock()
	defer r.mu.Unlock()

	order, ok := r.orders[id]
	if !ok {
		return nil
	}

	return order
}
