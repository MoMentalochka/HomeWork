package inventory

import (
	"context"

	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	"github.com/MoMentalochka/HomeWork/inventory/internal/repository/converter"
)

func (s *service) ListParts(ctx context.Context, filters *model.PartsFilter) ([]*model.Part, error) {
	parts, err := s.inventoryRepository.ListParts(ctx, converter.ModelFiltersToRepoModelFilters(filters))
	if err != nil {
	}
	return parts, nil
}
