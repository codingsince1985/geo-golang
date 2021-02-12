package geocod

import (
	"fmt"
	"strings"

	geo "github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		Results []struct {
			Components struct {
				Number  string
				Street  string
				City    string
				County  string
				State   string
				Zip     string
				Country string
			} `json:"address_components"`
			Address  string `json:"formatted_address"`
			Location struct {
				Lat float64
				Lng float64
			}
		}
	}
)

// Geocoder constructs Geocodio geocoder
func Geocoder(key string, baseURLs ...string) geo.Geocoder {
	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL(getUrl(key, baseURLs...)),
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func getUrl(key string, baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	return "https://api.geocod.io/v1/*&api_key=" + key
}

func (b baseURL) GeocodeURL(address string) string {
	params := fmt.Sprintf("geocode?q=%s", address)
	url := strings.Replace(string(b), "*", params, 1)
	return url
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	params := fmt.Sprintf("reverse?q=%f,%f", l.Lat, l.Lng)
	url := strings.Replace(string(b), "*", params, 1)
	return url
}

func (r *geocodeResponse) Location() (*geo.Location, error) {
	if len(r.Results) == 0 {
		return nil, nil
	}

	loc := r.Results[0].Location
	return &geo.Location{Lat: loc.Lat, Lng: loc.Lng}, nil
}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	if len(r.Results) == 0 {
		return nil, nil
	}

	r0 := r.Results[0]
	c := r0.Components
	addr := &geo.Address{
		FormattedAddress: r0.Address,
		Street:           c.Street,
		HouseNumber:      c.Number,
		Postcode:         c.Zip,
		State:            c.State,
		StateCode:        c.State,
		CountryCode:      c.Country,
	}

	return addr, nil
}
