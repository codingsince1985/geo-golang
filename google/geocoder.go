// Package google is a geo-golang based Google Maps geocode/reverse geocode client
package google

import (
	"encoding/json"
	"fmt"
	"github.com/codingsince1985/geo-golang"
)

type Endpoint geo.Endpoint

type GeocodeResponse struct {
	Results []struct {
		Formatted_Address string
		Geometry          struct {
			Location struct {
				Lat, Lng float64
			}
		}
	}
}

func NewGeocoder() geo.Geocoder {
	return geo.Geocoder{
		Endpoint("http://maps.googleapis.com/maps/api/geocode/json?sensor=false&"),
		GeocodeResponse{},
	}
}

func (e Endpoint) GeocodeUrl(address string) string {
	return string(e) + "address=" + address
}

func (e Endpoint) ReverseGeocodeUrl(l geo.Location) string {
	return string(e) + fmt.Sprintf("latlng=%f,%f", l.Lat, l.Lng)
}

func (r GeocodeResponse) Location(data []byte) (l geo.Location) {
	if json.Unmarshal(data, &r); len(r.Results) > 0 {
		loc := r.Results[0].Geometry.Location
		l = geo.Location{loc.Lat, loc.Lng}
	}
	return
}

func (r GeocodeResponse) Address(data []byte) (address string) {
	if json.Unmarshal(data, &r); len(r.Results) > 0 {
		address = r.Results[0].Formatted_Address
	}
	return
}
