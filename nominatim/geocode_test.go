package nominatim

import (
	"os"
	"strconv"
	"testing"
)

var TestServer = os.Getenv("NOMINATIM_SERVER")
var TestAddress = "221B Baker St, London NW1 6XE, UK"

func TestValidQuery(t *testing.T) {
	g := &GeocodeQuery{Address: TestAddress}
	if !g.Valid() {
		t.Fatal(g)
	}

	response, err := g.Query(TestServer)
	if err != nil {
		t.Fatal(err)
	}

	if len(response) <= 0 {
		t.Fatal(response)
	}

	for _, resp := range response {
		if resp.Lat == resp.Lon {
			t.Fatal(resp)
		}
	}

	resp := response[0]

	lat, err := strconv.ParseFloat(resp.Lat, 64)
	if err != nil {
		t.Fatal(err)
	}

	lon, err := strconv.ParseFloat(resp.Lon, 64)
	if err != nil {
		t.Fatal(err)
	}

	if lat != TestPoint.Lat || lon != TestPoint.Lon {
		t.Fatal(lat, lon)
	}
}
