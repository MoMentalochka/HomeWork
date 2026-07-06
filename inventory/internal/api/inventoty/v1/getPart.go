package v1

import (
	"context"
	"errors"

	"github.com/MoMentalochka/HomeWork/inventory/internal/converter"
	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	inventoryV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *api) GetPart(_ context.Context, req *inventoryV1.GetPartRequest) (*inventoryV1.GetPartResponse, error) {
	part, err := a.inventoryService.GetPart(context.Background(), req.GetUuid())

	if err != nil {
		if errors.Is(err, model.ErrPartNotFound) {
			return &inventoryV1.GetPartResponse{}, status.Errorf(codes.NotFound, err.Error())
		}
		return &inventoryV1.GetPartResponse{}, status.Errorf(codes.Internal, err.Error())
	}

	return &inventoryV1.GetPartResponse{
		Part: converter.ModelToPart(part),
	}, nil
}
