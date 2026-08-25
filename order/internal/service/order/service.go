package order

import (
	"github.com/MoMentalochka/HomeWork/order/internal/client/grpc"
	"github.com/MoMentalochka/HomeWork/order/internal/repository"
	"github.com/MoMentalochka/HomeWork/order/internal/service"
)

var _ service.OrderService = (*orderService)(nil)

type orderService struct {
	store           repository.OrderRepository
	paymentClient   grpc.PaymentClient
	inventoryClient grpc.InventoryClient
	producerService service.OrderPaidProducerService
}

func NewOrderService(rep repository.OrderRepository, paymentClient grpc.PaymentClient, inventoryClient grpc.InventoryClient, producerService service.OrderPaidProducerService) *orderService {
	return &orderService{
		rep,
		paymentClient,
		inventoryClient,
		producerService,
	}
}
