package inventory

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/MoMentalochka/HomeWork/inventory/internal/repository/mocks"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	inventoryRepository *mocks.InventoryRepository

	inventoryService *service
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.inventoryRepository = &mocks.InventoryRepository{}

	s.inventoryService = NewService(
		s.inventoryRepository,
	)
}

func (s *ServiceSuite) TearDownTest() {
}

func TestAPIIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
