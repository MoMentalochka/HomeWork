package inventory

import (
	"sync"

	inventoryV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
)

type repository struct {
	mu   sync.RWMutex
	data map[string]*inventoryV1.Part
}

func NewRepository() *repository {
	return &repository{
		data: make(map[string]*inventoryV1.Part),
	}
}
