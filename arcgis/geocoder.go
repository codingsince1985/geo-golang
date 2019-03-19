// Package arcgis is a geo-golang based ArcGIS geocode/reverse client
package arcgis

import (
	"fmt"
	"strings"

	geo "github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		Candidates []struct {
			Address  string
			Location struct {
				X float64
				Y float64
			}
		}

		ReverseAddress struct {
			MatchAddr    string `json:"Match_addr"`
			LongLabel    string
			ShortLabel   string
			AddNum       string
			Address      string
			Neighborhood string
			City         string
			Subregion    string
			Region       string
			Postal       string
			CountryCode  string
		} `json:"address"`
	}
)

// Geocoder constructs ArcGIS geocoder
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

	url := "http://geocode.arcgis.com/arcgis/rest/services/World/GeocodeServer/*"
	if len(token) > 0 {
		url += "&token=" + token
	}
	return url
}

func (b baseURL) GeocodeURL(address string) string {
	params := fmt.Sprintf("findAddressCandidates?f=json&maxLocations=%d&SingleLine=%s", 1, address)
	return strings.Replace(string(b), "*", params, 1)
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	params := fmt.Sprintf("reverseGeocode?f=json&location=%f,%f", l.Lng, l.Lat)
	return strings.Replace(string(b), "*", params, 1)
}

func (r *geocodeResponse) Location() (*geo.Location, error) {
	if len(r.Candidates) == 0 {
		return nil, nil
	}

	g := r.Candidates[0].Location
	return &geo.Location{
		Lat: g.Y,
		Lng: g.X,
	}, nil
}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	addr := &geo.Address{
		FormattedAddress: r.ReverseAddress.MatchAddr,
		Street:           r.ReverseAddress.Address,
		HouseNumber:      r.ReverseAddress.AddNum,
		Postcode:         r.ReverseAddress.Postal,
		State:            r.ReverseAddress.Region,
		CountryCode:      r.ReverseAddress.CountryCode,
	}

	return addr, nil
}
