// Package opencage has the implementation of OpenCage Data geocode and reverse geocode
package opencage

import (
	"encoding/json"
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"net/url"
)

const GEOCODE_BASE_URL = "https://api.opencagedata.com/geocode/v1/json?key=YOURKEY&q="

type Endpoint geo.Endpoint
type GeocodeResponse struct {
	Results []struct {
		Formatted string
		Geometry  struct {
			Lat float64
			Lng float64
		}
	}
}

var Geocoder = geo.Geocoder{Endpoint{GEOCODE_BASE_URL}, GeocodeResponse{}}

func (e Endpoint) GeocodeUrl(address string) string {
	return e.BaseUrl + url.QueryEscape(address)
}

func (e Endpoint) ReverseGeocodeUrl(l geo.Location) string {
	return e.BaseUrl + fmt.Sprintf("%+f,%+f", l.Lat, l.Lng)
}

func (r GeocodeResponse) Location(data []byte) (location geo.Location) {
	if json.Unmarshal(data, &r); len(r.Results) > 0 {
		g := r.Results[0].Geometry
		location = geo.Location{g.Lat, g.Lng}
	}
	return
}

func (r GeocodeResponse) Address(data []byte) (address string) {
	if json.Unmarshal(data, &r); len(r.Results) > 0 {
		address = r.Results[0].Formatted
	}
	return
}
