package inventory

import (
	"context"

	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	"github.com/MoMentalochka/HomeWork/inventory/internal/repository/converter"
	repomodel "github.com/MoMentalochka/HomeWork/inventory/internal/repository/model"
)

func (r *repository) ListParts(_ context.Context, filters *repomodel.PartsFilter) ([]*model.Part, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	parts := make([]*model.Part, 0, len(r.data))

	for _, v := range r.data {
		parts = append(parts, new(converter.PartToModel(v)))
	}
	return parts, nil
}
