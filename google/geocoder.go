// Package google is a geo-golang based Google Maps geocode/reverse geocode client
package google

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
)

type baseURL string

type geocodeResponse struct {
	Results []struct {
		Formatted_Address string
		Geometry          struct {
			Location geo.Location
		}
	}
}

// Geocoder constructs Google geocoder
func Geocoder() geo.Geocoder {
	return geo.Geocoder{
		baseURL("http://maps.googleapis.com/maps/api/geocode/json?sensor=false&"),
		&geocodeResponse{},
	}
}

func (b baseURL) GeocodeURL(address string) string {
	return string(b) + "address=" + address
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return string(b) + fmt.Sprintf("latlng=%f,%f", l.Lat, l.Lng)
}

func (r *geocodeResponse) Location() (l geo.Location) {
	if len(r.Results) > 0 {
		l = r.Results[0].Geometry.Location
	}
	return
}

func (r *geocodeResponse) Address() (address string) {
	if len(r.Results) > 0 {
		address = r.Results[0].Formatted_Address
	}
	return
}

func (r *geocodeResponse) ResponseObject() geo.ResponseParser {
	return r
}
