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

func (c *inventoryClient) GetPart(ctx context.Context, uuid string) (model.Part, error) {
	res, err := c.generatedClient.GetPart(ctx, &inventoryv1.GetPartRequest{Uuid: uuid})
	if err != nil {
		return model.Part{}, err
	}

	return converter.ProtoPartToModelPart(res.Part), nil
}
