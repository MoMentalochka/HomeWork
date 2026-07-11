package converter

import (
	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	repomodel "github.com/MoMentalochka/HomeWork/inventory/internal/repository/model"
)

func PartToModel(part repomodel.Part) model.Part {
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
		Category:      part.Category,
		Dimensions:    dimensions,
		Manufacturer:  manufacturer,
		Tags:          part.Tags,
		Metadata:      part.Metadata,
		CreatedAt:     part.CreatedAt,
		UpdatedAt:     part.UpdatedAt,
	}
}

func PartDimensionsToModel(dimensions *repomodel.Dimensions) *model.Dimensions {
	return &model.Dimensions{
		Width:  dimensions.Width,
		Height: dimensions.Height,
		Weight: dimensions.Weight,
		Length: dimensions.Length,
	}
}

func PartManufacturerToModel(manufacturer *repomodel.Manufacturer) *model.Manufacturer {
	return &model.Manufacturer{
		Name:    manufacturer.Name,
		Country: manufacturer.Country,
		Website: manufacturer.Website,
	}
}

func ModelFiltersToRepoModelFilters(filters *model.PartsFilter) *repomodel.PartsFilter {
	if filters == nil {
		return &repomodel.PartsFilter{}
	}
	var categories []string

	if filters.Categories != nil {
		categories = append(categories, filters.Categories...)
	}

	return &repomodel.PartsFilter{
		Uuids:                 filters.Uuids,
		Names:                 filters.Names,
		Categories:            categories,
		ManufacturerCountries: filters.ManufacturerCountries,
		Tags:                  filters.Tags,
	}
}
