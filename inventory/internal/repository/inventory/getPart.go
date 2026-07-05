package inventory

import (
	"context"

	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	repomodel "github.com/MoMentalochka/HomeWork/inventory/internal/repository/model"
)

func (r *repository) GetPart(_ context.Context, uuid string) (repomodel.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, ok := r.data[uuid]
	if !ok {
		return repomodel.Part{}, model.ErrPartNotFound
	}

	return p, nil
}
