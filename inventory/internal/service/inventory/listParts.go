package inventory

import (
	"context"

	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	inventoryV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
)

func (s *service) ListParts(_ context.Context, req *inventoryV1.ListPartsRequest) ([]*model.Part, error) {
	return []*model.Part{}, nil
}
