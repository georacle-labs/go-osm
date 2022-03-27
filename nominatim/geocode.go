package nominatim

import (
	"encoding/json"
	"net/url"
	"path"
)

// GeocodeQuery represents a Nominatim geocode query
type GeocodeQuery struct {
	Address    string `json:"address"`
	Street     string `json:"street"`
	City       string `json:"city"`
	County     string `json:"county"`
	State      string `json:"state"`
	Country    string `json:"country"`
	Postalcode string `json:"postalcode"`
}

// Valid checks for a valid geocode query
func (g *GeocodeQuery) Valid() bool {
	return g.Address != "" ||
		((g.Street != "") && (g.City != "") && (g.State != "") &&
			(g.Country != "") && (g.Postalcode != ""))
}

// BuildQuery constructs a geocode query
func (g *GeocodeQuery) BuildQuery() (*url.Values, error) {
	params := &url.Values{}
	if g.Address != "" {
		params.Add("q", g.Address)
	} else {
		params.Add("street", g.Street)
		params.Add("city", g.City)
		params.Add("county", g.County)
		params.Add("state", g.State)
		params.Add("country", g.Country)
		params.Add("postalcode", g.Postalcode)
	}
	return params, nil
}

// Query performs a geocode query and returns the response
func (g *GeocodeQuery) Query(server string) ([]Response, error) {
	var result []Response

	base, err := url.Parse(server)
	if err != nil {
		return result, err
	}

	base.Path = path.Join(base.Path, "search")

	queryParams, err := g.BuildQuery()
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
