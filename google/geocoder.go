// Package google is a geo-golang based Google Geo Location API
// https://developers.google.com/maps/documentation/geocoding/intro
package google

import (
	"fmt"
	geo "github.com/codingsince1985/geo-golang"
)

type baseURL string

type geocodeResponse struct {
	Results []struct {
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Location geo.Location
		}
	}
}

// Geocoder constructs Google geocoder
func Geocoder(apiKey string) geo.Geocoder {
	return geo.Geocoder{
		EndpointBuilder: baseURL(fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?&key=%s&", apiKey)),
		ResponseParser:  &geocodeResponse{},
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
		address = r.Results[0].FormattedAddress
	}
	return
}

func (r *geocodeResponse) ResponseObject() geo.ResponseParser {
	return r
}
