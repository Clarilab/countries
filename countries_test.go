package countries_test

import (
	"testing"

	"github.com/Clarilab/countries"
)

func TestAlpha2(t *testing.T) {
	for i := range countries.AllMappings() {
		country := countries.AllMappings()[i]

		t.Run("find by country name", func(t *testing.T) {
			alpha2 := countries.Alpha2(country.Translations[countries.EN].Common)
			if country.Alpha2 != alpha2 {
				t.Errorf("expected: %s, got: %s", country.Alpha2, alpha2)
			}
		})

		t.Run("find by alpha 2", func(t *testing.T) {
			alpha2 := countries.Alpha2(country.Alpha2)
			if country.Alpha2 != alpha2 {
				t.Errorf("expected: %s, got: %s", country.Alpha2, alpha2)
			}
		})

		t.Run("find by alpha 3", func(t *testing.T) {
			alpha2 := countries.Alpha2(country.Alpha3)
			if country.Alpha2 != alpha2 {
				t.Errorf("expected: %s, got: %s", country.Alpha2, alpha2)
			}
		})
	}
}

func Test_Exists(t *testing.T) {
	t.Run("exists by country name", func(t *testing.T) {
		if !countries.Exists("Australia") {
			t.Error("expected: true, got: false")
		}
	})

	t.Run("exists by alpha 2", func(t *testing.T) {
		if !countries.Exists("AU") {
			t.Error("expected: true, got: false")
		}
	})

	t.Run("exists by alpha 3", func(t *testing.T) {
		if !countries.Exists("AUS") {
			t.Error("expected: true, got: false")
		}
	})

	t.Run("does not exist", func(t *testing.T) {
		if countries.Exists("NonexistentCountry") {
			t.Error("expected: false, got: true")
		}
	})
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
				_ = countries.Alpha2(input)
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
				_ = countries.Alpha2(input)
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
				_ = countries.Alpha2(input)
			}
		})
	}
}
