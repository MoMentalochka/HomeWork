package v1

import (
	"context"

	"github.com/MoMentalochka/HomeWork/inventory/internal/converter"
	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	inventoryV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
	"github.com/go-faster/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *api) ListParts(ctx context.Context, req *inventoryV1.ListPartsRequest) (*inventoryV1.ListPartsResponse, error) {
	parts, err := a.inventoryService.ListParts(ctx, converter.PartsFilterToModel(req.Filter))
	if err != nil {
		if errors.Is(err, model.ErrPartNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return &inventoryV1.ListPartsResponse{}, err
	}
	return &inventoryV1.ListPartsResponse{Parts: converter.ModelsToParts(parts)}, nil
}
