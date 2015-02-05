// Package opencage is a geo-golang based OpenCage geocode/reverse geocode client
package opencage

import (
	"encoding/json"
	"fmt"
	"github.com/codingsince1985/geo-golang"
)

type Endpoint geo.Endpoint

type GeocodeResponse struct {
	Results []struct {
		Formatted string
		Geometry  struct {
			Lat, Lng float64
		}
	}
}

func NewGeocoder(key string) geo.Geocoder {
	return geo.Geocoder{
		Endpoint("http://api.opencagedata.com/geocode/v1/json?key=" + key + "&q="),
		GeocodeResponse{},
	}
}

func (e Endpoint) GeocodeUrl(address string) string {
	return string(e) + address
}

func (e Endpoint) ReverseGeocodeUrl(l geo.Location) string {
	return string(e) + fmt.Sprintf("%+f,%+f", l.Lat, l.Lng)
}

func (r GeocodeResponse) Location(data []byte) (l geo.Location) {
	if json.Unmarshal(data, &r); len(r.Results) > 0 {
		g := r.Results[0].Geometry
		l = geo.Location{g.Lat, g.Lng}
	}
	return
}

func (r GeocodeResponse) Address(data []byte) (address string) {
	if json.Unmarshal(data, &r); len(r.Results) > 0 {
		address = r.Results[0].Formatted
	}
	return
}
