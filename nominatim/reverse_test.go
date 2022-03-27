package nominatim

import (
	"go-osm/geometry"
	"strings"
	"testing"
)

var TestPoint = geometry.Point{Lat: 51.5233879, Lon: -0.1582367}

func TestValidReverseQuery(t *testing.T) {
	g := &ReverseGeocodeQuery{TestPoint}

	if !g.Valid() {
		t.Fatal("Invalid query")
	}

	resp, err := g.Query(TestServer)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(resp.Name, strings.Split(TestAddress, ",")[0]) {
		t.Fatal(resp.Name)
	}
}
