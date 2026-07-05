package converter

import (
	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	repomodel "github.com/MoMentalochka/HomeWork/inventory/internal/repository/model"
)

func PartToModel(part repomodel.Part) model.Part {
	return model.Part{
		Uuid:          part.Uuid,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      PartCategoryToModel(part.Category),
		Dimensions:    PartDimensionsToModel(part.Dimensions),
		Manufacturer:  PartManufacturerToModel(part.Manufacturer),
		Tags:          part.Tags,
		Metadata:      part.Metadata,
		CreatedAt:     part.CreatedAt,
		UpdatedAt:     part.UpdatedAt,
	}
}

func PartCategoryToRepoModel(part model.Category) repomodel.Category {
	return repomodel.Category(part)
}

func PartDimensionsToRepoModel(part *model.Dimensions) *repomodel.Dimensions {
	return &repomodel.Dimensions{
		Width:  part.Width,
		Height: part.Height,
		Weight: part.Weight,
		Length: part.Length,
	}
}

func PartManufacturerToRepoModel(part *model.Manufacturer) *repomodel.Manufacturer {
	return &repomodel.Manufacturer{
		Name:    part.Name,
		Country: part.Country,
		Website: part.Website,
	}
}

func PartCategoryToModel(part repomodel.Category) model.Category {
	return model.Category(part)
}

func PartDimensionsToModel(part *repomodel.Dimensions) *model.Dimensions {
	return &model.Dimensions{
		Width:  part.Width,
		Height: part.Height,
		Weight: part.Weight,
		Length: part.Length,
	}
}

func PartManufacturerToModel(part *repomodel.Manufacturer) *model.Manufacturer {
	return &model.Manufacturer{
		Name:    part.Name,
		Country: part.Country,
		Website: part.Website,
	}
}
