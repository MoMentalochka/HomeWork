package inventory

import (
	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	repomodel "github.com/MoMentalochka/HomeWork/inventory/internal/repository/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func filteredParts(parts []*model.Part, filters *repomodel.PartsFilter) []*model.Part {
	if isEmptyFilter(filters) {
		return parts
	}
	result := make([]*model.Part, len(parts))

	copy(result, parts)

	if len(filters.Uuids) > 0 {
		result = filterByUUID(result, filters.Uuids)
	}

	if len(filters.Names) > 0 {
		result = filterByName(result, filters.Names)
	}

	if len(filters.Categories) > 0 {
		result = filterByCategory(result, filters.Categories)
	}

	if len(filters.ManufacturerCountries) > 0 {
		result = filterByManufacturerCountries(result, filters.ManufacturerCountries)
	}

	if len(filters.Tags) > 0 {
		result = filterByTags(result, filters.Tags)
	}

	return result
}

func filterByName(parts []*model.Part, filters []string) []*model.Part {
	var result []*model.Part

	allowedSet := make(map[string]struct{})
	for _, filter := range filters {
		allowedSet[filter] = struct{}{}
	}

	for _, part := range parts {
		if _, ok := allowedSet[part.Name]; ok {
			result = append(result, part)
		}
	}
	return result
}

func filterByUUID(parts []*model.Part, filters []string) []*model.Part {
	var result []*model.Part

	allowedSet := makeSet(filters)

	for _, part := range parts {
		if _, ok := allowedSet[part.Uuid]; ok {
			result = append(result, part)
		}
	}
	return result
}

func filterByCategory(parts []*model.Part, filters []string) []*model.Part {
	var result []*model.Part

	allowedSet := makeSet(filters)

	for _, part := range parts {
		if _, ok := allowedSet[part.Category]; ok {
			result = append(result, part)
		}
	}
	return result
}

func filterByManufacturerCountries(parts []*model.Part, filters []string) []*model.Part {
	var result []*model.Part

	allowedSet := makeSet(filters)

	for _, part := range parts {
		if _, ok := allowedSet[part.Manufacturer.Country]; ok {
			result = append(result, part)
		}
	}
	return result
}

func filterByTags(parts []*model.Part, filters []string) []*model.Part {
	var result []*model.Part

	allowedSet := makeSet(filters)

	for _, part := range parts {
		if hasAny(part.Tags, allowedSet) {
			result = append(result, part)
		}
	}
	return result
}

func isEmptyFilter(filter *repomodel.PartsFilter) bool {
	return len(filter.Uuids) == 0 &&
		len(filter.Names) == 0 &&
		len(filter.Categories) == 0 &&
		len(filter.ManufacturerCountries) == 0 &&
		len(filter.Tags) == 0
}

func bsonFilterFromPartsFilter(f *repomodel.PartsFilter) bson.M {
	filter := bson.M{}

	if len(f.Uuids) > 0 {
		filter["uuid"] = bson.M{"$in": f.Uuids}
	}
	if len(f.Names) > 0 {
		filter["name"] = bson.M{"$in": f.Names}
	}
	if len(f.ManufacturerCountries) > 0 {
		filter["manufacturer"] = bson.M{"$in": f.ManufacturerCountries}
	}
	if len(f.Tags) > 0 {
		filter["tags"] = bson.M{"$in": f.Tags}
	}
	if len(f.Categories) > 0 {
		filter["category"] = bson.M{"$in": f.Categories}
	}

	return filter
}
