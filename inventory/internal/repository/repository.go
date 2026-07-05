package repository

import (
	"context"

	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	repomodel "github.com/MoMentalochka/HomeWork/inventory/internal/repository/model"
)

type InventoryRepository interface {
	GetPart(_ context.Context, uuid string) (repomodel.Part, error)
	ListParts(_ context.Context, filters *repomodel.PartsFilter) ([]*model.Part, error)
}
