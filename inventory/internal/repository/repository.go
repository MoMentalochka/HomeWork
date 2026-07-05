package repository

import (
	"context"

	repomodel "github.com/MoMentalochka/HomeWork/inventory/internal/repository/model"
)

type InventoryRepository interface {
	GetPart(_ context.Context, uuid string) (repomodel.Part, error)
	ListParts(_ context.Context, uuid string) ([]*repomodel.Part, error)
}
