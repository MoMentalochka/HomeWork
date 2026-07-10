package converter

import (
	"github.com/MoMentalochka/HomeWork/order/internal/model"
	inventoryv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
)

func ModelPartsFilterToProtoPartsFilter(filters model.PartsFilter) *inventoryv1.PartsFilter {
	var categories []inventoryv1.Category

	if filters.Categories != nil {
		for _, category := range filters.Categories {
			val, ok := inventoryv1.Category_value[category]
			if !ok {
				categories = append(categories, inventoryv1.Category_CATEGORY_UNKNOWN_UNSPECIFIED)
				continue
			}
			categories = append(categories, inventoryv1.Category(val))
		}
	}

	return &inventoryv1.PartsFilter{
		Uuids:                 filters.Uuids,
		Names:                 filters.Names,
		Categories:            categories,
		ManufacturerCountries: filters.ManufacturerCountries,
		Tags:                  filters.Tags,
	}
}

func ProtoPartToModelPart(part *inventoryv1.Part) model.Part {
	var dimensions *model.Dimensions
	if part.Dimensions != nil {
		dimensions = &model.Dimensions{
			Width:  part.Dimensions.Width,
			Height: part.Dimensions.Height,
			Length: part.Dimensions.Length,
			Weight: part.Dimensions.Weight,
		}
	}

	var manufacturer *model.Manufacturer
	if part.Manufacturer != nil {
		manufacturer = &model.Manufacturer{
			Name:    part.Manufacturer.Name,
			Country: part.Manufacturer.Name,
			Website: part.Manufacturer.Name,
		}
	}

	return model.Part{
		Uuid:          part.Uuid,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      part.Category.String(),
		Dimensions:    dimensions,
		Manufacturer:  manufacturer,
		Tags:          part.Tags,
		Metadata:      part.Metadata,
		CreatedAt:     new(part.CreatedAt.AsTime()),
		UpdatedAt:     new(part.UpdatedAt.AsTime()),
	}
}

func PartsToModelParts(parts []*inventoryv1.Part) []model.Part {
	result := make([]model.Part, len(parts))
	for _, part := range parts {
		result = append(result, ProtoPartToModelPart(part))
	}
	return result
}
