// Package nominatim is a geo-golang based MapRequest Nominatim geocode/reverse geocode client
package nominatim

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"strconv"
)

type baseURL string

type geocodeResponse struct {
	Display_Name, Lat, Lon, Error string
}

func Geocoder() geo.Geocoder {
	return geo.Geocoder{
		baseURL("http://open.mapquestapi.com/nominatim/v1/"),
		&geocodeResponse{},
	}
}

func (b baseURL) GeocodeURL(address string) string {
	return string(b) + "search.php?format=json&limit=1&q=" + address
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return string(b) + fmt.Sprintf("reverse.php?format=json&lat=%f&lon=%f", l.Lat, l.Lng)
}

func (r *geocodeResponse) Location() (l geo.Location) {
	if r.Error == "" {
		l = geo.Location{parseFloat(r.Lat), parseFloat(r.Lon)}
	}
	return
}

func (r *geocodeResponse) Address() (address string) {
	if r.Error == "" {
		address = r.Display_Name
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
