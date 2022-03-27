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

	// RadEarth is the average radius of Earth (in meters)
	RadEarth = 6378100
)

// Point represents a single coordinate
type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// Valid checks for a valid coordinate
func (p *Point) Valid() bool {
	latRad := p.Lat * math.Pi / 180.0
	lonRad := p.Lon * math.Pi / 180.0

	validLat := latRad <= MaxLat && latRad >= MinLat
	validLon := lonRad <= MaxLon && lonRad >= MinLon

	return validLat && validLon
}

// SphericalDistance computes the great circle distance between two points
// Haversine formula:  a = sin²(Δφ/2) + cos φ1 ⋅ cos φ2 ⋅ sin²(Δλ/2)
// c = 2 ⋅ atan2( √a, √(1−a) )
// d = R ⋅ c
// where φ is latitude, λ is longitude, R is earth’s avg radius
func (p *Point) SphericalDistance(p2 *Point) float64 {
	phi1 := p.Lat * math.Pi / 180.0  // radians
	phi2 := p2.Lat * math.Pi / 180.0 // radians
	deltaPhi := phi2 - phi1
	deltaLam := (p2.Lon - p.Lon) * math.Pi / 180.0

	a := math.Pow(math.Sin(deltaPhi/2), 2)
	b := math.Cos(phi1) * math.Cos(phi2)
	b *= math.Pow(math.Sin(deltaLam/2), 2)
	a += b

	c := 2.0 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return RadEarth * c // meters
}
