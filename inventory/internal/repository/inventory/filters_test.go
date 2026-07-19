package inventory

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"

	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	repomodel "github.com/MoMentalochka/HomeWork/inventory/internal/repository/model"
)

var (
	uuid  = gofakeit.UUID()
	parts = []*model.Part{
		{
			Uuid:         uuid,
			Tags:         []string{"ag"},
			Manufacturer: &model.Manufacturer{Country: "France"},
			Category:     "Колёса",
		},
		{
			Uuid:         "2",
			Tags:         []string{"gar"},
			Manufacturer: &model.Manufacturer{Country: "France"},
		},
		{
			Uuid:         "3",
			Name:         "Пушка 1",
			Category:     "Пушки",
			Tags:         []string{"gar"},
			Manufacturer: &model.Manufacturer{Country: "Russia"},
		},
	}
	repoFilters = &repomodel.PartsFilter{Uuids: []string{"3"}, Tags: []string{"gar"}, ManufacturerCountries: []string{"Russia"}}
)

func TestBsonFilterFromPartsFilter(t *testing.T) {
	filters := bsonFilterFromPartsFilter(repoFilters)
	require.Len(t, filters, 3)
}

func TestEmptyFilters(t *testing.T) {
	require.True(t, isEmptyFilter(&repomodel.PartsFilter{}))
	require.Len(t, repomodel.PartsFilter{}.Tags, 0)
}

func TestFilteredParts(t *testing.T) {
	FilteredParts := filteredParts(parts, repoFilters)

	require.Len(t, FilteredParts, 1)
}

func TestFilterByTags(t *testing.T) {
	FilteredParts := filterByTags(parts, []string{"ag"})

	require.Len(t, FilteredParts, 1)

	FilteredParts = filterByTags(parts, []string{"ags"})

	require.Len(t, FilteredParts, 0)
}

func TestFilterByManufacturerCountries(t *testing.T) {
	FilteredParts := filterByManufacturerCountries(parts, []string{"France"})

	require.Len(t, FilteredParts, 2)

	FilteredParts = filterByManufacturerCountries(parts, []string{"Russia"})
	require.Len(t, FilteredParts, 1)
}

func TestFilterByUUID(t *testing.T) {
	FilteredParts := filterByUUID(parts, []string{uuid})

	require.Len(t, FilteredParts, 1)

	FilteredParts = filterByUUID(parts, []string{uuid, "2"})

	require.Len(t, FilteredParts, 2)
}

func TestFilterByName(t *testing.T) {
	FilteredParts := filterByName(parts, []string{"fff"})

	require.Len(t, FilteredParts, 0)

	FilteredParts = filterByName(parts, []string{"Пушка 1"})

	require.True(t, FilteredParts[0].Name == "Пушка 1")
}

func TestFilterByCategory(t *testing.T) {
	FilteredParts := filterByCategory(parts, []string{"Пушки"})

	require.Len(t, FilteredParts, 1)

	FilteredParts = filterByCategory(parts, []string{"Колёса", "Пушки"})

	require.Len(t, FilteredParts, 2)
}
