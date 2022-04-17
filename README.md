<p align="center">
  <img src="img.png" />
</p>

# Go OpenStreetMap

[![Go Reference](https://pkg.go.dev/badge/github.com/georacle-labs/go-osm.svg)](https://pkg.go.dev/github.com/georacle-labs/go-osm)
[![Go Report Card](https://goreportcard.com/badge/github.com/georacle-labs/go-osm)](https://goreportcard.com/report/github.com/georacle-labs/go-osm)
[![CircleCI](https://circleci.com/gh/georacle-labs/go-osm/tree/main.svg?style=shield)](https://circleci.com/gh/georacle-labs/go-osm/tree/main)

The Go Client for [OpenStreetMap](https://www.openstreetmap.org/).

This package provides an open source framework for **location search** and **geocoding** using the OpenStreetMap dataset.

# Features

- [X] Area Search
- [x] Bounding Box Search
- [x] Proximity Search
- [X]  Geocoding
- [X] Reverse Geocoding
- [X] Geometry Sampling

# Usage

```go
package main

import (
        "fmt"
        "log"

        "github.com/georacle-labs/go-osm/nominatim"
        "github.com/georacle-labs/go-osm/overpass"
)

func main() {
        // location query
        Tags := overpass.Tags{Key: "leisure", Value: "park"}
        a := &overpass.Area{Name: "London", Tags: Tags}
        location, err := a.Query("https://overpass-api.de/api/interpreter")
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println(location)

        // geocode
        g := &nominatim.GeocodeQuery{Address: "221B Baker St, London NW1 6XE, UK"}
        geocode, err := g.Query("https://nominatim.openstreetmap.org/")
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println(geocode)
}
```

# Backends

| Package                                                                                          | Functionality                                                                                            |
|------------------------------------------------------------------------------------- |----------------------------------------------------------------------------------------------|
| [Nominatim](https://github.com/osm-search/Nominatim)  | Fuzzy Search, Geocoding, Reverse Geocoding                                  |
| [Overpass](https://github.com/drolbr/Overpass-API)           | Location Search and Geometry Sampling                                           |


# Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/georacle-labs/go-osm

# License

The package is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT)
