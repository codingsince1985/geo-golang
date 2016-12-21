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
		}
	}
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
	return "https://api.mapbox.com/geocoding/v5/mapbox.places/*.json?access_token=" + token
}

func (b baseURL) GeocodeURL(address string) string { return strings.Replace(string(b), "*", address, 1) }

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return strings.Replace(string(b), "*", fmt.Sprintf("%+f,%+f", l.Lng, l.Lat), 1)
}

func (r *geocodeResponse) Location() geo.Location {
	if len(r.Features) > 0 {
		g := r.Features[0]
		return geo.Location{g.Center[1], g.Center[0]}
	}
	return geo.Location{}
}

func (r *geocodeResponse) Address() string {
	if len(r.Features) > 0 {
		return r.Features[0].PlaceName
	}
	return ""
}
