// Package mapbox is a geo-golang based Mapbox geocode/reverse geocode client
package mapbox

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"strings"
)

type baseURL string

type geocodeResponse struct {
	Features []struct {
		PlaceName string `json:"place_name"`
		Center    [2]float64
	}
}

// Geocoder constructs Mapbox geocoder
func Geocoder(token string) geo.Geocoder {
	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL("https://api.mapbox.com/geocoding/v5/mapbox.places/*.json?access_token=" + token),
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func (b baseURL) GeocodeURL(address string) string {
	return strings.Replace(string(b), "*", address, 1)
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return strings.Replace(string(b), "*", fmt.Sprintf("%+f,%+f", l.Lng, l.Lat), 1)
}

func (r *geocodeResponse) Location() (l geo.Location) {
	if len(r.Features) > 0 {
		g := r.Features[0]
		l = geo.Location{g.Center[1], g.Center[0]}
	}
	return
}

func (r *geocodeResponse) Address() (address string) {
	if len(r.Features) > 0 {
		address = r.Features[0].PlaceName
	}
	return
}
