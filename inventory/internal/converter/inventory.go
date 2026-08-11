package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	inventoryV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
)

func ModelToPart(part model.Part) *inventoryV1.Part {
	var dimensions *inventoryV1.Dimensions
	if part.Dimensions != nil {
		dimensions = &inventoryV1.Dimensions{
			Length: part.Dimensions.Length,
			Width:  part.Dimensions.Width,
			Height: part.Dimensions.Height,
			Weight: part.Dimensions.Weight,
		}
	}

	var manufacturer *inventoryV1.Manufacturer
	if part.Manufacturer != nil {
		manufacturer = &inventoryV1.Manufacturer{
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

	return &inventoryV1.Part{
		Uuid:          part.Uuid,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      convertCategoryToGRPC(part.Category),
		Dimensions:    dimensions,
		Manufacturer:  manufacturer,
		Tags:          part.Tags,
		Metadata:      part.Metadata,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}

func ModelsToParts(models []*model.Part) []*inventoryV1.Part {
	parts := make([]*inventoryV1.Part, 0, len(models))
	for _, mod := range models {
		parts = append(parts, ModelToPart(*mod))
	}
	return parts
}

func PartsFilterToModel(filters *inventoryV1.PartsFilter) *model.PartsFilter {
	if filters == nil {
		return &model.PartsFilter{}
	}
	var categories []model.Category

	if filters.Categories != nil {
		for _, category := range filters.Categories {
			categories = append(categories, convertGRPCCategoryToModelCategory(category))
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

func convertCategoryToGRPC(category model.Category) inventoryV1.Category {
	switch category {
	case model.CategoryEngine:
		return inventoryV1.Category_CATEGORY_ENGINE
	case model.CategoryFuel:
		return inventoryV1.Category_CATEGORY_FUEL
	case model.CategoryPorthole:
		return inventoryV1.Category_CATEGORY_PORTHOLE
	case model.CategoryWing:
		return inventoryV1.Category_CATEGORY_WING
	default:
		return inventoryV1.Category_CATEGORY_UNKNOWN_UNSPECIFIED
	}
}

func convertGRPCCategoryToModelCategory(grpcCategory inventoryV1.Category) model.Category {
	switch grpcCategory {
	case inventoryV1.Category_CATEGORY_ENGINE:
		return model.CategoryEngine
	case inventoryV1.Category_CATEGORY_FUEL:
		return model.CategoryFuel
	case inventoryV1.Category_CATEGORY_PORTHOLE:
		return model.CategoryPorthole
	case inventoryV1.Category_CATEGORY_WING:
		return model.CategoryWing
	default:
		return model.CategoryUnknown
	}
}
