package order

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	clientMocks "github.com/MoMentalochka/HomeWork/order/internal/client/grpc/mocks"
	"github.com/MoMentalochka/HomeWork/order/internal/repository/mocks"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	orderRepository  *mocks.OrderRepository
	paymentService   *clientMocks.PaymentClient
	inventoryService *clientMocks.InventoryClient

	orderService *orderService
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.orderRepository = &mocks.OrderRepository{}

	s.orderService = NewOrderService(
		s.orderRepository,
		s.paymentService,
		s.inventoryService,
	)
}

func (s *ServiceSuite) TearDownTest() {
}

func TestAPIIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
