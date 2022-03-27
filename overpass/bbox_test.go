package overpass

import (
	"testing"
)

func TestValidBBox(t *testing.T) {
	bbox := BBox{South: 40.771900, West: -73.974600, North: 40.797500, East: -73.946900}
	bbox.Key = "public_transport"
	bbox.Value = "station"

	if !bbox.Valid() {
		t.Fatal(bbox)
	}

	_, err := bbox.Query(TestServer)
	if err != nil {
		t.Fatal(err)
	}
}
