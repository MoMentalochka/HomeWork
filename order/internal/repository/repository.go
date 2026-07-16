package repository

import (
	"github.com/MoMentalochka/HomeWork/order/internal/model"
)

type OrderRepository interface {
	Create(order *model.OrderDto) error
	Get(id string) (*model.OrderDto, error)
	Update(id string, data model.OrderDto) (*model.OrderDto, error)
}
