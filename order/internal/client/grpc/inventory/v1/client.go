package v1

import (
	"context"

	"github.com/MoMentalochka/HomeWork/order/internal/client/converter"
	"github.com/MoMentalochka/HomeWork/order/internal/model"
	inventoryv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
)

type inventoryClient struct {
	generatedClient inventoryv1.InventoryServiceClient
}

func NewInventoryClient(generatedClient inventoryv1.InventoryServiceClient) *inventoryClient {
	return &inventoryClient{generatedClient: generatedClient}
}

func (c *inventoryClient) ListParts(ctx context.Context, filters model.PartsFilter) ([]model.Part, error) {
	res, err := c.generatedClient.ListParts(ctx, &inventoryv1.ListPartsRequest{Filter: converter.ModelPartsFilterToProtoPartsFilter(filters)})
	if err != nil {
		return nil, err
	}

	return converter.PartsToModelParts(res.Parts), nil
}
