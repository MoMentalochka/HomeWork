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

	var category model.Category
	if part.Category != "" {
		category = PartCategoryToModel(part.Category)
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
		Category:      category,
		Dimensions:    dimensions,
		Manufacturer:  manufacturer,
		Tags:          part.Tags,
		Metadata:      part.Metadata,
		CreatedAt:     part.CreatedAt,
		UpdatedAt:     part.UpdatedAt,
	}
}

//func PartCategoryToRepoModel(category model.Category) repomodel.Category {
//	return repomodel.Category(category)
//}
//
//func PartDimensionsToRepoModel(dimensions *model.Dimensions) *repomodel.Dimensions {
//	return &repomodel.Dimensions{
//		Width:  dimensions.Width,
//		Height: dimensions.Height,
//		Weight: dimensions.Weight,
//		Length: dimensions.Length,
//	}
//}
//
//func PartManufacturerToRepoModel(manufacturer *model.Manufacturer) *repomodel.Manufacturer {
//	return &repomodel.Manufacturer{
//		Name:    manufacturer.Name,
//		Country: manufacturer.Country,
//		Website: manufacturer.Website,
//	}
//}

func PartCategoryToModel(part repomodel.Category) model.Category {
	return model.Category(part)
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
	var categories []repomodel.Category

	if filters.Categories != nil && len(filters.Categories) > 0 {
		for _, category := range filters.Categories {
			categories = append(categories, repomodel.Category(category))
		}
	}

	return &repomodel.PartsFilter{
		Uuids:                 filters.Uuids,
		Names:                 filters.Names,
		Categories:            categories,
		ManufacturerCountries: filters.ManufacturerCountries,
		Tags:                  filters.Tags,
	}
}
