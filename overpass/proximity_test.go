package overpass

import (
	"testing"

	"github.com/georacle-labs/go-osm/geometry"
)

func TestValidProximity(t *testing.T) {
	p := Proximity{Radius: 500}
	p.Point = geometry.Point{Lat: 48.858093, Lon: 2.294694}
	p.Tags = Tags{Key: "amenity", Value: "restaurant"}

	if !p.Valid() {
		t.Fatal(p)
	}

	_, err := p.Query(TestServer)
	if err != nil {
		t.Fatal(err)
	}
}
