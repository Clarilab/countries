package countries

import (
	"strings"
)

// GetAllMappings returns the list of all country mappings.
func GetAllMappings() []Mapping {
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
		if isCountryNameOrNationality(mappings[i].Translations[string(EN)], query) ||
			isCountryNameOrNationality(mappings[i].Translations[string(DE)], query) {
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

// GetAlpha2 looks up any matching occurrence for the query and returns the ISO-3166-1 Alpha-2 code.
func GetAlpha2(query string) string {
	country, err := FindCountry(query)
	if err != nil {
		return ""
	}

	return country.Alpha2
}

// GetAlpha3 looks up any matching occurrence for the query and returns the ISO-3166-1 Alpha-3 code.
func GetAlpha3(query string) string {
	country, err := FindCountry(query)
	if err != nil {
		return ""
	}

	return country.Alpha3
}

// GetCountryName looks up any matching occurrence for the query and returns the official country name in english.
func GetCountryName(query string) string {
	country, err := FindCountry(query)
	if err != nil {
		return ""
	}

	return country.Translations[string(EN)].Common
}

// GetCountryTranslation looks up any matching occurrence for the query and returns the country translation.
func GetCountryTranslation(query string, lang Language) (*Translation, error) {
	country, err := FindCountry(query)
	if err != nil {
		return nil, ErrCountryNotFound
	}

	translation, ok := country.Translations[string(lang)]
	if !ok {
		return nil, ErrTranslationNotFound
	}

	return &translation, nil
}
