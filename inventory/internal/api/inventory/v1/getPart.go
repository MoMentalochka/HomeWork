package v1

import (
	"context"

	"github.com/MoMentalochka/HomeWork/inventory/internal/converter"
	inventoryV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
)

func (a *api) GetPart(ctx context.Context, req *inventoryV1.GetPartRequest) (*inventoryV1.GetPartResponse, error) {
	part, err := a.inventoryService.GetPart(ctx, req.GetUuid())
	if err != nil {
		return &inventoryV1.GetPartResponse{}, err
	}

	return &inventoryV1.GetPartResponse{
		Part: converter.ModelToPart(part),
	}, nil
}
