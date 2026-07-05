package inventory

import (
	"sync"

	repomodel "github.com/MoMentalochka/HomeWork/inventory/internal/repository/model"
)

type repository struct {
	mu   sync.RWMutex
	data map[string]repomodel.Part
}

func NewRepository() *repository {
	return &repository{
		data: make(map[string]repomodel.Part),
	}
}
