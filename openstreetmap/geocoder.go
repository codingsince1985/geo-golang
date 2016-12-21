// Package openstreetmap is a geo-golang based OpenStreetMap geocode/reverse geocode client
package openstreetmap

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

// Geocoder constructs OpenStreetMap geocoder
func Geocoder() geo.Geocoder { return GeocoderWithURL("https://nominatim.openstreetmap.org/") }

// GeocoderWithURL constructs OpenStreetMap geocoder using a custom installation of Nominatim
func GeocoderWithURL(nominatimURL string) geo.Geocoder {
	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL(nominatimURL),
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func (b baseURL) GeocodeURL(address string) string {
	return string(b) + "search?format=json&limit=1&q=" + address
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return string(b) + "reverse?" + fmt.Sprintf("format=json&lat=%f&lon=%f", l.Lat, l.Lng)
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
