package nominatim

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"path"

	"go-osm/geometry"
)

// ReverseGeocodeQuery represents a Nominatim reverse geocode query
type ReverseGeocodeQuery struct {
	geometry.Point
}

// BuildQuery constructs a reverse geocode query for Nominatim
func (r *ReverseGeocodeQuery) BuildQuery() (*url.Values, error) {
	if !r.Valid() {
		return nil, errors.New("InvalidReverseGeocodeQuery")
	}

	params := &url.Values{}
	params.Add("lat", fmt.Sprintf("%f", r.Lat))
	params.Add("lon", fmt.Sprintf("%f", r.Lon))
	return params, nil
}

// Query performs a reverse geocode query and returns the response
func (r *ReverseGeocodeQuery) Query(server string) (Response, error) {
	var result Response

	base, err := url.Parse(server)
	if err != nil {
		return result, err
	}

	base.Path = path.Join(base.Path, "reverse")

	queryParams, err := r.BuildQuery()
	if err != nil {
		return result, err
	}

	queryParams.Add("format", "jsonv2")
	base.RawQuery = queryParams.Encode()

	resp, err := Nominatim(base.String())
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
