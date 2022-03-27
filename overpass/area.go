package overpass

import (
	"errors"
	"fmt"
	"strconv"
)

// Area represents an overpass named area query
type Area struct {
	Tags
	Name string `json:"name"`
}

// Valid checks for a valid named area
func (a *Area) Valid() bool {
	return len(a.Name) > 0
}

// BuildQuery constructs a named area query
func (a *Area) BuildQuery() (string, error) {
	var query string

	if !a.Valid() {
		return query, errors.New("InvalidAreaQuery")
	}

	query = fmt.Sprintf("[out:json];area[name=\"%s\"];", a.Name)
	query += fmt.Sprintf("nwr[%s=%s](area);out", a.Key, a.Value)

	if a.Limit > 0 {
		query += (" " + strconv.Itoa(int(a.Limit)))
	}

	return query + ";", nil
}

// Query makes a named area query with overpass
func (a *Area) Query(server string) (*Response, error) {
	query, err := a.BuildQuery()
	if err != nil {
		return nil, err
	}

	return Overpass(server, []byte(query))
}
