// Package google is a geo-golang based Google Maps geocode/reverse geocode client
package google

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
)

type baseUrl string

type geocodeResponse struct {
	Results []struct {
		Formatted_Address string
		Geometry          struct {
			Location struct {
				Lat, Lng float64
			}
		}
	}
}

func Geocoder() geo.Geocoder {
	return geo.Geocoder{
		baseUrl("http://maps.googleapis.com/maps/api/geocode/json?sensor=false&"),
		&geocodeResponse{},
	}
}

func (b baseUrl) GeocodeUrl(address string) string {
	return string(b) + "address=" + address
}

func (b baseUrl) ReverseGeocodeUrl(l geo.Location) string {
	return string(b) + fmt.Sprintf("latlng=%f,%f", l.Lat, l.Lng)
}

func (r *geocodeResponse) Location() (l geo.Location) {
	if len(r.Results) > 0 {
		loc := r.Results[0].Geometry.Location
		l = geo.Location{loc.Lat, loc.Lng}
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
