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
		Geometry  struct {
			Lat, Lng float64
		}
	}
}

// Geocoder constructs OpenCage geocoder
func Geocoder(key string) geo.Geocoder {
	return geo.Geocoder{
		baseURL("http://api.opencagedata.com/geocode/v1/json?key=" + key + "&q="),
		&geocodeResponse{},
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
		g := r.Results[0].Geometry
		l = geo.Location{g.Lat, g.Lng}
	}
	return
}

func (r *geocodeResponse) Address() (address string) {
	if len(r.Results) > 0 {
		address = r.Results[0].Formatted
	}
	return
}

func (r *geocodeResponse) ResponseObject() geo.ResponseParser {
	return r
}
