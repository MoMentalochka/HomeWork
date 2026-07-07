package repository

import ordersv1 "github.com/MoMentalochka/HomeWork/shared/pkg/openapi/order/v1"

type OrderRepository interface {
	AddOrder(order *ordersv1.OrderDto)
	GetOrder(id string) *ordersv1.OrderDto
}
