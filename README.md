# ClariLab Countries

This library provides a list of mappings for country <=> iso-3166-1 alpha 2 <=> iso-3166-1 alpha 3.
With the given functions, you can search with any query and it will try to be matched to a country.

## Dependencies

This library has zero dependencies and aims to also have none in the future.

## Usage

### Available functions

```go
    GetAllMappings() []Mapping
    FindCountry(query string) (*Mapping, error)
    GetCountryName(query string) string
    GetAlpha2(query string) string
    GetAlpha3(query string) string
    GetCountryTranslation(query string, lang Language) (*Translation, error)
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
    countryMapping, err := countries.FindCountry("DE")
    // handle error case
```

## Benchmarks

Because we're doing a linear search over the country list, the worst case performance is O(n).

In all benchmarks, Australia is the best case scenario as it's the first entry of the country list and Japan is the worst case scenario as it's the last entry of the country list.

### Find country by country name

```text
BenchmarkGetByCountryName/input_Australia-8             63535485                18.57 ns/op            0 B/op          0 allocs/op
BenchmarkGetByCountryName/input_Australia-8             64542153                18.51 ns/op            0 B/op          0 allocs/op
BenchmarkGetByCountryName/input_Australia-8             63800798                18.53 ns/op            0 B/op          0 allocs/op
BenchmarkGetByCountryName/input_Australia-8             62710806                18.74 ns/op            0 B/op          0 allocs/op
BenchmarkGetByCountryName/input_Australia-8             63743749                18.73 ns/op            0 B/op          0 allocs/op
BenchmarkGetByCountryName/input_Japan-8                   150270              7703 ns/op               0 B/op          0 allocs/op
BenchmarkGetByCountryName/input_Japan-8                   157118              7591 ns/op               0 B/op          0 allocs/op
BenchmarkGetByCountryName/input_Japan-8                   152257              7612 ns/op               0 B/op          0 allocs/op
BenchmarkGetByCountryName/input_Japan-8                   158622              7559 ns/op               0 B/op          0 allocs/op
BenchmarkGetByCountryName/input_Japan-8                   158541              7608 ns/op               0 B/op          0 allocs/op
```

### Find country by ISO 3166-1 code

```text
BenchmarkGetByAlpha2/input_AU-8                 127186076                9.391 ns/op           0 B/op          0 allocs/op
BenchmarkGetByAlpha2/input_AU-8                 128749902                9.370 ns/op           0 B/op          0 allocs/op
BenchmarkGetByAlpha2/input_AU-8                 128010331                9.315 ns/op           0 B/op          0 allocs/op
BenchmarkGetByAlpha2/input_AU-8                 128852181                9.333 ns/op           0 B/op          0 allocs/op
BenchmarkGetByAlpha2/input_AU-8                 127954191                9.460 ns/op           0 B/op          0 allocs/op
BenchmarkGetByAlpha2/input_JP-8                  1620708               779.1 ns/op             0 B/op          0 allocs/op
BenchmarkGetByAlpha2/input_JP-8                  1479526               740.4 ns/op             0 B/op          0 allocs/op
BenchmarkGetByAlpha2/input_JP-8                  1624162               740.7 ns/op             0 B/op          0 allocs/op
BenchmarkGetByAlpha2/input_JP-8                  1613852               749.5 ns/op             0 B/op          0 allocs/op
BenchmarkGetByAlpha2/input_JP-8                  1591786               737.8 ns/op             0 B/op          0 allocs/op
```

```text
BenchmarkGetByAlpha3/input_AUS-8                126209905                9.466 ns/op           0 B/op          0 allocs/op
BenchmarkGetByAlpha3/input_AUS-8                125323408                9.592 ns/op           0 B/op          0 allocs/op
BenchmarkGetByAlpha3/input_AUS-8                123060835               10.15 ns/op            0 B/op          0 allocs/op
BenchmarkGetByAlpha3/input_AUS-8                121636543                9.583 ns/op           0 B/op          0 allocs/op
BenchmarkGetByAlpha3/input_AUS-8                125560383                9.513 ns/op           0 B/op          0 allocs/op
BenchmarkGetByAlpha3/input_JPN-8                 1654291               724.6 ns/op             0 B/op          0 allocs/op
BenchmarkGetByAlpha3/input_JPN-8                 1652850               722.2 ns/op             0 B/op          0 allocs/op
BenchmarkGetByAlpha3/input_JPN-8                 1400821               728.3 ns/op             0 B/op          0 allocs/op
BenchmarkGetByAlpha3/input_JPN-8                 1644262               732.0 ns/op             0 B/op          0 allocs/op
BenchmarkGetByAlpha3/input_JPN-8                 1598776               783.7 ns/op             0 B/op          0 allocs/op
```
