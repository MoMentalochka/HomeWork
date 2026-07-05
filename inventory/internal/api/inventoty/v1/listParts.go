package v1

import (
	"context"

	inventoryV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
)

func (a *api) ListParts(_ context.Context, req *inventoryV1.ListPartsRequest) (*inventoryV1.ListPartsResponse, error) {
	return &inventoryV1.ListPartsResponse{}, nil
}
