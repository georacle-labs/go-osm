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

// BuildQuery constructs a proximity query
func (p *Proximity) BuildQuery() (string, error) {
	lat := p.Lat * math.Pi / 180 // radians
	lon := p.Lon * math.Pi / 180 // radians

	// construct a bounding box around the ref point
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

	return b.BuildQuery()
}

// Query makes a proximity query with overpass
func (p *Proximity) Query(server string) (*Response, error) {
	query, err := p.BuildQuery()
	if err != nil {
		return nil, err
	}

	resp, err := Overpass(server, []byte(query))
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
