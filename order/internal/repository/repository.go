package repository

import (
	"github.com/MoMentalochka/HomeWork/order/internal/model"
)

type OrderRepository interface {
	AddOrder(order *model.OrderDto) error
	GetOrder(id string) (*model.OrderDto, error)
}
