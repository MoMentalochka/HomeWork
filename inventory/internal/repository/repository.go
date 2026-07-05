package repository

import (
	"context"

	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
)

type InventoryRepository interface {
	GetParts(_ context.Context, uuid string) (model.Part, error)
}
