// Package mapbox is a geo-golang based Mapbox geocode/reverse geocode client
package mapbox

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"strings"
)

type (
	baseURL         string
	geocodeResponse struct {
		Features []struct {
			PlaceName string `json:"place_name"`
			Center    [2]float64
			Text      string `json:"text"`    // usually street name
			Address   string `json:"address"` // potentially house number
			Context   []struct {
				Text      string `json:"text"`
				Id        string `json:"id"`
				ShortCode string `json:"short_code"`
				Wikidata  string `json:"wikidata"`
			}
		}
		Message string `json:"message"`
	}
)

const (
	mapboxPrefixLocality = "place"
	mapboxPrefixPostcode = "postcode"
	mapboxPrefixState    = "region"
	mapboxPrefixCountry  = "country"
)

// Geocoder constructs Mapbox geocoder
func Geocoder(token string, baseURLs ...string) geo.Geocoder {
	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL(getUrl(token, baseURLs...)),
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func getUrl(token string, baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	return "https://api.mapbox.com/geocoding/v5/mapbox.places/*.json?limit=1&access_token=" + token
}

func (b baseURL) GeocodeURL(address string) string { return strings.Replace(string(b), "*", address, 1) }

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return strings.Replace(string(b), "*", fmt.Sprintf("%+f,%+f", l.Lng, l.Lat), 1)
}

func (r *geocodeResponse) Location() (*geo.Location, error) {
	if len(r.Features) == 0 {
		// error in response
		if r.Message != "" {
			return nil, fmt.Errorf("reverse geocoding error: %s", r.Message)
		} else { // no results
			return nil, nil
		}
	}

	g := r.Features[0]
	return &geo.Location{
		Lat: g.Center[1],
		Lng: g.Center[0],
	}, nil
}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	if len(r.Features) == 0 {
		// error in response
		if r.Message != "" {
			return nil, fmt.Errorf("reverse geocoding error: %s", r.Message)
		} else { // no results
			return nil, nil
		}
	}

	return parseMapboxResponse(r), nil
}

func parseMapboxResponse(r *geocodeResponse) *geo.Address {
	addr := &geo.Address{}
	f := r.Features[0]
	addr.FormattedAddress = f.PlaceName
	addr.Street = f.Text
	addr.HouseNumber = f.Address
	for _, c := range f.Context {
		if strings.HasPrefix(c.Id, mapboxPrefixLocality) {
			addr.City = c.Text
		} else if strings.HasPrefix(c.Id, mapboxPrefixPostcode) {
			addr.Postcode = c.Text
		} else if strings.HasPrefix(c.Id, mapboxPrefixState) {
			addr.State = c.Text
		} else if strings.HasPrefix(c.Id, mapboxPrefixCountry) {
			addr.Country = c.Text
			addr.CountryCode = strings.ToUpper(c.ShortCode)
		}
	}

	return addr
}
