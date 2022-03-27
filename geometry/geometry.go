package geometry

import (
	"math"
)

const (
	// MinLat is the minimum valid latitude (in radians)
	MinLat = -90 * math.Pi / 180.0

	// MaxLat is the maximum valid latitude (in radians)
	MaxLat = 90 * math.Pi / 180.0

	// MinLon is the minimum valid longitude (in radians)
	MinLon = -180 * math.Pi / 180.0

	// MaxLon is the maximum valid longitude (in radians)
	MaxLon = 180 * math.Pi / 180.0
)

// Point represents a single coordinate
type Point struct {
	Lat float64 `json:"lat" form:"lat"`
	Lon float64 `json:"lon" form:"lon"`
}

// Valid checks for a valid coordinate
func (p *Point) Valid() bool {
	latRad := p.Lat * math.Pi / 180.0
	lonRad := p.Lon * math.Pi / 180.0

	validLat := latRad <= MaxLat && latRad >= MinLat
	validLon := lonRad <= MaxLon && lonRad >= MinLon

	return validLat && validLon
}
