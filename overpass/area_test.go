package overpass

import (
	"os"
	"testing"
)

var TestServer = os.Getenv("OVERPASS_SERVER")

func TestValidArea(t *testing.T) {

	a := &Area{Name: "London", Tags: Tags{Key: "leisure", Value: "park"}}

	if !a.Valid() {
		t.Fatal(a)
	}

	_, err := a.Query(TestServer)
	if err != nil {
		t.Fatal(err)
	}
}
