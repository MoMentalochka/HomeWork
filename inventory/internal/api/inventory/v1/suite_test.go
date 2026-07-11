package v1

import (
	"context"
	"testing"

	"github.com/MoMentalochka/HomeWork/inventory/internal/service/mocks"
	"github.com/stretchr/testify/suite"
)

type APISuite struct {
	suite.Suite

	ctx context.Context

	inventoryService *mocks.InventoryService

	inventoryApi *api
}

func (s *APISuite) SetupTest() {
	s.ctx = context.Background()

	s.inventoryService = mocks.NewInventoryService(s.T())

	s.inventoryApi = NewApi(
		s.inventoryService,
	)
}

func (s *APISuite) TearDownTest() {
}

func TestAPIIntegration(t *testing.T) {
	suite.Run(t, new(APISuite))
}
