package service

import (
	"context"

	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
)

type InventoryService interface {
	GetPart(_ context.Context, uuid string) (model.Part, error)
	ListParts(_ context.Context, filters *model.PartsFilter) ([]*model.Part, error)
}
