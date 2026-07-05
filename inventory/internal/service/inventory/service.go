package inventory

import (
	"github.com/MoMentalochka/HomeWork/inventory/internal/repository"
)

type service struct {
	inventoryRepository repository.InventoryRepository
}
