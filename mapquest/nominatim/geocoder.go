// Package nominatim is a geo-golang based MapRequest Nominatim geocode/reverse geocode client
package nominatim

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		DisplayName     string `json:"display_name"`
		Lat, Lon, Error string
	}
)

var key string

// Geocoder constructs MapRequest Nominatim geocoder
func Geocoder(k string, baseURLs ...string) geo.Geocoder {
	key = k
	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL(getUrl(baseURLs...)),
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func getUrl(baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	return "http://open.mapquestapi.com/nominatim/v1/"
}

func (b baseURL) GeocodeURL(address string) string {
	return string(b) + "search.php?key=" + key + "&format=json&limit=1&q=" + address
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return string(b) + "reverse.php?key=" + key + fmt.Sprintf("&format=json&lat=%f&lon=%f", l.Lat, l.Lng)
}

func (r *geocodeResponse) Location() geo.Location {
	if r.Error == "" {
		return geo.Location{geo.ParseFloat(r.Lat), geo.ParseFloat(r.Lon)}
	}
	return geo.Location{}
}

func (r *geocodeResponse) Address() string {
	if r.Error == "" {
		return r.DisplayName
	}
	return ""
}
