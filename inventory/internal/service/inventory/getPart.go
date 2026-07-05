package inventory

import (
	"context"

	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	"github.com/MoMentalochka/HomeWork/inventory/internal/repository/converter"
)

func (s *service) GetPart(_ context.Context, uuid string) (model.Part, error) {
	part, err := s.inventoryRepository.GetPart(context.Background(), uuid)

	if err != nil {
		return model.Part{}, err
	}

	return converter.PartToModel(part), nil
}
