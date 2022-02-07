package countries

import (
	"errors"
	"sort"
)

// Mapping holds a country code.
type Mapping struct {
	Country string `json:"country" example:"Germany"`
	Alpha2  string `json:"alpha2" example:"DE"`
	Alpha3  string `json:"alpha3" example:"DEU"`
}

// ErrCountryNotFound indicates that the searched country was not found in our list.
var ErrCountryNotFound = errors.New("country not found")

// GetAllMappings returns the list of all country mappings.
func GetAllMappings() []Mapping {
	return mappings
}

// FindCountry looks up any matching occurence of the query.
func FindCountry(query string) (*Mapping, error) {
	switch {
	case len(query) < 2:
		return nil, ErrCountryNotFound
	case len(query) == 2:
		return findCountryByAlpha2(query)
	case len(query) == 3:
		return findCountryByAlpha3(query)
	default:
		return findCountryByName(query)
	}
}

func findCountryByAlpha2(query string) (*Mapping, error) {
	for i := range mappings {
		if mappings[i].Alpha2 == query {
			return &mappings[i], nil
		}
	}

	return nil, ErrCountryNotFound
}

func findCountryByAlpha3(query string) (*Mapping, error) {
	for i := range mappings {
		if mappings[i].Alpha3 == query {
			return &mappings[i], nil
		}
	}

	return nil, ErrCountryNotFound
}

func findCountryByName(query string) (*Mapping, error) {
	idx := sort.Search(len(mappings), func(i int) bool {
		return mappings[i].Country >= query
	})

	if idx < len(mappings) && mappings[idx].Country == query {
		return &mappings[idx], nil
	}

	return nil, ErrCountryNotFound
}

// GetAlpha2 looks up any matching occurence for the query and returns the ISO-3166-1 Alpha-2 code.
func GetAlpha2(query string) string {
	country, err := FindCountry(query)
	if err != nil {
		return ""
	}

	return country.Alpha2
}

// GetAlpha3 looks up any matching occurence for the query and returns the ISO-3166-1 Alpha-3 code.
func GetAlpha3(query string) string {
	country, err := FindCountry(query)
	if err != nil {
		return ""
	}

	return country.Alpha3
}

// GetCountryName looks up any matching occurence for the query and returns the country name in english.
func GetCountryName(query string) string {
	country, err := FindCountry(query)
	if err != nil {
		return ""
	}

	return country.Country
}
