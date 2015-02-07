// Package mapquest is a geo-golang based MapRequest geocode/reverse geocode client
package mapquest

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"strconv"
)

type baseUrl string

type geocodeResponse map[string]interface{}

func Geocoder() geo.Geocoder {
	return geo.Geocoder{
		baseUrl("http://open.mapquestapi.com/nominatim/v1/"),
		&geocodeResponse{},
	}
}

func (b baseUrl) GeocodeUrl(address string) string {
	return string(b) + "search.php?format=json&limit=1&q=" + address
}

func (b baseUrl) ReverseGeocodeUrl(l geo.Location) string {
	return string(b) + fmt.Sprintf("reverse.php?format=json&lat=%f&lon=%f", l.Lat, l.Lng)
}

func (r *geocodeResponse) Location() (l geo.Location) {
	if (*r)["lat"] != nil && (*r)["lon"] != nil {
		return geo.Location{parseFloat((*r)["lat"]), parseFloat((*r)["lon"])}
	}
	return
}

func (r *geocodeResponse) Address() (address string) {
	if (*r)["error"] == nil {
		address = (*r)["display_name"].(string)
	}
	return
}

func (r *geocodeResponse) ResponseObject() geo.ResponseParser {
	return r
}

func parseFloat(value interface{}) float64 {
	f, _ := strconv.ParseFloat(value.(string), 64)
	return f
}
