package inventory

import (
	"context"

	repomodel "github.com/MoMentalochka/HomeWork/inventory/internal/repository/model"
	inventoryV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
)

func (r *repository) ListParts(_ context.Context, req *inventoryV1.ListPartsRequest) ([]*repomodel.Part, error) {
	return []*repomodel.Part{}, nil
}
