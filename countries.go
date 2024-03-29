package countries

import (
	"strings"
)

// AllMappings returns the list of all country mappings.
func AllMappings() []Mapping {
	return mappings
}

// FindCountry looks up any matching occurrence of the query.
func FindCountry(query string) (*Mapping, error) {
	const (
		alpha2Len = 2
		alpha3Len = 3
	)

	switch {
	case len(query) < alpha2Len:
		return nil, ErrCountryNotFound
	case len(query) == alpha2Len:
		return findCountryByAlpha2(query)
	case len(query) == alpha3Len:
		return findCountryByAlpha3(query)
	default:
		return findCountryByNameOrNationality(query)
	}
}

func findCountryByAlpha2(query string) (*Mapping, error) {
	query = strings.ToUpper(query)

	for i := range mappings {
		if mappings[i].Alpha2 == query {
			return &mappings[i], nil
		}
	}

	return nil, ErrCountryNotFound
}

func findCountryByAlpha3(query string) (*Mapping, error) {
	query = strings.ToUpper(query)

	for i := range mappings {
		if mappings[i].Alpha3 == query {
			return &mappings[i], nil
		}
	}

	return nil, ErrCountryNotFound
}

func findCountryByNameOrNationality(query string) (*Mapping, error) {
	for i := range mappings {
		if isCountryNameOrNationality(mappings[i].Translations[EN], query) ||
			isCountryNameOrNationality(mappings[i].Translations[DE], query) {
			return &mappings[i], nil
		}
	}

	return nil, ErrCountryNotFound
}

func isCountryNameOrNationality(translation Translation, query string) bool {
	return strings.EqualFold(translation.Common, query) ||
		strings.EqualFold(translation.Official, query) ||
		strings.EqualFold(translation.Nationality, query)
}

// Alpha2 looks up any matching occurrence for the query and returns the ISO-3166-1 Alpha-2 code.
func Alpha2(query string) string {
	country, err := FindCountry(query)
	if err != nil {
		return ""
	}

	return country.Alpha2
}

// Alpha3 looks up any matching occurrence for the query and returns the ISO-3166-1 Alpha-3 code.
func Alpha3(query string) string {
	country, err := FindCountry(query)
	if err != nil {
		return ""
	}

	return country.Alpha3
}

// CountryName looks up any matching occurrence for the query and returns the official country name in english.
func CountryName(query string) string {
	country, err := FindCountry(query)
	if err != nil {
		return ""
	}

	return country.Translations[EN].Common
}

// CountryTranslation looks up any matching occurrence for the query and returns the country translation.
func CountryTranslation(query string, lang Language) (*Translation, error) {
	country, err := FindCountry(query)
	if err != nil {
		return nil, ErrCountryNotFound
	}

	translation, ok := country.Translations[lang]
	if !ok {
		return nil, ErrTranslationNotFound
	}

	return &translation, nil
}
