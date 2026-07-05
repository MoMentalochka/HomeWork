package v1

import (
	"context"

	"github.com/MoMentalochka/HomeWork/inventory/internal/converter"
	inventoryV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
)

func (a *api) ListParts(ctx context.Context, req *inventoryV1.ListPartsRequest) (*inventoryV1.ListPartsResponse, error) {
	parts, err := a.inventoryService.ListParts(ctx, converter.PartsFilterToModel(req.Filter))
	if err != nil {
		return &inventoryV1.ListPartsResponse{}, err
	}
	return &inventoryV1.ListPartsResponse{Parts: converter.ModelsToParts(parts)}, nil
}
