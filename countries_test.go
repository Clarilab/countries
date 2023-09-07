package countries_test

import (
	"testing"

	"github.com/Clarilab/countries/v2"
)

func TestGetAlpha2(t *testing.T) {
	for i := range countries.GetAllMappings() {
		country := countries.GetAllMappings()[i]

		t.Run("find by country name", func(t *testing.T) {
			alpha2 := countries.GetAlpha2(country.Translations[string(countries.EN)].Common)
			if country.Alpha2 != alpha2 {
				t.Errorf("expected: %s, got: %s", country.Alpha2, alpha2)
			}
		})

		t.Run("find by alpha 2", func(t *testing.T) {
			alpha2 := countries.GetAlpha2(country.Alpha2)
			if country.Alpha2 != alpha2 {
				t.Errorf("expected: %s, got: %s", country.Alpha2, alpha2)
			}
		})

		t.Run("find by alpha 3", func(t *testing.T) {
			alpha2 := countries.GetAlpha2(country.Alpha3)
			if country.Alpha2 != alpha2 {
				t.Errorf("expected: %s, got: %s", country.Alpha2, alpha2)
			}
		})
	}
}

func BenchmarkGetByCountryName(b *testing.B) {
	b.ReportAllocs()

	inputs := []string{
		"Australia", // best case
		"Japan",     // worst case
	}

	for _, input := range inputs {
		b.Run("input_"+input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = countries.GetAlpha2(input)
			}
		})
	}
}

func BenchmarkGetByAlpha2(b *testing.B) {
	b.ReportAllocs()

	inputs := []string{
		"AU", // best case
		"JP", // worst case
	}

	for _, input := range inputs {
		b.Run("input_"+input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = countries.GetAlpha2(input)
			}
		})
	}
}

func BenchmarkGetByAlpha3(b *testing.B) {
	b.ReportAllocs()

	inputs := []string{
		"AUS", // best case
		"JPN", // worst case
	}

	for _, input := range inputs {
		b.Run("input_"+input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = countries.GetAlpha2(input)
			}
		})
	}
}
