package converter

import (
	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	inventoryv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
)

func ModelToPart(part model.Part) *inventoryv1.Part {
	return &inventoryv1.Part{}
}
