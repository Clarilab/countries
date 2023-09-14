package countries

import "errors"

// Language is a ISO 639-1 language code.
type Language string

const (
	// DE is the german ISO 639-1 language code.
	DE Language = "de"
	// EN is for english ISO 639-1 language code.
	EN Language = "en"
)

// Mapping holds a country code.
type Mapping struct {
	Alpha2       string                   `json:"alpha2" example:"DE"`
	Alpha3       string                   `json:"alpha3" example:"DEU"`
	Translations map[Language]Translation `json:"translations"`
}

// Translation is a single translation for a country.
type Translation struct {
	Common      string `json:"common" example:"Germany"`
	Official    string `json:"official" example:"Germany"`
	Nationality string `json:"nationality" example:"German"`
}

// ErrCountryNotFound indicates that the searched country was not found in our list.
var ErrCountryNotFound = errors.New("country not found")

// ErrTranslationNotFound indicates that the translation for the given language was not found in our list.
var ErrTranslationNotFound = errors.New("translation not found")
