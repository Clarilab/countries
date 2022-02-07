package countries_test

import (
	"sort"
	"testing"

	"github.com/Clarilab/countries"
	"github.com/stretchr/testify/assert"
)

func TestGetAlpha2(t *testing.T) {
	for i := range countries.GetAllMappings() {
		country := countries.GetAllMappings()[i]

		t.Run("find by country name", func(t *testing.T) {
			alpha2 := countries.GetAlpha2(country.Country)
			assert.Equal(t, country.Alpha2, alpha2)
		})

		t.Run("find by alpha 2", func(t *testing.T) {
			alpha2 := countries.GetAlpha2(country.Alpha2)
			assert.Equal(t, country.Alpha2, alpha2)
		})

		t.Run("find by alpha 3", func(t *testing.T) {
			alpha2 := countries.GetAlpha2(country.Alpha3)
			assert.Equal(t, country.Alpha2, alpha2)
		})
	}
}

func BenchmarkGetByCountryName(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = countries.GetCountryName("Germany")
	}
}

func TestMappingsAreSortedByCountryName(t *testing.T) {
	mappings := countries.GetAllMappings()

	sorted := make([]countries.Mapping, len(mappings))
	copy(sorted, mappings)

	// sort the "sorted" slice
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Country <= sorted[j].Country
	})

	// if both slices are equal after sorting "sorted", mappings was already sorted
	assert.Equal(t, sorted, mappings)
}
