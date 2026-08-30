package order

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	clientMocks "github.com/MoMentalochka/HomeWork/order/internal/client/grpc/mocks"
	"github.com/MoMentalochka/HomeWork/order/internal/repository/mocks"
	serviceMocks "github.com/MoMentalochka/HomeWork/order/internal/service/mocks"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	orderRepository  *mocks.OrderRepository
	paymentService   *clientMocks.PaymentClient
	inventoryService *clientMocks.InventoryClient

	orderService  *orderService
	orderProducer *serviceMocks.OrderPaidProducerService
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.orderRepository = &mocks.OrderRepository{}

	s.orderService = NewOrderService(
		s.orderRepository,
		s.paymentService,
		s.inventoryService,
		s.orderProducer,
	)
}

func (s *ServiceSuite) TearDownTest() {
}

func TestAPIIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
