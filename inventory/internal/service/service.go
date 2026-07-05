package service

import (
	"context"

	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
)

type InventoryService interface {
	GetParts(_ context.Context, uuid string) (model.Part, error)
}
