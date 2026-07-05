package inventory

import (
	"github.com/MoMentalochka/HomeWork/inventory/internal/repository"
)

type service struct {
	inventoryRepository repository.InventoryRepository
}

func NewService(inventoryRepository repository.InventoryRepository) *service {
	return &service{
		inventoryRepository: inventoryRepository,
	}
}
