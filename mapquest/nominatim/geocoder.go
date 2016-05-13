// Package nominatim is a geo-golang based MapRequest Nominatim geocode/reverse geocode client
package nominatim

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"strconv"
)

type baseURL string

type geocodeResponse struct {
	DisplayName     string `json:"display_name"`
	Lat, Lon, Error string
}

var key string

// Geocoder constructs MapRequest Nominatim geocoder
func Geocoder(k string, baseURLs ...string) geo.Geocoder {
	var url string
	if len(baseURLs) > 0 {
		url = baseURLs[0]
	} else {
		url = "http://open.mapquestapi.com/nominatim/v1/"
	}
	key = k
	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL(url),
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func (b baseURL) GeocodeURL(address string) string {
	return string(b) + "search.php?key=" + key + "&format=json&limit=1&q=" + address
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return string(b) + "reverse.php?key=" + key + fmt.Sprintf("&format=json&lat=%f&lon=%f", l.Lat, l.Lng)
}

func (r *geocodeResponse) Location() (l geo.Location) {
	if r.Error == "" {
		l = geo.Location{parseFloat(r.Lat), parseFloat(r.Lon)}
	}
	return
}

func (r *geocodeResponse) Address() (address string) {
	if r.Error == "" {
		address = r.DisplayName
	}
	return
}

func parseFloat(value interface{}) float64 {
	f, _ := strconv.ParseFloat(value.(string), 64)
	return f
}
