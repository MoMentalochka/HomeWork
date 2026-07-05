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

	repo := repository{
		data: make(map[string]repomodel.Part),
	}

	repo.data["1"] = repomodel.Part{Uuid: "1"}
	repo.data["2"] = repomodel.Part{Uuid: "2"}

	return &repo
}
