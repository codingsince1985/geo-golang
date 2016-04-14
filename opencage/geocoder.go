// Package opencage is a geo-golang based OpenCage geocode/reverse geocode client
package opencage

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
)

type baseURL string

type geocodeResponse struct {
	Results []struct {
		Formatted string
		Geometry  geo.Location
	}
}

// Geocoder constructs OpenCage geocoder
func Geocoder(key string) geo.Geocoder {
	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL("http://api.opencagedata.com/geocode/v1/json?key=" + key + "&q="),
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func (b baseURL) GeocodeURL(address string) string {
	return string(b) + address
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return string(b) + fmt.Sprintf("%+f,%+f", l.Lat, l.Lng)
}

func (r *geocodeResponse) Location() (l geo.Location) {
	if len(r.Results) > 0 {
		l = r.Results[0].Geometry
	}
	return
}

func (r *geocodeResponse) Address() (address string) {
	if len(r.Results) > 0 {
		address = r.Results[0].Formatted
	}
	return
}
