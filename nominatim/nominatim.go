package nominatim

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Response represents a Nominatim query response
type Response struct {
	PlaceID     int64             `json:"place_id"`
	Licence     string            `json:"licence"`
	OsmType     string            `json:"osm_type"`
	OsmID       int64             `json:"osm_id"`
	Lat         string            `json:"lat"`
	Lon         string            `json:"lon"`
	PlaceRank   int64             `json:"place_rank"`
	Category    string            `json:"category"`
	Type        string            `json:"type"`
	Importance  float64           `json:"importance"`
	AddressType string            `json:"addresstype"`
	Name        string            `json:"name"`
	DisplayName string            `json:"display_name"`
	Address     map[string]string `json:"address"`
	BoundingBox []string          `json:"boundingbox"`
}

// Nominatim performs a Nominatim query
func Nominatim(query string) ([]byte, error) {
	start := time.Now()

	resp, err := http.Get(query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	elapsed := time.Now().Sub(start)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("[Nominatim] Server Responded: %d bytes (%s)", len(body), elapsed)
	return body, nil
}
