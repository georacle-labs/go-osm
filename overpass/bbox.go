package overpass

import (
	"errors"
	"fmt"
	"strconv"
)

// BBox represents a bounding box
type BBox struct {
	Tags
	South float64 `json:"south"`
	West  float64 `json:"west"`
	North float64 `json:"north"`
	East  float64 `json:"east"`
}

// Valid checks for a valid bounding box
func (b *BBox) Valid() bool {
	return b.South < b.North && b.West < b.East
}

// BuildQuery constructs a bounding box query
func (b *BBox) BuildQuery() (string, error) {
	var query string

	if !b.Valid() {
		return query, errors.New("InvalidBBoxQuery")
	}

	query = fmt.Sprintf("[out:json];nwr[%s=%s]", b.Key, b.Value)
	query += fmt.Sprintf("(%f, %f, %f, %f);out", b.South, b.West, b.North, b.East)
	if b.Limit > 0 {
		query += (" " + strconv.Itoa(int(b.Limit)))
	}

	return query + ";", nil
}

// Query makes a named bbox query with overpass
func (b *BBox) Query(server string) (*Response, error) {
	query, err := b.BuildQuery()
	if err != nil {
		return nil, err
	}

	return Overpass(server, []byte(query))
}
