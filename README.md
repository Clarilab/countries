# ClariLab Countries

This library provides a list of mappings for country <=> iso-3166-1 alpha 2 <=> iso-3166-1 alpha 3. 
With the given functions, you can search with any query and it will try to be matched to a country.

## Usage:

### Available functions

```go
    GetAllMappings()
    FindCountry(query string)
    GetCountryName(query string)
    GetAlpha2(query string)
    GetAlpha3(query string)
```

### Examples

```go
    countryName := countries.GetCountryName("DE")
    countryName = countries.GetCountryName("DEU")
    countryName = countries.GetCountryName("Germany")
```
will all result in "Germany". If the country couldn't be found with the given query, your result will be an empty string.

You can also retrieve the object of type Mapping with the FindCountry function.

```go
    countryMapping := countries.FindCountry("DE")
```