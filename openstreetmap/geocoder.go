// Package openstreetmap is a geo-golang based OpenStreetMap geocode/reverse geocode client
package openstreetmap

import (
	"fmt"
	"strconv"

	"github.com/codingsince1985/geo-golang"
)

type baseURL string

type geocodeResponse struct {
	DisplayName     string `json:"display_name"`
	Lat, Lon, Error string
}

// Geocoder constructs OpenStreetMap geocoder
func Geocoder() geo.Geocoder {
	return geo.Geocoder{
		baseURL("https://nominatim.openstreetmap.org/"),
		&geocodeResponse{},
	}
}

func (b baseURL) GeocodeURL(address string) string {
	return string(b) + "search?format=json&limit=1&q=" + address
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return string(b) + "reverse?" + fmt.Sprintf("format=json&lat=%f&lon=%f", l.Lat, l.Lng)
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

func (r *geocodeResponse) ResponseObject() geo.ResponseParser {
	return r
}

func parseFloat(value interface{}) float64 {
	f, _ := strconv.ParseFloat(value.(string), 64)
	return f
}
