package overpass

import (
	"math"

	geo "github.com/georacle-labs/go-osm/geometry"
)

// Proximity represents a proximity query
type Proximity struct {
	Tags
	geo.Point
	Radius float64 `json:"radius"`
}

// Valid checks for a valid proximity query
func (p *Proximity) Valid() bool {
	return p.Radius > 0
}

// BBox constructs a bounding box around a reference point
func (p *Proximity) BBox() *BBox {
	lat := p.Lat * math.Pi / 180 // radians
	lon := p.Lon * math.Pi / 180 // radians

	rad := p.Radius / geo.RadEarth
	minLon := geo.MinLon
	maxLon := geo.MaxLon
	minLat := lat - rad
	maxLat := lat + rad

	// compute bounds
	if minLat > geo.MinLat && maxLat < geo.MaxLat {
		dLon := math.Asin(math.Sin(rad) / math.Cos(lat))
		minLon = lon - dLon
		maxLon = lon + dLon

		// normalize
		if minLon < geo.MinLon {
			minLon += 2 * math.Pi
		}
		if maxLon > geo.MaxLon {
			maxLon -= 2 * math.Pi
		}
	} else {
		minLat = math.Max(minLat, geo.MinLat)
		maxLat = math.Max(maxLat, geo.MaxLat)
	}

	b := &BBox{Tags: p.Tags}
	b.South = minLat * 180 / math.Pi
	b.West = minLon * 180 / math.Pi
	b.North = maxLat * 180 / math.Pi
	b.East = maxLon * 180 / math.Pi

	return b
}

// Query makes a proximity query with overpass
func (p *Proximity) Query(server string) (*Response, error) {
	bbox := p.BBox()
	resp, err := bbox.Query(server)
	if err != nil {
		return nil, err
	}

	refPoint := geo.Point{Lat: p.Lat, Lon: p.Lon}
	validNodes := make([]Element, 0)

	// filter valid nodes
	for _, e := range resp.Elements {
		pt := &geo.Point{Lat: e.Lat, Lon: e.Lon}
		dist := refPoint.SphericalDistance(pt)
		if dist <= p.Radius {
			validNodes = append(validNodes, e)
		}
	}

	resp.Elements = validNodes
	return resp, nil
}
