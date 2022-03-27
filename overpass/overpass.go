package overpass

import (
	"bytes"
	"encoding/json"
	"go-osm/geometry"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Response represents an overpass query response
type Response struct {
	Version   float64           `json:"version"`
	Generator string            `json:"generator"`
	Osm3s     map[string]string `json:"osm3s"`
	Elements  []Element         `json:"elements"`
}

// Element represents a single response element
type Element struct {
	Type     string             `json:"type"`
	ID       int64              `json:"id"`
	Center   geometry.Point     `json:"center"`
	Bounds   map[string]float64 `json:"bounds"`
	Nodes    []int64            `json:"nodes"`
	Members  []Member           `json:"members"`
	Geometry []geometry.Point   `json:"geometry"`
	Lat      float64            `json:"lat"`
	Lon      float64            `json:"lon"`
	Tags     map[string]string  `json:"tags"`
}

// Member represents an OSM relation member
type Member struct {
	Type string `json:"type"`
	Ref  int64  `json:"ref"`
	Role string `json:"role"`
}

// Tags represents an osm key-value pair
// see: https://taginfo.openstreetmap.org/
type Tags struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Limit uint   `json:"limit"`
}

// Overpass sends an overpass query
func Overpass(server string, query []byte) (*Response, error) {
	postBody := bytes.NewBuffer(query)

	start := time.Now()
	resp, err := http.Post(server, "application/json", postBody)
	if err != nil {
		return nil, err

	}
	defer resp.Body.Close()

	elapsed := time.Now().Sub(start)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("Server Responded: %d bytes (%s)", len(body), elapsed)

	res := new(Response)
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}
