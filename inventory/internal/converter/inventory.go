package converter

import (
	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	inventoryv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ModelToPart(part model.Part) *inventoryv1.Part {

	category, ok := inventoryv1.Category_value[string(part.Category)]
	if !ok {
		category = int32(inventoryv1.Category_CATEGORY_UNKNOWN_UNSPECIFIED)
	}

	var dimensions *inventoryv1.Dimensions
	if part.Dimensions != nil {
		dimensions = &inventoryv1.Dimensions{
			Length: part.Dimensions.Length,
			Width:  part.Dimensions.Width,
			Height: part.Dimensions.Height,
			Weight: part.Dimensions.Weight,
		}
	}

	var manufacturer *inventoryv1.Manufacturer
	if part.Manufacturer != nil {
		manufacturer = &inventoryv1.Manufacturer{
			Name:    part.Manufacturer.Name,
			Country: part.Manufacturer.Country,
			Website: part.Manufacturer.Website,
		}
	}

	var updatedAt *timestamppb.Timestamp
	if part.UpdatedAt != nil {
		updatedAt = timestamppb.New(*part.UpdatedAt)
	}
	var createdAt *timestamppb.Timestamp
	if part.UpdatedAt != nil {
		createdAt = timestamppb.New(*part.CreatedAt)
	}

	return &inventoryv1.Part{
		Uuid:          part.Uuid,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      inventoryv1.Category(category),
		Dimensions:    dimensions,
		Manufacturer:  manufacturer,
		Tags:          part.Tags,
		Metadata:      part.Metadata,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}

func ModelsToParts(models []*model.Part) []*inventoryv1.Part {
	parts := make([]*inventoryv1.Part, 0, len(models))
	for _, mod := range models {
		parts = append(parts, ModelToPart(*mod))
	}
	return parts
}

func PartsFilterToModel(filters *inventoryv1.PartsFilter) *model.PartsFilter {
	if filters == nil {
		return &model.PartsFilter{}
	}
	var categories []string

	if filters.Categories != nil && len(filters.Categories) > 0 {
		for _, category := range filters.Categories {
			categories = append(categories, string(category))
		}
	}

	return &model.PartsFilter{
		Uuids:                 filters.Uuids,
		Names:                 filters.Names,
		Categories:            categories,
		ManufacturerCountries: filters.ManufacturerCountries,
		Tags:                  filters.Tags,
	}
}
