package order

import (
	"sync"

	"github.com/MoMentalochka/HomeWork/order/internal/model"
)

type orderRepository struct {
	mu     sync.Mutex
	orders map[string]*model.OrderDto
}

func NewOrderRepository() *orderRepository {
	return &orderRepository{
		orders: make(map[string]*model.OrderDto),
	}
}

func (r *orderRepository) AddOrder(order *model.OrderDto) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.orders[order.OrderUUID] = order
}

func (r *orderRepository) GetOrder(id string) *model.OrderDto {
	r.mu.Lock()
	defer r.mu.Unlock()

	order, ok := r.orders[id]
	if !ok {
		return nil
	}

	return order
}
