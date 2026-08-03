package converter

import (
	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	repoModel "github.com/MoMentalochka/HomeWork/inventory/internal/repository/model"
)

func PartToModel(part repoModel.Part) model.Part {
	var dimensions *model.Dimensions
	if part.Dimensions != nil {
		dimensions = PartDimensionsToModel(part.Dimensions)
	}

	var manufacturer *model.Manufacturer
	if part.Manufacturer != nil {
		manufacturer = PartManufacturerToModel(part.Manufacturer)
	}

	return model.Part{
		Uuid:          part.Uuid,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      convertRepoCategoryToServiceCategory(part.Category),
		Dimensions:    dimensions,
		Manufacturer:  manufacturer,
		Tags:          part.Tags,
		Metadata:      part.Metadata,
		CreatedAt:     part.CreatedAt,
		UpdatedAt:     part.UpdatedAt,
	}
}

func PartDimensionsToModel(dimensions *repoModel.Dimensions) *model.Dimensions {
	return &model.Dimensions{
		Width:  dimensions.Width,
		Height: dimensions.Height,
		Weight: dimensions.Weight,
		Length: dimensions.Length,
	}
}

func PartManufacturerToModel(manufacturer *repoModel.Manufacturer) *model.Manufacturer {
	return &model.Manufacturer{
		Name:    manufacturer.Name,
		Country: manufacturer.Country,
		Website: manufacturer.Website,
	}
}

func ModelFiltersToRepoModelFilters(filters *model.PartsFilter) *repoModel.PartsFilter {
	if filters == nil {
		return &repoModel.PartsFilter{}
	}
	var categories []repoModel.Category

	if filters.Categories != nil {
		for _, category := range filters.Categories {
			categories = append(categories, convertServiceCategoryToRepoCategory(category))
		}
	}

	return &repoModel.PartsFilter{
		Uuids:                 filters.Uuids,
		Names:                 filters.Names,
		Categories:            categories,
		ManufacturerCountries: filters.ManufacturerCountries,
		Tags:                  filters.Tags,
	}
}

func convertServiceCategoryToRepoCategory(serviceCategory model.Category) repoModel.Category {
	switch serviceCategory {
	case model.CategoryEngine:
		return repoModel.CategoryEngine
	case model.CategoryFuel:
		return repoModel.CategoryFuel
	case model.CategoryPorthole:
		return repoModel.CategoryPorthole
	case model.CategoryWing:
		return repoModel.CategoryWing
	default:
		return repoModel.CategoryUnknown
	}
}

func convertRepoCategoryToServiceCategory(repoCategory repoModel.Category) model.Category {
	switch repoCategory {
	case repoModel.CategoryEngine:
		return model.CategoryEngine
	case repoModel.CategoryFuel:
		return model.CategoryFuel
	case repoModel.CategoryPorthole:
		return model.CategoryPorthole
	case repoModel.CategoryWing:
		return model.CategoryWing
	default:
		return model.CategoryUnknown
	}
}
