package repository

import (
	"github.com/MoMentalochka/HomeWork/order/internal/model"
)

type OrderRepository interface {
	AddOrder(order *model.OrderDto)
	GetOrder(id string) *model.OrderDto
}
