// Package google has the implementation of using Google geocode and reverse geocode, in ~50 LoC.
package google

import (
	"encoding/json"
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"net/url"
)

const GEOCODE_BASE_URL = "http://maps.googleapis.com/maps/api/geocode/json?sensor=false"

type Endpoint geo.Endpoint
type GeocodeResponse struct {
	Results []struct {
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Location struct {
				Lat, Lng float64
			}
		}
	}
}

var Geocoder = geo.Geocoder{Endpoint{GEOCODE_BASE_URL}, GeocodeResponse{}}

func (e Endpoint) GeocodeUrl(address string) string {
	return e.BaseUrl + "&address=" + url.QueryEscape(address)
}

func (e Endpoint) ReverseGeocodeUrl(l geo.Location) string {
	return e.BaseUrl + fmt.Sprintf("&latlng=%f,%f", l.Lat, l.Lng)
}

func (r GeocodeResponse) Location(data []byte) (location geo.Location) {
	if json.Unmarshal(data, &r); len(r.Results) > 0 {
		l := r.Results[0].Geometry.Location
		location = geo.Location{l.Lat, l.Lng}
	}
	return
}

func (r GeocodeResponse) Address(data []byte) string {
	if json.Unmarshal(data, &r); len(r.Results) > 0 {
		return r.Results[0].FormattedAddress
	}
	return ""
}
